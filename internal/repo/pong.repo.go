package repo

import "fmt"

// type PongRepo struct {
// 	//model
// }

// func NewPongRepo() *PongRepo {
// 	return &PongRepo{}
// }

// func (pr *PongRepo) Pong() string {
// 	return "pong-repo..."
// }

type IPongRepo interface {
	Pong() string
}

type pongRepo struct {
	id int
	//model
}

// Pong implements IPongRepo.
func (p *pongRepo) Pong() string {
	fmt.Printf("pointer: %v \n", &p)
	p.id = p.id + 1
	return fmt.Sprintf("pong-repo-id: %v", p.id)
}

func NewPongRepo(id int) IPongRepo {
	return &pongRepo{
		id: id,
	}
}
