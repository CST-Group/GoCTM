package application

import (
	"godct/codelet"
	"godct/memory"
)

type CodeletApplication struct {
	Brokers         string
	Codelets        []codelet.Codelet
	CodeletMemories map[*codelet.Codelet][]memory.DistributedMemory
}
