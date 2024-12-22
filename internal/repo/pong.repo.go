package repo

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
	//model
}

// Pong implements IPongRepo.
func (p *pongRepo) Pong() string {
	return "pong-repo..."
}

func NewPongRepo() IPongRepo {
	return &pongRepo{}
}
