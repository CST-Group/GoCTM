package memory

import (
	"godct/constants"
	"godct/kafka/config"
	"testing"
	"time"
)

func TestDistributedMemorySetI(t *testing.T) {

	message := "Test message in the distribute!"

	inputDistributedMemory, outputDistributedMemory := InitializeDistributedMemories()

	time.Sleep(5 * time.Second)

	outputDistributedMemory.SetI(message)

	for inputDistributedMemory.GetI() == nil {
	}

	if inputDistributedMemory.GetI() != message {
		t.Errorf("The input distributed memory does not receive the message!")
	}

}

func TestDistributedMemorySetIWithPrefix(t *testing.T) {

	message := "Test message in the distribute!"

	inputDistributedMemory, outputDistributedMemory := InitializeDistributedMemoriesWithPrefix()

	time.Sleep(5 * time.Second)

	outputDistributedMemory.SetI(message)

	for inputDistributedMemory.GetI() == nil {
	}

	if inputDistributedMemory.GetI() != message {
		t.Errorf("The input distributed memory does not receive the message!")
	}

}

func InitializeDistributedMemories() (DistributedMemory, DistributedMemory) {

	inputDistributedMemory := DistributedMemory{
		ID:   1,
		Name: "MEMORY_TEST_INPUT",
		Type: constants.INPUT_MEMORY,
		TopicsConfigs: []config.TopicConfig{
			{
				Name:                      "topic-1",
				DistributedMemoryBehavior: constants.DISTRIBUTED_MEMORY_BEHAVIOR_PULLED,
			},
		},
	}

	inputDistributedMemory.InitMemory("127.0.0.1:9092")

	outputDistributedMemory := DistributedMemory{
		ID:   2,
		Name: "MEMORY_TEST_OUTPUT",
		Type: constants.OUTPUT_MEMORY,
		TopicsConfigs: []config.TopicConfig{
			{
				Name:                      "topic-1",
				DistributedMemoryBehavior: constants.DISTRIBUTED_MEMORY_BEHAVIOR_TRIGGERED,
			},
		},
	}

	outputDistributedMemory.InitMemory("127.0.0.1:9092")

	return inputDistributedMemory, outputDistributedMemory
}

func InitializeDistributedMemoriesWithPrefix() (DistributedMemory, DistributedMemory) {

	inputDistributedMemory := DistributedMemory{
		ID:   1,
		Name: "MEMORY_TEST_INPUT",
		Type: constants.INPUT_MEMORY,
		TopicsConfigs: []config.TopicConfig{
			{
				DistributedMemoryBehavior: constants.DISTRIBUTED_MEMORY_BEHAVIOR_PULLED,
				Prefix:                    "topic-*",
			},
		},
	}

	inputDistributedMemory.InitMemory("127.0.0.1:9092")

	outputDistributedMemory := DistributedMemory{
		ID:   2,
		Name: "MEMORY_TEST_OUTPUT",
		Type: constants.OUTPUT_MEMORY,
		TopicsConfigs: []config.TopicConfig{
			{
				Name:                      "topic-1",
				DistributedMemoryBehavior: constants.DISTRIBUTED_MEMORY_BEHAVIOR_TRIGGERED,
			},
		},
	}

	outputDistributedMemory.InitMemory("127.0.0.1:9092")

	return inputDistributedMemory, outputDistributedMemory
}
