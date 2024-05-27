package main

import (
	"fmt"

	"github.com/dr4g0n7ly/AutoTariff-Service/types"
)

const basePrice = 3.15

type Aggregator interface {
	AggregateDistance(types.Distance) error
	CalculateInvoice(int) (*types.Invoice, error)
}

type InvoiceAggregator struct {
	store Storer
}

func (i *InvoiceAggregator) AggregateDistance(distance types.Distance) error {
	fmt.Println("processing and inserting distance in the storage: ", distance)
	return i.store.Insert(distance)
}

func (i *InvoiceAggregator) CalculateInvoice(obuID int) (*types.Invoice, error) {
	fmt.Println("getting distance from storage for obuID: ", obuID)
	distance, err := i.store.GetDistance(obuID)
	if err != nil {
		return &types.Invoice{}, err
	}
	invoice := types.Invoice{
		OBUID:         obuID,
		TotalDistance: distance,
		TotalAmount:   distance * basePrice,
	}
	return &invoice, nil
}

func NewInvoiceAggregator(store Storer) Aggregator {
	return &InvoiceAggregator{
		store: store,
	}
}
