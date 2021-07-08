package memory

import (
	"fmt"
	"time"
)

type MemoryContainer struct {
	ID        int64  `json:"id"`
	Name      string `json:"Name"`
	Memories  []Memory
	Timestamp int64 `json:"timestamp"`
}

func (memoryContainer *MemoryContainer) GetID() int64 {
	return memoryContainer.ID
}

func (memoryContainer *MemoryContainer) SetID(id int64) {
	memoryContainer.ID = id
}

func (memoryContainer *MemoryContainer) GetName() string {
	return memoryContainer.Name
}

func (memoryContainer *MemoryContainer) SetName(name string) {
	memoryContainer.Name = name
}

func (memoryContainer *MemoryContainer) SetI(i interface{}) {
	memoryContainer.SetIEvaluation(i, -1)
}

func (memoryContainer *MemoryContainer) checkMemory() {
	if memoryContainer.Memories == nil {
		memoryContainer.Memories = []Memory{}
	}
}

func (memoryContainer *MemoryContainer) SetIEvaluation(i interface{}, evaluation float64) {

	memoryContainer.checkMemory()

	memoryObject := MemoryObject{
		Evaluation: -1,
		I:          i,
	}

	if evaluation > -1.0 && evaluation < 1.0 {
		memoryObject.SetEvaluation(evaluation)
	}

	var memory Memory = &memoryObject

	memoryContainer.Memories = append(memoryContainer.Memories, memory)

	memoryContainer.Timestamp = time.Now().Unix()

}

func (memoryContainer *MemoryContainer) GetI() interface{} {

	var i interface{} = nil

	var maxMemoryEvaluation float64 = 0

	for _, memory := range memoryContainer.Memories {

		memoryEvaluation := memory.GetEvaluation()

		if memoryEvaluation >= maxMemoryEvaluation {
			i = memory.GetI()
			maxMemoryEvaluation = memoryEvaluation
		}
	}

	return i

}

func (memoryContainer *MemoryContainer) GetIByIndex(index int32) interface{} {
	if index >= 0 && index < int32(len(memoryContainer.Memories)) {
		return memoryContainer.Memories[index].GetI()
	} else {
		fmt.Errorf("Index %v for the method greater than the number of MemoryObjects within the MemoryContainer", index)
		return nil
	}
}

func (memoryContainer *MemoryContainer) SetEvaluation(evaluation float64) {
	memoryContainer.SetEvaluationByIndex(evaluation, 0)
}

func (memoryContainer *MemoryContainer) SetEvaluationByIndex(evaluation float64, index int32) {
	if memoryContainer.Memories != nil && len(memoryContainer.Memories) > int(index) {

		memory := memoryContainer.Memories[index]

		if memory != nil {

			memoryElement, ok := memory.(*MemoryContainer)

			if ok {
				memoryElement.SetEvaluationByIndex(evaluation, index)
			} else {
				memory.SetEvaluation(evaluation)
			}

			memoryContainer.Timestamp = time.Now().Unix()
		}
	}
}

func (memoryContainer *MemoryContainer) GetEvaluation() float64 {

	var maxMemoryEvaluation float64 = 0

	for _, memory := range memoryContainer.Memories {

		memoryEvaluation := memory.GetEvaluation()

		if memoryEvaluation >= maxMemoryEvaluation {
			maxMemoryEvaluation = memoryEvaluation
		}
	}

	return maxMemoryEvaluation
}

func (memoryContainer *MemoryContainer) Add(memory Memory) {
	memoryContainer.checkMemory()
	memoryContainer.Memories = append(memoryContainer.Memories, memory)
	memoryContainer.Timestamp = time.Now().Unix()
}

func (memoryContainer *MemoryContainer) GetTimestamp() int64 {
	return memoryContainer.Timestamp
}

func (memoryContainer *MemoryContainer) SetTimestamp(timestamp int64) {
	memoryContainer.Timestamp = timestamp
}
