package memory

import (
	"encoding/json"
	"errors"
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
	ID                             int64
	Name                           string
	Brokers                        string
	Memories                       []Memory
	TopicsConfigs                  []config.TopicConfig
	Type                           string
	memoryContentReceiverRoutines  []func(memory Memory, consumer *kafka.Consumer, topicConfig config.TopicConfig)
	memoryContentPublisherRoutines []func(memory Memory, producer *kafka.Producer, topicConfig config.TopicConfig)
	Timestamp                      int64
}

func (distributedMemory *DistributedMemory) GetID() int64 {
	return distributedMemory.ID
}

func (distributedMemory *DistributedMemory) SetID(id int64) {
	distributedMemory.ID = id
}

func (distributedMemory *DistributedMemory) GetName() string {
	return distributedMemory.Name
}

func (distributedMemory *DistributedMemory) SetName(name string) {
	distributedMemory.Name = name
}

func (distributedMemory *DistributedMemory) GetI() interface{} {
	var i interface{} = nil

	var maxMemoryEvaluation float64 = 0

	for _, memory := range distributedMemory.Memories {

		memoryEvaluation := memory.GetEvaluation()

		if memoryEvaluation >= maxMemoryEvaluation {
			i = memory.GetI()
			maxMemoryEvaluation = memoryEvaluation
		}
	}

	return i
}

func (distributedMemory *DistributedMemory) GetIByIndex(index int64) interface{} {
	if len(distributedMemory.Memories) > 0 {
		if distributedMemory.Memories[index] != nil {
			return distributedMemory.Memories[index].GetI()
		} else {
			message := "Index does not exist."

			err := errors.New(message)
			handler.ErrorCheck(err, message)
		}
	}

	return -1
}

func (distributedMemory *DistributedMemory) SetI(i interface{}) {
	if len(distributedMemory.Memories) > 0 {
		distributedMemory.Memories[0].SetI(i)
		distributedMemory.TopicsConfigs[0].Command <- true
		distributedMemory.SetTimestamp(time.Now().Unix())
	}
}

func (distributedMemory *DistributedMemory) SetIByIndex(i interface{}, index int64) {
	if len(distributedMemory.Memories) > 0 {
		if distributedMemory.Memories[index] != nil {
			distributedMemory.Memories[index].SetI(i)
			distributedMemory.TopicsConfigs[index].Command <- true
			distributedMemory.SetTimestamp(time.Now().Unix())
		} else {
			message := "Index does not exist."

			err := errors.New(message)
			handler.ErrorCheck(err, message)
		}
	}
}

func (distributedMemory *DistributedMemory) SetEvaluation(evaluation float64) {
	if len(distributedMemory.Memories) > 0 {
		distributedMemory.Memories[0].SetEvaluation(evaluation)
		distributedMemory.TopicsConfigs[0].Command <- true
		distributedMemory.SetTimestamp(time.Now().Unix())
	}
}

func (distributedMemory *DistributedMemory) SetEvaluationByIndex(evaluation float64, index int64) {
	if len(distributedMemory.Memories) > 0 {
		if distributedMemory.Memories[index] != nil {
			distributedMemory.Memories[index].SetEvaluation(evaluation)
			distributedMemory.TopicsConfigs[index].Command <- true
			distributedMemory.SetTimestamp(time.Now().Unix())
		} else {
			message := "Index does not exist."

			err := errors.New(message)
			handler.ErrorCheck(err, message)
		}
	}
}

func (distributedMemory *DistributedMemory) GetEvaluation() float64 {

	var maxMemoryEvaluation float64 = 0

	for _, memory := range distributedMemory.Memories {

		memoryEvaluation := memory.GetEvaluation()

		if memoryEvaluation >= maxMemoryEvaluation {
			maxMemoryEvaluation = memoryEvaluation
		}
	}

	return maxMemoryEvaluation
}

func (distributedMemory *DistributedMemory) GetEvaluationByIndex(index int64) float64 {
	if len(distributedMemory.Memories) > 0 {
		if distributedMemory.Memories[index] != nil {
			return distributedMemory.Memories[index].GetEvaluation()
		} else {
			message := "Index does not exist."

			err := errors.New(message)
			handler.ErrorCheck(err, message)
		}
	}

	return -1
}

func (distributedMemory *DistributedMemory) InitMemory(brokers string) {

	distributedMemory.Brokers = brokers
	distributedMemory.SetTimestamp(time.Now().Unix())

	if distributedMemory.Type == constants.INPUT_MEMORY {
		distributedMemory.consumerSetup(distributedMemory.TopicsConfigs)
	} else {
		distributedMemory.producerSetup(distributedMemory.TopicsConfigs)
	}

}

func (distributedMemory *DistributedMemory) SetTimestamp(timestamp int64) {
	distributedMemory.Timestamp = timestamp
}

func (distributedMemory *DistributedMemory) GetTimestamp() int64 {
	return distributedMemory.Timestamp
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

		topicsCofigs[i].Command = make(chan bool)

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
		if topicConfig.DistributedMemoryBehavior == constants.DISTRIBUTED_MEMORY_BEHAVIOR_TRIGGERED {

			command := <-topicConfig.Command

			memoryJson, err := json.Marshal(memory)

			handler.ErrorCheck(err, "Error to convert memory to json.")

			if command {
				producer.Produce(&kafka.Message{
					TopicPartition: kafka.TopicPartition{Topic: &topicConfig.Name,
						Partition: kafka.PartitionAny},
					Value: []byte(memoryJson)}, nil)
			}

		} else {

			if memory.GetI() != lastI || memory.GetEvaluation() != float64(lastEvaluation) {

				memoryJson, err := json.Marshal(memory)

				handler.ErrorCheck(err, "Error to convert memory to json.")

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
		// initialTime := time.Now().Unix()

		message, _ := consumer.ReadMessage(-1)

		// handler.ErrorCheck(err, fmt.Sprintf("Error to receive message from topic %s", topicConfig.Name))

		if topicConfig.StructName != nil {
			object := reflect.New(topicConfig.StructName).Elem()

			json.Unmarshal(message.Value, memory)

			iField := []byte(fmt.Sprintf("%v", memory.GetI()))

			json.Unmarshal(iField, object)

			memory.SetI(object)
		} else {
			json.Unmarshal(message.Value, memory)
		}
	}
}
