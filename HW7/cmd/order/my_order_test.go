package main

import (
	"context"
	"github.com/stretchr/testify/require"
	"runtime"
	"slices"
	"testing"
	"time"

	"hw7/internal/generator"
	"hw7/internal/model"
	"hw7/internal/workerpool"
)

// Показательный тест. Чтобы он заработал, реализуйте generator
func TestSmallSuccessOneWorker(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	orders := []model.OrderInitialized{
		{
			OrderID:     7,
			ProductID:   5,
			OrderStates: []model.OrderState{model.Initialized},
		},
	}

	actions, countChecker := getDefaultAdditionalActions()
	orderWorkerPool := workerpool.NewOrderWorkerPoolImplementation()
	orderGenerator := generator.NewOrderGeneratorImplementation()

	result := orderWorkerPool.StartWorkerPool(ctx, orderGenerator.GenerateOrdersStream(ctx, orders), actions, 1)
	processedOrders := make([]model.Order, 0)

	for order := range result {
		processedOrders = append(processedOrders, finishedStateToOrder(order))
	}

	slices.SortFunc(processedOrders, func(a, b model.Order) int {
		return a.OrderID - b.OrderID
	})

	processedOrder := processedOrders[0]

	require.Equal(t, 1, len(processedOrders))
	require.Equal(t, 7, processedOrder.OrderID)
	require.Equal(t, 5, processedOrder.ProductID)
	require.Equal(t, 2, processedOrder.StorageID)
	require.Equal(t, 3, processedOrder.PickupPointID)
	require.Equal(t, true, processedOrder.IsProcessed)
	require.Equal(t, getFullOrderStates(), processedOrder.OrderStates)

	require.Equal(t, 1, countChecker.initToStartedCounter)
	require.Equal(t, 1, countChecker.StartedToFinishedExternalInteractionCounter)
	require.Equal(t, 1, countChecker.FinishedExternalInteractionToProcessFinishedCounter)
}

func TestSmallSuccessTenWorker(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	orders := []model.OrderInitialized{
		{
			OrderID:     7,
			ProductID:   5,
			OrderStates: []model.OrderState{model.Initialized},
		},
	}

	actions, countChecker := getDefaultAdditionalActions()
	orderWorkerPool := workerpool.NewOrderWorkerPoolImplementation()
	orderGenerator := generator.NewOrderGeneratorImplementation()

	result := orderWorkerPool.StartWorkerPool(ctx, orderGenerator.GenerateOrdersStream(ctx, orders), actions, 10)
	processedOrders := make([]model.Order, 0)

	for order := range result {
		processedOrders = append(processedOrders, finishedStateToOrder(order))
	}

	slices.SortFunc(processedOrders, func(a, b model.Order) int {
		return a.OrderID - b.OrderID
	})

	processedOrder := processedOrders[0]

	require.Equal(t, 1, len(processedOrders))
	require.Equal(t, 7, processedOrder.OrderID)
	require.Equal(t, 5, processedOrder.ProductID)
	require.Equal(t, 2, processedOrder.StorageID)
	require.Equal(t, 3, processedOrder.PickupPointID)
	require.Equal(t, true, processedOrder.IsProcessed)
	require.Equal(t, getFullOrderStates(), processedOrder.OrderStates)

	require.Equal(t, 1, countChecker.initToStartedCounter)
	require.Equal(t, 1, countChecker.StartedToFinishedExternalInteractionCounter)
	require.Equal(t, 1, countChecker.FinishedExternalInteractionToProcessFinishedCounter)
}

func TestPanicInitToStarted(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	orders := []model.OrderInitialized{
		{
			OrderID:     7,
			ProductID:   5,
			OrderStates: []model.OrderState{model.Initialized},
		},
	}

	actions, countChecker := getDefaultAdditionalActions()
	actions.InitToStarted = func() {
		countChecker.initToStartedCounter++
		panic("test")
	}

	orderWorkerPool := workerpool.NewOrderWorkerPoolImplementation()
	orderGenerator := generator.NewOrderGeneratorImplementation()

	result := orderWorkerPool.StartWorkerPool(ctx, orderGenerator.GenerateOrdersStream(ctx, orders), actions, 2)
	processedOrders := make([]model.Order, 0)

	for order := range result {
		processedOrders = append(processedOrders, finishedStateToOrder(order))
	}

	slices.SortFunc(processedOrders, func(a, b model.Order) int {
		return a.OrderID - b.OrderID
	})

	processedOrder := processedOrders[0]

	require.Equal(t, 1, len(processedOrders))
	require.Equal(t, 7, processedOrder.OrderID)
	require.Equal(t, 5, processedOrder.ProductID)
	require.Equal(t, 0, processedOrder.StorageID)
	require.Equal(t, 0, processedOrder.PickupPointID)
	require.Equal(t, false, processedOrder.IsProcessed)
	require.Equal(t, []model.OrderState{model.Initialized}, processedOrder.OrderStates)

	require.Equal(t, 1, countChecker.initToStartedCounter)
	require.Equal(t, 0, countChecker.StartedToFinishedExternalInteractionCounter)
	require.Equal(t, 0, countChecker.FinishedExternalInteractionToProcessFinishedCounter)
}

func TestPanicStartedToFinishedExternalInteraction(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	orders := []model.OrderInitialized{
		{
			OrderID:     7,
			ProductID:   5,
			OrderStates: []model.OrderState{model.Initialized},
		},
	}

	actions, countChecker := getDefaultAdditionalActions()
	actions.StartedToFinishedExternalInteraction = func() {
		countChecker.StartedToFinishedExternalInteractionCounter++
		panic("test")
	}

	orderWorkerPool := workerpool.NewOrderWorkerPoolImplementation()
	orderGenerator := generator.NewOrderGeneratorImplementation()

	result := orderWorkerPool.StartWorkerPool(ctx, orderGenerator.GenerateOrdersStream(ctx, orders), actions, 2)
	processedOrders := make([]model.Order, 0)

	for order := range result {
		processedOrders = append(processedOrders, finishedStateToOrder(order))
	}

	slices.SortFunc(processedOrders, func(a, b model.Order) int {
		return a.OrderID - b.OrderID
	})

	processedOrder := processedOrders[0]

	require.Equal(t, 1, len(processedOrders))
	require.Equal(t, 7, processedOrder.OrderID)
	require.Equal(t, 5, processedOrder.ProductID)
	require.Equal(t, 0, processedOrder.StorageID)
	require.Equal(t, 0, processedOrder.PickupPointID)
	require.Equal(t, false, processedOrder.IsProcessed)
	require.Equal(t, []model.OrderState{model.Initialized, model.ProcessStarted}, processedOrder.OrderStates)

	require.Equal(t, 1, countChecker.initToStartedCounter)
	require.Equal(t, 1, countChecker.StartedToFinishedExternalInteractionCounter)
	require.Equal(t, 0, countChecker.FinishedExternalInteractionToProcessFinishedCounter)
}

func TestPanicFinishedExternalInteractionToProcessFinished0(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	orders := []model.OrderInitialized{
		{
			OrderID:     7,
			ProductID:   5,
			OrderStates: []model.OrderState{model.Initialized},
		},
	}

	actions, countChecker := getDefaultAdditionalActions()
	actions.FinishedExternalInteractionToProcessFinished = func() {
		countChecker.FinishedExternalInteractionToProcessFinishedCounter++
		panic("test")
	}

	orderWorkerPool := workerpool.NewOrderWorkerPoolImplementation()
	orderGenerator := generator.NewOrderGeneratorImplementation()

	result := orderWorkerPool.StartWorkerPool(ctx, orderGenerator.GenerateOrdersStream(ctx, orders), actions, 2)
	processedOrders := make([]model.Order, 0)

	for order := range result {
		processedOrders = append(processedOrders, finishedStateToOrder(order))
	}

	slices.SortFunc(processedOrders, func(a, b model.Order) int {
		return a.OrderID - b.OrderID
	})

	processedOrder := processedOrders[0]

	require.Equal(t, 1, len(processedOrders))
	require.Equal(t, 7, processedOrder.OrderID)
	require.Equal(t, 5, processedOrder.ProductID)
	require.Equal(t, 2, processedOrder.StorageID)
	require.Equal(t, 3, processedOrder.PickupPointID)
	require.Equal(t, false, processedOrder.IsProcessed)
	require.Equal(t, []model.OrderState{model.Initialized, model.ProcessStarted, model.FinishedExternalInteraction},
		processedOrder.OrderStates)

	require.Equal(t, 1, countChecker.initToStartedCounter)
	require.Equal(t, 1, countChecker.StartedToFinishedExternalInteractionCounter)
	require.Equal(t, 1, countChecker.FinishedExternalInteractionToProcessFinishedCounter)
}

func TestBigSuccess(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	orders := []model.OrderInitialized{
		{
			OrderID:     6,
			ProductID:   19,
			OrderStates: []model.OrderState{model.Initialized},
		},
		{
			OrderID:     12424,
			ProductID:   9900,
			OrderStates: []model.OrderState{model.Initialized},
		},
		{
			OrderID:     55315,
			ProductID:   590531,
			OrderStates: []model.OrderState{model.Initialized},
		},
		{
			OrderID:     0,
			ProductID:   9,
			OrderStates: []model.OrderState{model.Initialized},
		},
		{
			OrderID:     12414,
			ProductID:   10,
			OrderStates: []model.OrderState{model.Initialized},
		},
		{
			OrderID:     95125,
			ProductID:   995,
			OrderStates: []model.OrderState{model.Initialized},
		},
		{
			OrderID:     110,
			ProductID:   0x5,
			OrderStates: []model.OrderState{model.Initialized},
		},
		{
			OrderID:     51,
			ProductID:   012,
			OrderStates: []model.OrderState{model.Initialized},
		},
	}

	actions, countChecker := getDefaultAdditionalActions()

	orderWorkerPool := workerpool.NewOrderWorkerPoolImplementation()
	orderGenerator := generator.NewOrderGeneratorImplementation()

	result := orderWorkerPool.StartWorkerPool(ctx, orderGenerator.GenerateOrdersStream(ctx, orders), actions, 5)
	processedOrders := make([]model.Order, 0)

	ticker := time.NewTicker(time.Second * 10)

	func() {
		for {
			select {
			case o, ok := <-result:
				if !ok {
					return
				}

				processedOrders = append(processedOrders, finishedStateToOrder(o))
			case <-ticker.C:
				t.Fail()
			}
		}
	}()

	slices.SortFunc(processedOrders, func(a, b model.Order) int {
		return a.OrderID - b.OrderID
	})

	require.Equal(t, 8, len(processedOrders))

	require.Equal(t, 8, countChecker.initToStartedCounter)
	require.Equal(t, 8, countChecker.StartedToFinishedExternalInteractionCounter)
	require.Equal(t, 8, countChecker.FinishedExternalInteractionToProcessFinishedCounter)
}

func TestBigSuccessFan(t *testing.T) {
	t.Parallel()

	runtime.GOMAXPROCS(1)
	ctx := context.Background()

	orders := []model.OrderInitialized{
		{
			OrderID:     6,
			ProductID:   19,
			OrderStates: []model.OrderState{model.Initialized},
		},
		{
			OrderID:     12424,
			ProductID:   9900,
			OrderStates: []model.OrderState{model.Initialized},
		},
		{
			OrderID:     55315,
			ProductID:   590531,
			OrderStates: []model.OrderState{model.Initialized},
		},
		{
			OrderID:     0,
			ProductID:   9,
			OrderStates: []model.OrderState{model.Initialized},
		},
		{
			OrderID:     12414,
			ProductID:   10,
			OrderStates: []model.OrderState{model.Initialized},
		},
		{
			OrderID:     95125,
			ProductID:   995,
			OrderStates: []model.OrderState{model.Initialized},
		},
		{
			OrderID:     110,
			ProductID:   0x5,
			OrderStates: []model.OrderState{model.Initialized},
		},
		{
			OrderID:     51,
			ProductID:   012,
			OrderStates: []model.OrderState{model.Initialized},
		},
	}

	actions, countChecker := getDefaultAdditionalActions()

	actions.StartedToFinishedExternalInteraction = func() {
		countChecker.StartedToFinishedExternalInteractionCounter++
		time.Sleep(time.Second * 5)
	}

	orderWorkerPool := workerpool.NewOrderWorkerPoolImplementation()
	orderGenerator := generator.NewOrderGeneratorImplementation()

	result := orderWorkerPool.StartWorkerPool(ctx, orderGenerator.GenerateOrdersStream(ctx, orders), actions, 1)
	processedOrders := make([]model.Order, 0)

	ticker := time.NewTicker(time.Second * 10)

	func() {
		for {
			select {
			case o, ok := <-result:
				if !ok {
					return
				}

				processedOrders = append(processedOrders, finishedStateToOrder(o))
			case <-ticker.C:
				t.Fail()
			}
		}
	}()

	slices.SortFunc(processedOrders, func(a, b model.Order) int {
		return a.OrderID - b.OrderID
	})

	require.Equal(t, 8, len(processedOrders))

	require.Equal(t, 8, countChecker.initToStartedCounter)
	require.Equal(t, 8, countChecker.StartedToFinishedExternalInteractionCounter)
	require.Equal(t, 8, countChecker.FinishedExternalInteractionToProcessFinishedCounter)
}

func TestBigPanic(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	orders := []model.OrderInitialized{
		{
			OrderID:     6,
			ProductID:   19,
			OrderStates: []model.OrderState{model.Initialized},
		},
		{
			OrderID:     12424,
			ProductID:   9900,
			OrderStates: []model.OrderState{model.Initialized},
		},
		{
			OrderID:     55315,
			ProductID:   590531,
			OrderStates: []model.OrderState{model.Initialized},
		},
		{
			OrderID:     0,
			ProductID:   9,
			OrderStates: []model.OrderState{model.Initialized},
		},
		{
			OrderID:     12414,
			ProductID:   10,
			OrderStates: []model.OrderState{model.Initialized},
		},
		{
			OrderID:     95125,
			ProductID:   995,
			OrderStates: []model.OrderState{model.Initialized},
		},
		{
			OrderID:     110,
			ProductID:   0x5,
			OrderStates: []model.OrderState{model.Initialized},
		},
		{
			OrderID:     51,
			ProductID:   012,
			OrderStates: []model.OrderState{model.Initialized},
		},
	}

	actions, countChecker := getDefaultAdditionalActions()

	actions.FinishedExternalInteractionToProcessFinished = func() {
		countChecker.FinishedExternalInteractionToProcessFinishedCounter++
		panic("test")
	}

	orderWorkerPool := workerpool.NewOrderWorkerPoolImplementation()
	orderGenerator := generator.NewOrderGeneratorImplementation()

	result := orderWorkerPool.StartWorkerPool(ctx, orderGenerator.GenerateOrdersStream(ctx, orders), actions, 5)
	processedOrders := make([]model.Order, 0)

	ticker := time.NewTicker(time.Second * 10)

	func() {
		for {
			select {
			case o, ok := <-result:
				if !ok {
					return
				}

				processedOrders = append(processedOrders, finishedStateToOrder(o))
			case <-ticker.C:
				t.Fail()
			}
		}
	}()

	slices.SortFunc(processedOrders, func(a, b model.Order) int {
		return a.OrderID - b.OrderID
	})

	require.Equal(t, 8, len(processedOrders))

	require.Equal(t, 8, countChecker.initToStartedCounter)
	require.Equal(t, 8, countChecker.StartedToFinishedExternalInteractionCounter)
	require.Equal(t, 5, countChecker.FinishedExternalInteractionToProcessFinishedCounter)
}
