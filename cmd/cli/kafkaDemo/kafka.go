package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
	kafka "github.com/segmentio/kafka-go"
)

var (
	wg       sync.WaitGroup
	producer *kafka.Writer
	exit     atomic.Bool
)

const (
	kafkaUrl = "127.0.0.1:9092"
	topic    = "test"
)

func closeConsumer(consumer *kafka.Reader, id int) {
	fmt.Printf("shutdown consummer %v....\n", id)
	if err := consumer.Close(); err != nil {
		fmt.Printf("failed to close reader: %s", err)
	}
}

func registerConsumer(id int) {
	defer wg.Done()
	consumer := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{kafkaUrl},
		Topic:          topic,
		GroupID:        fmt.Sprintf("consumer-group-%v", id),
		Partition:      0,
		MaxBytes:       10e6, // 10MB
		StartOffset:    kafka.FirstOffset,
		CommitInterval: time.Second,
	})

	defer closeConsumer(consumer, id)

	for {
		ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
		m, err := consumer.ReadMessage(ctx)
		//m, err := consumer.FetchMessage(ctx)
		if err != nil {
			if ctx.Err() == context.DeadlineExceeded && exit.Load() {
				break
			} else if ctx.Err() != context.DeadlineExceeded {
				fmt.Printf("failed to commit messages: %s \n", err)
				break
			} else {
				continue
			}
		}
		fmt.Printf("consumer %v - read at topic: %s- partition: %v - offset: %v - message: %s = %s\n", id, m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		// if err := consumer.CommitMessages(ctx, m); err != nil {
		// 	fmt.Printf("failed to commit messages: %s \n", err)
		// }
	}

	// if err := consumer.Close(); err != nil {
	// 	log.Fatal("failed to close reader:", err)
	// }
	//defer closeConsumer(consumer, id)
}

// func checkStatus() {
// 	for {
// 		time.Sleep(2 * time.Second)
// 		fmt.Printf("status exit %v \n", exit.Load())
// 	}
// }

func closeProducer() {
	fmt.Println("shutdown producer....")
	if err := producer.Close(); err != nil {
		fmt.Printf("failed to close producer: %s", err)
	}
}

func initProducer() {
	producer = &kafka.Writer{
		Addr:                   kafka.TCP(kafkaUrl),
		Topic:                  topic,
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: true,
	}
	//defer closeProducer()
}

func pushMessAction(key string, value string) {
	err := producer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(key),
			Value: []byte(value),
		},
	)
	if err != nil {
		fmt.Printf("failed to write messages : %s \n", err)
	}
}

func main() {
	exit.Store(false)
	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, os.Interrupt)

	r := gin.Default()
	initProducer()
	r.POST("/message", func(ctx *gin.Context) {
		pushMessAction(ctx.Query("key"), ctx.Query("value"))
	})

	wg.Add(2)
	go registerConsumer(1)
	go registerConsumer(2)
	//go checkStatus()
	go func() {
		r.Run(":8081")
	}()

	<-exitSignal
	fmt.Println("cleanup resource before shutdow...")
	exit.Store(true)
	time.Sleep(1 * time.Second)
	closeProducer()
	wg.Wait()
	fmt.Println("shutdow success...")
}
