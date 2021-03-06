package memory

import (
	"time"
)

func CreateMemoryObject(name string) Memory {
	memoryObject := MemoryObject{
		Name:       name,
		Timestamp:  time.Now().Unix(),
		Evaluation: 0,
		I:          nil,
	}

	var memory Memory = &memoryObject

	return memory
}

func CreateMemoryContainer(name string) MemoryContainer {
	memoryContainer := MemoryContainer{
		Name: name,
	}

	return memoryContainer
}
