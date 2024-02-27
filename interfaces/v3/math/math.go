package math

type Service interface {
	Operate(a, b int) int
}

func NewService(op Service) Service {
	return &service{
		op: op,
	}
}

type service struct {
	op Service
}

func (svc *service) Operate(a, b int) int {
	return svc.op.Operate(a, b)
}
