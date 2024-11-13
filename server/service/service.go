package service

type service struct {
	r Repositorier
}

func NewService(repo Repositorier) Servicer {
	return service{repo}
}
