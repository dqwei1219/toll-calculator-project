package main

import (
	"math"
	"time"

	"github.com/dqwei1219/toll-calculator-project/types"
	"github.com/sirupsen/logrus"
)

type LogMiddleware struct {
	next DataProducer
}

func NewLogMiddleware(next DataProducer) *LogMiddleware {
	return &LogMiddleware{
		next: next,
	}
}

func (l *LogMiddleware) ProduceData(data types.UnitCoordinate) error {

	defer func(start time.Time) { // execute after return
		logrus.WithFields(logrus.Fields{
			"vehicle_id": data.UnitId,
			"latitude":   math.Trunc(data.Latitude),
			"longitude":  math.Trunc(data.Longitude),
			"duration":   time.Since(start),
		}).Info("Producing data to Kafka")
	}(time.Now())
	return l.next.ProduceData(data)
}
