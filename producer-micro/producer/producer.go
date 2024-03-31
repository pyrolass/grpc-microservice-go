package producer

import (
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	types "github.com/pyrolass/grpc-microservice-go/proto"
	"github.com/sirupsen/logrus"
)

type ProducerInterface interface {
	ProduceDriverDistance(data *types.InsertDriverLogRequest) error
}

type KafkaProducer struct {
	producer *kafka.Producer
	topic    string
}

func NewKafkaProducer(topic string) (ProducerInterface, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					logrus.Errorf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					logrus.Infof("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	if err != nil {
		return nil, err
	}

	return &KafkaProducer{
		producer: p,
		topic:    topic,
	}, nil
}

func (k *KafkaProducer) ProduceDriverDistance(data *types.InsertDriverLogRequest) error {

	b, err := json.Marshal(data)

	if err != nil {
		return err
	}

	err = k.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &k.topic, Partition: kafka.PartitionAny},
		Value:          b,
	}, nil)

	if err != nil {
		return err
	}

	return nil
}
