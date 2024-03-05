package generator

import (
	"context"
	"hw7/internal/model"
)

type OrderGenerator interface {
	GenerateOrdersStream(ctx context.Context, orders []model.OrderInitialized) <-chan model.OrderInitialized
}

type OrderGeneratorImplementation struct{}

func NewOrderGeneratorImplementation() *OrderGeneratorImplementation {
	return &OrderGeneratorImplementation{}
}

func (o *OrderGeneratorImplementation) GenerateOrdersStream(ctx context.Context, orders []model.OrderInitialized) <-chan model.OrderInitialized {
	retOrders := make(chan model.OrderInitialized, len(orders))
	for _, order := range orders {
		retOrders <- order
	}
	close(retOrders)
	return retOrders
}
