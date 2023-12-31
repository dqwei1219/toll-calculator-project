package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/dqwei1219/toll-calculator-project/aggregator/client"
	"github.com/dqwei1219/toll-calculator-project/types"
	"github.com/sirupsen/logrus"
)

type KafkaConsumer struct {
	consumer    *kafka.Consumer
	isRunning   bool
	calcService CalculatorServicer
	aggClient   client.Client
}

func NewKafkaComsumer(topic string, svc CalculatorServicer, aggClient client.Client) (*KafkaConsumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		return nil, err
	}
	c.SubscribeTopics([]string{topic}, nil)

	return &KafkaConsumer{
		consumer:    c,
		calcService: svc,
		aggClient:  aggClient,
	}, nil
}

func (c *KafkaConsumer) Start() {
	fmt.Println("Start consuming data")
	c.isRunning = true
	c.readMessageLoop()
}

func (c *KafkaConsumer) readMessageLoop() {
	for c.isRunning {
		msg, err := c.consumer.ReadMessage(-1) // timeout indefinitely
		if err != nil {
			logrus.Errorf("kafka consumer error: %v", err)
			continue
		}

		var data types.UnitCoordinate
		if err := json.Unmarshal(msg.Value, &data); err != nil {
			logrus.Errorf("error unmarshalling message: %v", err)
			continue
		}

		distance, err := c.calcService.CalculateDistance(data)
		if err != nil {
			logrus.Errorf("error calculating distance: %v", err)
			continue
		}

		req := &types.AggregateDistReq{
			Value:  distance,
			Unix:   time.Now().UnixNano(),
			UnitId: int32(data.UnitId),
		}

		if err := c.aggClient.AggregateDist(context.Background(), req); err != nil {
			logrus.Errorf("error aggregating distance: %v", err)
			continue
		}

		fmt.Printf("ID: %d, calculated distance: %.2f\n", data.UnitId, distance)
	}
}
