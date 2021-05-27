package memory

type MemoryObject struct {
	Memory

	ID         int64
	Name       string
	I          interface{}
	Evaluation float64
	Timestamp  int64
}
