package application

import (
	"godct/codelet"
	"godct/constants"
	"godct/memory"
)

type CodeletApplication struct {
	Brokers         string
	Codelets        []codelet.Codelet
	CodeletMemories map[*codelet.Codelet][]memory.DistributedMemory
}

func (codeletApplication *CodeletApplication) SetupCodeletApplication() {
	codeletApplication.initializeDistributedMemories()
	codeletApplication.startCodelets()
}

func (codeletApplication *CodeletApplication) startCodelets() {
	for _, codelet := range codeletApplication.Codelets {
		codelet.Start()
	}
}

func (codeletApplication *CodeletApplication) initializeDistributedMemories() {
	for codelet, memories := range codeletApplication.CodeletMemories {
		for _, memory := range memories {

			memory.InitMemory(codeletApplication.Brokers)

			if memory.Type == constants.INPUT_MEMORY {
				codelet.AddInput(&memory)
			} else {
				codelet.AddOutput(&memory)
			}

		}

		codeletApplication.Codelets = append(codeletApplication.Codelets, *codelet)
	}
}
