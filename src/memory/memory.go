package memory

type Memory interface {
	GetID() int64
	SetID(id int64)
	GetName() string
	SetName(name string)
	GetEvaluation() float64
	SetEvaluation(evalution float64)
	GetI() interface{}
	SetI(i interface{})
	SetTimestamp(timestamp int64)
	GetTimestamp() int64
}
