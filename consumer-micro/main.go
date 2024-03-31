package main

import (
	"github.com/sirupsen/logrus"
)

var topic = "driver_distance"

func main() {
	cons, err := NewKafkaConsumer(topic)

	cons.Start()

	if err != nil {
		logrus.Fatalf("Error creating Kafka consumer: %v", err)
	}

}
