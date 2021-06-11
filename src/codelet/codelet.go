package codelet

import (
	"errors"
	"godct/handler"
	"godct/memory"
	"strings"
	"time"
)

type Codelet struct {
	Activation float64 `default:"0"`
	Threshold  float64 `default:"0"`
	Enabled    bool    `default:"true"`
	Inputs     []memory.Memory
	Outputs    []memory.Memory
	Broadcasts []memory.Memory
	StartTime  int64
	EndTime    int64
	Duration   int64 `default:"1"`

	AccessMemoryObjects func()
	Proc                func()
	CalculateActivation func()
}

func (codelet *Codelet) AddInput(memory memory.Memory) {
	codelet.Inputs = append(codelet.Inputs, memory)
}

func (codelet *Codelet) AddOutput(memory memory.Memory) {
	codelet.Outputs = append(codelet.Outputs, memory)
}

func (codelet *Codelet) AddBroadcast(memory memory.Memory) {
	codelet.Broadcasts = append(codelet.Broadcasts, memory)
}

func (codelet *Codelet) GetOutput(name string) memory.Memory {
	return codelet.getMemoryInList(codelet.Outputs, name)
}

func (codelet *Codelet) GetInput(name string) memory.Memory {
	return codelet.getMemoryInList(codelet.Inputs, name)
}

func (codelet *Codelet) GetBroadcast(name string) memory.Memory {
	return codelet.getMemoryInList(codelet.Broadcasts, name)
}

func (codelet *Codelet) getMemoryInList(memories []memory.Memory, name string) memory.Memory {
	if len(memories) > 0 {
		for _, memory := range memories {
			if strings.EqualFold(strings.ToLower(memory.GetName()), strings.ToLower(name)) {
				return memory
			}
		}
	}

	return nil
}

func (codelet *Codelet) SetActivation(activation float64) {
	if activation > 1.0 {
		codelet.Activation = 1
	} else if activation < 0 {
		codelet.Activation = 0
	} else {
		codelet.Activation = activation
	}
}

func (codelet *Codelet) SetThreshold(threshold float64) {
	if threshold > 1 {
		codelet.Threshold = 1
	} else if threshold < 0 {
		codelet.Threshold = 0
	} else {
		codelet.Threshold = threshold
	}
}

func (codelet *Codelet) Start() {
	codelet.checkCodeletProperties()

	go codelet.run()
}

func (codelet *Codelet) checkCodeletProperties() {

	var err error = nil
	var message string = ""

	if codelet.Proc == nil {
		message = "Proc function not be nullable, please implement it"
		err = errors.New(message)
	}

	if codelet.CalculateActivation == nil {
		message = "CalculateActivation function not be nullable, please implement it"
		err = errors.New(message)
	}

	if codelet.AccessMemoryObjects == nil {
		message = "AccessMemoryObjects function not be nullable, please implement it"
		err = errors.New(message)
	}

	handler.ErrorCheck(err, message)

}

func (codelet *Codelet) run() {

	codelet.StartTime = time.Now().Unix()

	for codelet.Enabled {
		codelet.AccessMemoryObjects()

		codelet.CalculateActivation()

		if codelet.Activation >= codelet.Threshold {
			codelet.Proc()
		}

		time.Sleep(time.Duration(codelet.Duration) * time.Millisecond)
	}

	codelet.EndTime = time.Now().Unix()

}
