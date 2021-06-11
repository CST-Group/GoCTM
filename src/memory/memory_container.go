package memory

import "fmt"

type MemoryContainer struct {
	Memory

	ID       int64
	Name     string
	Memories []Memory
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

	var memoryObject = MemoryObject{
		Evaluation: -1,
		I:          i,
	}

	if evaluation != -1 {
		memoryObject.SetEvaluation(evaluation)
	}

	memoryContainer.Memories = append(memoryContainer.Memories, memoryObject.Memory)

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
	fmt.Errorf("This method is not available for MemoryContainer. Use setEvaluation(Double eval, int index) instead")
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
}
