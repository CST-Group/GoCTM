package memory

type Memory interface {
	GetEvaluation() float64
	SetEvaluation(evalution float64)
	GetName() string
	SetName(name string)
	GetI() interface{}
	SetI(i interface{})
	GetTimestamp() int64
}
