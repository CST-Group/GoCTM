package memory

type Memory interface {
	GetEvaluation() float64
	SetEvaluation(evalution float64)
	GetI() interface{}
	SetI(i interface{})
}
