package service

import "go/go-backend-api/internal/repo"

type PongService struct {
	pongRepo *repo.PongRepo
}

func NewPongService() *PongService {
	return &PongService{
		pongRepo: repo.NewPongRepo(),
	}
}

func (ps *PongService) Pong() string {
	return ps.pongRepo.Pong()
}
