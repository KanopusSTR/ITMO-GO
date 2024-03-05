package workerpool

import (
	"context"
	"hw7/internal/model"
	"hw7/internal/pipeline"
	"sync"
)

type OrderWorkerPool interface {
	StartWorkerPool(
		ctx context.Context,
		orders <-chan model.OrderInitialized,
		additionalActions model.OrderActions,
		workersCount int,
	) <-chan model.OrderProcessFinished
}

type OrderWorkerPoolImplementation struct{}

func NewOrderWorkerPoolImplementation() *OrderWorkerPoolImplementation {
	return &OrderWorkerPoolImplementation{}
}

func (o *OrderWorkerPoolImplementation) StartWorkerPool(
	ctx context.Context,
	orders <-chan model.OrderInitialized,
	additionalActions model.OrderActions,
	workersCount int,
) <-chan model.OrderProcessFinished {
	retOrders := make(chan model.OrderProcessFinished)
	wg := sync.WaitGroup{}
	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go worker(ctx, additionalActions, &wg, orders, retOrders)
	}
	go func() {
		wg.Wait()
		close(retOrders)
	}()
	return retOrders
}

func worker(
	ctx context.Context, additionalActions model.OrderActions,
	wg *sync.WaitGroup,
	orders <-chan model.OrderInitialized,
	retOrders chan model.OrderProcessFinished,
) {
	defer wg.Done()
	pipeline.NewOrderPipelineImplementation().Start(ctx, additionalActions, orders, retOrders)
}
