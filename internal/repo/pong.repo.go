package repo

type PongRepo struct {
	//model
}

func NewPongRepo() *PongRepo {
	return &PongRepo{}
}

func (pr *PongRepo) Pong() string {
	return "pong-repo..."
}
