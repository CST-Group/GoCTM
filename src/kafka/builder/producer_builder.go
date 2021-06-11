package builder

import (
	"godct/handler"
	"godct/kafka/config"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func BuildProducer(brokers string) *kafka.Producer {

	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": brokers})

	handler.ErrorCheck(err, "Error to create the producer.")

	return producer
}

func GenerateProducers(topicsConfigs []config.TopicConfig, brokers string) []*kafka.Producer {

	var producers []*kafka.Producer

	for i := 0; i < len(topicsConfigs); i++ {
		producers = append(producers, BuildProducer(brokers))
	}

	return producers
}
