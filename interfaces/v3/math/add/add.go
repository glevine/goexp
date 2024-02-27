package add

type Service interface {
	Operate(a, b int) int
}

func NewService() Service {
	return &service{}
}

type service struct{}

func (svc *service) Operate(a, b int) int {
	return a + b
}
