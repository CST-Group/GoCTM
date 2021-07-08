package builder

import (
	"errors"
	"fmt"
	"godct/constants"
	"godct/handler"
	"godct/kafka/config"
	"reflect"

	"regexp"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func BuildConsumer(brokers string, consumerGroupId string) *kafka.Consumer {

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  brokers,
		"group.id":           consumerGroupId,
		"auto.offset.reset":  "latest",
		"enable.auto.commit": true,
	})

	handler.ErrorCheck(err, "Error to connect in bootstrap servers.")

	return consumer

}

func generateTopicConfigsPrefix(brokers string, prefix string, structName reflect.Type) []config.TopicConfig {

	topicsFound := []config.TopicConfig{}

	consumer := BuildConsumer(brokers, "any")

	metadata, err := consumer.GetMetadata(nil, true, 5000)

	handler.ErrorCheck(err, "Error to get the topics metadata.")

	for key, _ := range metadata.Topics {

		re := regexp.MustCompile(prefix)

		if re.MatchString(key) {
			topicsFound = append(topicsFound, config.TopicConfig{
				Name:                      key,
				Prefix:                    prefix,
				DistributedMemoryBehavior: constants.DISTRIBUTED_MEMORY_BEHAVIOR_PULLED,
			})
		}

	}

	if len(topicsFound) == 0 {
		errorMessage := fmt.Sprintf("Topic prefix not found - Prefix - %s.", prefix)
		err = errors.New(errorMessage)

		handler.ErrorCheck(err, errorMessage)
	}

	return topicsFound

}

func GenerateConsumers(topicsConfigs []config.TopicConfig, brokers string, consumerGroupId string) map[*config.TopicConfig]*kafka.Consumer {

	consumers := map[*config.TopicConfig]*kafka.Consumer{}

	for _, topicConfig := range topicsConfigs {

		if len(topicConfig.Prefix) > 0 {
			foundTopics := generateTopicConfigsPrefix(brokers, topicConfig.Prefix, topicConfig.StructName)

			newConsumers := GenerateConsumers(foundTopics, brokers, consumerGroupId)

			for newFoundTopic, newConsumer := range newConsumers {
				consumers[newFoundTopic] = newConsumer
			}
		}

		consumer := BuildConsumer(brokers, consumerGroupId)

		createCheckTopic(brokers, topicConfig.Name, 1, 1)

		err := consumer.Subscribe(topicConfig.Name, nil)

		handler.ErrorCheck(err, "Error to subscribe consumer in topics.")

		consumers[&topicConfig] = consumer

	}

	return consumers
}
