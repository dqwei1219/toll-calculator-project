package main

import (
	"fmt"

	"github.com/dqwei1219/toll-calculator-project/types"
)

type MemoryStore struct {
	data map[int]float64
}

func NewInMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make(map[int]float64),
	}
}

func (m *MemoryStore) Insert(distance types.Distance) error {
	m.data[distance.UnitId] += distance.Value
	return nil
}

func (m *MemoryStore) Get(unitId int) (float64, error) {
	dist, ok := m.data[unitId]
	if (!ok) {
		return 0.0, fmt.Errorf("could not find distance for id %d\n", unitId) 
	}
	return dist, nil
}