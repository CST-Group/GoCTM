package memory

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

}

func (memoryContainer *MemoryContainer) SetIEvaluation(i interface{}, evaluation float64) {

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
		// System.out.println("Index for the " + getName()	+ ".getI(index) method greater than the number of MemoryObjects within the MemoryContainer");
		return nil
	}
}

func (memoryContainer *MemoryContainer) SetEvaluation(evaluation float64) {
	// memoryContainer.Evaluation = evaluation
}

func (memoryContainer *MemoryContainer) GetEvaluation() float64 {
	// return memoryContainer.Evaluation
	return 0
}
