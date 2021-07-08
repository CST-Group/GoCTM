package memory

import "time"

type MemoryObject struct {
	ID         int64       `json:"id"`
	Name       string      `json:"name"`
	I          interface{} `json:"i"`
	Evaluation float64     `json:"evaluation"`
	Timestamp  int64       `json:"timestamp"`
}

func (memoryObject *MemoryObject) GetID() int64 {
	return memoryObject.ID
}

func (memoryObject *MemoryObject) SetID(id int64) {
	memoryObject.ID = id
}

func (memoryObject *MemoryObject) GetName() string {
	return memoryObject.Name
}

func (memoryObject *MemoryObject) SetName(name string) {
	memoryObject.Name = name
}

func (memoryObject *MemoryObject) GetI() interface{} {
	return memoryObject.I
}

func (memoryObject *MemoryObject) SetI(i interface{}) {
	memoryObject.I = i
	memoryObject.Timestamp = time.Now().Unix()
}

func (memoryObject *MemoryObject) SetEvaluation(evaluation float64) {
	memoryObject.Evaluation = evaluation
	memoryObject.Timestamp = time.Now().Unix()
}

func (memoryObject *MemoryObject) GetEvaluation() float64 {
	return memoryObject.Evaluation
}

func (memoryObject *MemoryObject) GetTimestamp() int64 {
	return memoryObject.Timestamp
}

func (memoryObject *MemoryObject) SetTimestamp(timestamp int64) {
	memoryObject.Timestamp = timestamp
}
