package memory

type DistributedMemory struct {
	Memory

	ID         float64
	Name       string
	I          interface{}
	Evaluation float64
	Timestamp  int64
}

func (distributedMemory *DistributedMemory) GetName() string {
	return distributedMemory.Name
}

func (distributedMemory *DistributedMemory) SetName(name string) {
	distributedMemory.Name = name
}

func (distributedMemory *DistributedMemory) SetI(i interface{}) {
	distributedMemory.I = i
}

func (distributedMemory *DistributedMemory) SetEvaluation(evaluation float64) {
	distributedMemory.Evaluation = evaluation
}

func (distributedMemory *DistributedMemory) GetEvaluation() float64 {
	return distributedMemory.Evaluation
}
