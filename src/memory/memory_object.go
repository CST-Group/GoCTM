package memory

type MemoryObject struct {
	Memory

	ID         int64
	Name       string
	I          interface{}
	Evaluation float64
	Timestamp  int64
}

func (memoryObject *MemoryObject) GetName() string {
	return memoryObject.Name
}

func (memoryObject *MemoryObject) SetName(name string) {
	memoryObject.Name = name
}

func (memoryObject *MemoryObject) SetI(i interface{}) {
	memoryObject.I = i
}

func (memoryObject *MemoryObject) SetEvaluation(evaluation float64) {
	memoryObject.Evaluation = evaluation
}

func (memoryObject *MemoryObject) GetEvaluation() float64 {
	return memoryObject.Evaluation
}
