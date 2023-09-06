package main

import (
	"math"

	"github.com/dqwei1219/toll-calculator-project/types"
)

type CalculatorServicer interface {
	CalculateDistance(types.UnitCoordinate) (float64, error)
}

type CalculatorService struct {
	prevLocation []float64
}

func NewCalculatorService() (*CalculatorService, error) {
	return &CalculatorService{
		prevLocation: make([]float64, 0),
	}, nil
}

func (s *CalculatorService) CalculateDistance(data types.UnitCoordinate) (float64, error) {
	dist := 0.0
	if len(s.prevLocation) > 0 {
		dist = calculateAbsDist(s.prevLocation[0], s.prevLocation[1],
			 data.Latitude, data.Longitude)
	}
	s.prevLocation = []float64{data.Latitude, data.Longitude}
	return dist, nil
}

func calculateAbsDist(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))
}
