package main

import (
	"context"
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	types "github.com/pyrolass/grpc-microservice-go/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type KafkaConsumer struct {
	consumer  *kafka.Consumer
	isRunning bool
	conn      *grpc.ClientConn
}

func NewKafkaConsumer(topic string) (*KafkaConsumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"auto.offset.reset": "earliest",
		"group.id":          "myGroup",
	})

	c.SubscribeTopics([]string{topic}, nil)

	if err != nil {
		return nil, err
	}

	conn, err := grpc.Dial("localhost:3003", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}
	// TODO close connection
	// defer conn.Close()

	return &KafkaConsumer{consumer: c, conn: conn}, nil
}

func (c *KafkaConsumer) Start() {
	logrus.Info("Starting Kafka Consumer")
	c.isRunning = true

	c.readMessageLoop()
}

func (c *KafkaConsumer) readMessageLoop() {

	var messagesBatch []*types.InsertDriverLogRequest

	for c.isRunning {
		msg, err := c.consumer.ReadMessage(-1)

		if err != nil {
			logrus.Errorf("kafka consumer error: %v", err)
			continue
		}

		var driverData types.InsertDriverLogRequest

		err = json.Unmarshal(msg.Value, &driverData)

		if err != nil {
			logrus.Errorf("Error unmarshalling driver data: %v", err)
			continue
		}

		messagesBatch = append(messagesBatch, &driverData)

		if len(messagesBatch) == 5 {

			batchRequest := &types.InsertLogListRequest{Logs: messagesBatch}

			c.writeDataToDB(batchRequest)

			messagesBatch = []*types.InsertDriverLogRequest{}
		}

		logrus.Infof("Received driver data: %v", &driverData)

	}

}

func (c *KafkaConsumer) writeDataToDB(data *types.InsertLogListRequest) {
	client := types.NewDriverWriteServiceClient(c.conn)

	resp, err := client.InsertDriverLogDB(context.Background(), data)

	if err != nil {
		logrus.Errorf("Error writing to db: %v", err)
	}

	logrus.Infof("Response from db: %v", resp)
}
