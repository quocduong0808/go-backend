package service

import "go/go-backend-api/internal/repo"

// type PongService struct {
// 	pongRepo *repo.PongRepo
// }

// func NewPongService() *PongService {
// 	return &PongService{
// 		pongRepo: repo.NewPongRepo(),
// 	}
// }

// func (ps *PongService) Pong() string {
// 	return ps.pongRepo.Pong()
// }

type IPongService interface {
	Pong() string
}

type pongService struct {
	pongRepo repo.IPongRepo
}

// Pong implements IPongService.
func (ps *pongService) Pong() string {
	return ps.pongRepo.Pong()
}

func NewPongService(pongRepo repo.IPongRepo) IPongService {
	return &pongService{
		pongRepo: pongRepo,
	}
}
