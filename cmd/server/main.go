package main

import (
	"go/go-backend-api/internal/routers"
)

func main() {
	r := routers.CreateRouter()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
