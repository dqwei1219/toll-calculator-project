package main

import (
	"time"

	"github.com/dqwei1219/toll-calculator-project/types"
	"github.com/sirupsen/logrus"
)

type LogMiddleware struct {
	next Aggregator
}

func NewLogMiddleware(next Aggregator) *LogMiddleware {
	return &LogMiddleware{
		next: next,
	}
}

func (l *LogMiddleware) AggregateDistance(distance types.Distance) (err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"took": time.Since(start),
			"err":  err,
			"dist": distance,
		}).Info("Aggregating distance")
	}(time.Now())
	err = l.next.AggregateDistance(distance)
	return
}

func (l *LogMiddleware) CalculateInvoice(id int) (inv *types.Invoice, err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"took": time.Since(start),
			"err":  err,
			"id":   id,
			"inv":  inv,
		}).Info("Calculating invoice")
	}(time.Now())
	inv, err = l.next.CalculateInvoice(id)
	return
}