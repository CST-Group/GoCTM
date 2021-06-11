package memory

import (
	"encoding/json"
	"fmt"
	"godct/constants"
	"godct/handler"
	"godct/kafka/builder"
	"godct/kafka/config"
	"reflect"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type DistributedMemory struct {
	Memory

	ID                             int64
	Name                           string
	Brokers                        string
	Memories                       []Memory
	TopicsConfigs                  []config.TopicConfig
	Type                           string
	memoryContentReceiverRoutines  []func(memory Memory, consumer *kafka.Consumer, topicConfig config.TopicConfig)
	memoryContentPublisherRoutines []func(memory Memory, producer *kafka.Producer, topicConfig config.TopicConfig)
}

func (distributedMemory *DistributedMemory) GetI() interface{} {
	return nil
}

func (distributedMemory *DistributedMemory) SetI(i interface{}) {
	// distributedMemory.I = i
}

func (distributedMemory *DistributedMemory) SetEvaluation(evaluation float64) {
	// distributedMemory.Evaluation = evaluation
}

func (distributedMemory *DistributedMemory) GetEvaluation() float64 {
	return 0
}

func (distributedMemory *DistributedMemory) InitMemory(brokers string) {

	distributedMemory.Brokers = brokers

	if distributedMemory.Type == constants.INPUT_MEMORY {
		distributedMemory.consumerSetup(distributedMemory.TopicsConfigs)
	} else {
		distributedMemory.producerSetup(distributedMemory.TopicsConfigs)
	}

}

func (distributedMemory *DistributedMemory) consumerSetup(topicsCofigs []config.TopicConfig) {

	consumers := builder.GenerateConsumers(topicsCofigs, distributedMemory.Brokers, distributedMemory.Name)

	for topicConfig, consumer := range consumers {
		memoryObject := CreateMemoryObject(fmt.Sprintf("%v_DM", topicConfig.Name))

		distributedMemory.Memories = append(distributedMemory.Memories, memoryObject)

		receiverProccessFunction := distributedMemory.receiverProccess

		distributedMemory.memoryContentReceiverRoutines = append(distributedMemory.memoryContentReceiverRoutines, receiverProccessFunction)

		go receiverProccessFunction(memoryObject, consumer, *topicConfig)
	}

}

func (distributedMemory *DistributedMemory) producerSetup(topicsCofigs []config.TopicConfig) {

	producers := builder.GenerateProducers(topicsCofigs, distributedMemory.Brokers)

	for i := 0; i < len(producers); i++ {
		memoryObject := CreateMemoryObject(fmt.Sprintf("%v_DM", topicsCofigs[i].Name))

		distributedMemory.Memories = append(distributedMemory.Memories, memoryObject)

		publisherProcessFunction := distributedMemory.publisherProccess

		distributedMemory.memoryContentPublisherRoutines = append(distributedMemory.memoryContentPublisherRoutines, publisherProcessFunction)

		go publisherProcessFunction(memoryObject, producers[i], topicsCofigs[i])
	}
}

func (distributedMemory *DistributedMemory) publisherProccess(memory Memory, producer *kafka.Producer, topicConfig config.TopicConfig) {

	var lastI interface{} = nil
	var lastEvaluation float64 = 0

	defer producer.Close()

	for {
		memoryJson, err := json.Marshal(memory)

		handler.ErrorCheck(err, "Error to convert memory to json.")

		if topicConfig.DistributedMemoryBehavior == constants.DISTRIBUTED_MEMORY_BEHAVIOR_TRIGGERED {

			command := <-topicConfig.Command

			if command {
				producer.Produce(&kafka.Message{
					TopicPartition: kafka.TopicPartition{Topic: &topicConfig.Name,
						Partition: kafka.PartitionAny},
					Value: []byte(memoryJson)}, nil)
			}

		} else {

			if memory.GetI() != lastI || memory.GetEvaluation() != float64(lastEvaluation) {

				producer.Produce(&kafka.Message{
					TopicPartition: kafka.TopicPartition{Topic: &topicConfig.Name,
						Partition: kafka.PartitionAny},
					Value: []byte(memoryJson)}, nil)

				lastI = memory.GetI()
				lastEvaluation = memory.GetEvaluation()
			}

			time.Sleep(10 * time.Millisecond)

		}
	}

}

func (DistributedMemory *DistributedMemory) receiverProccess(memory Memory, consumer *kafka.Consumer, topicConfig config.TopicConfig) {

	defer consumer.Close()

	for {
		message, err := consumer.ReadMessage(10 * time.Millisecond)

		handler.ErrorCheck(err, fmt.Sprintf("Error to receive message from topic %s", topicConfig.Name))

		if topicConfig.StructName != nil {
			object := reflect.New(topicConfig.StructName).Elem()

			json.Unmarshal(message.Value, &object)

			memory.SetI(object)
		} else {
			memory.SetI(string(message.Value))
		}
	}

}
