package main

import (

	"github.com/dqwei1219/toll-calculator-project/types"
)
const basePrice = 5.0

type Aggregator interface {
	AggregateDistance(types.Distance) error
	CalculateInvoice(int) (*types.Invoice, error)
}

type Storer interface {
	Insert(types.Distance) error
	Get(int) (float64, error)
}

type InvoiceAggregator struct {
	store Storer
}

func NewInvoiceAggregator(store Storer) Aggregator{
	return &InvoiceAggregator{store: store}
}

func (i *InvoiceAggregator) AggregateDistance(distance types.Distance) error {
	return i.store.Insert(distance)
}

func (i *InvoiceAggregator) CalculateInvoice(id int) (*types.Invoice, error) {
	dist, err := i.store.Get(id)
	if err != nil {
		return nil, err

	}
	inv := &types.Invoice{
		UnitId: id,
		TotalDistance: dist,
		TotalCharge: basePrice * dist,
	}
	return inv, nil
}