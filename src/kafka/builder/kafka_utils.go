package builder

import (
	"context"
	"godct/handler"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func createCheckTopic(brokers string, topic string, partitionsNumber int, replicationFactor int) {

	kafkaAdminClient, err := kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers": brokers})

	handler.ErrorCheck(err, "Error to connect in bootstrap servers.")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	kafkaAdminClient.CreateTopics(
		ctx,
		[]kafka.TopicSpecification{{
			Topic:             topic,
			NumPartitions:     partitionsNumber,
			ReplicationFactor: replicationFactor}},
		kafka.SetAdminOperationTimeout(10000))

}
