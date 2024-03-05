package pipeline

import (
	"context"
	"errors"
	"hw7/internal/model"
	"sync"
)

const fanLimit = 10

type OrderPipeline interface {
	Start(
		ctx context.Context,
		actions model.OrderActions,
		orders <-chan model.OrderInitialized,
		processed chan<- model.OrderProcessFinished,
	)
}

type OrderPipelineImplementation struct{}

func NewOrderPipelineImplementation() *OrderPipelineImplementation {
	return &OrderPipelineImplementation{}
}

func (o *OrderPipelineImplementation) Start(
	ctx context.Context,
	actions model.OrderActions,
	orders <-chan model.OrderInitialized,
	processed chan model.OrderProcessFinished,
) {

	chanProcessStarted := make(chan model.OrderProcessStarted)
	chanFinishedExternalInteraction := make(chan model.OrderFinishedExternalInteraction)

	go func() {
		defer close(chanProcessStarted)
		InitToStarted(actions, orders, chanProcessStarted, processed)
	}()

	go func() {
		fanOutProgress := make([]chan model.OrderFinishedExternalInteraction, fanLimit)

		for i := 0; i < fanLimit; i++ {
			fanOutProgress[i] = make(chan model.OrderFinishedExternalInteraction)
		}

		for it := 0; it < fanLimit; it++ {
			go func(it int) {
				defer close(fanOutProgress[it])
				StartedToFinishedExternalInteraction(
					actions,
					chanProcessStarted,
					fanOutProgress[it],
					processed,
				)
			}(it)
		}

		var wg sync.WaitGroup
		wg.Add(fanLimit)
		for _, c := range fanOutProgress {
			go func(c chan model.OrderFinishedExternalInteraction) {
				for v := range c {
					chanFinishedExternalInteraction <- v
				}
				wg.Done()
			}(c)
		}
		go func() {
			wg.Wait()
			close(chanFinishedExternalInteraction)
		}()
	}()

	func() {
		FinishedExternalInteractionToProcessFinished(
			actions,
			chanFinishedExternalInteraction,
			processed,
		)
	}()
}

func InitToStarted(
	actions model.OrderActions,
	orders <-chan model.OrderInitialized,
	updateOrders chan model.OrderProcessStarted,
	processed chan model.OrderProcessFinished,
) {
	reserveOrders := make(chan model.OrderInitialized)
	wg := sync.WaitGroup{}

	defer func() {
		go func() {
			wg.Wait()
			close(reserveOrders)
		}()

		if err := recover(); err != nil {
			updateOrders = make(chan model.OrderProcessStarted)
			makeProcessFinishedFromInit(reserveOrders, processed, "panic: panic in InitToStart()")
			makeProcessFinishedFromInit(orders, processed, "panic: panic in InitToStart()")
		}
	}()

	for order := range orders {
		wg.Add(1)
		go func() {
			reserveOrders <- order
			wg.Done()
		}()
		actions.InitToStarted()
		order.OrderStates = append(order.OrderStates, model.ProcessStarted)
		updateOrders <- model.OrderProcessStarted{
			OrderInitialized: order,
			OrderStates:      order.OrderStates,
			Error:            order.Error,
		}
	}
}

func StartedToFinishedExternalInteraction(
	actions model.OrderActions,
	orders <-chan model.OrderProcessStarted,
	updateOrders chan model.OrderFinishedExternalInteraction,
	processed chan model.OrderProcessFinished,
) {
	reserveOrders := make(chan model.OrderProcessStarted)
	wg := sync.WaitGroup{}

	defer func() {
		go func() {
			wg.Wait()
			close(reserveOrders)
		}()

		if err := recover(); err != nil {
			updateOrders = make(chan model.OrderFinishedExternalInteraction)
			go func() {
				makeProcessFinishedFromStarted(
					reserveOrders,
					processed,
					"panic: panic in StartedToFinishedExternalInteraction()",
				)
				makeProcessFinishedFromStarted(
					orders,
					processed,
					"panic: panic in StartedToFinishedExternalInteraction()",
				)
			}()
		}
	}()

	for order := range orders {
		wg.Add(1)
		go func() {
			reserveOrders <- order
			wg.Done()
		}()
		actions.StartedToFinishedExternalInteraction()
		order.OrderStates = append(order.OrderStates, model.FinishedExternalInteraction)
		updateOrders <- model.OrderFinishedExternalInteraction{
			OrderProcessStarted: order,
			StorageID:           order.OrderInitialized.ProductID%2 + 1,
			PickupPointID:       order.OrderInitialized.ProductID%3 + 1,
			OrderStates:         order.OrderStates,
			Error:               order.Error,
		}
	}
}

func FinishedExternalInteractionToProcessFinished(
	actions model.OrderActions,
	orders <-chan model.OrderFinishedExternalInteraction,
	processed chan model.OrderProcessFinished,
) {
	reserveOrders := make(chan model.OrderFinishedExternalInteraction)
	wg := sync.WaitGroup{}

	defer func() {
		go func() {
			wg.Wait()
			close(reserveOrders)
		}()

		if err := recover(); err != nil {
			makeProcessFinishedFromFinishedExternalInteraction(
				reserveOrders,
				processed,
				"panic: panic in FinishedExternalInteractionToProcessFinished()",
			)

			makeProcessFinishedFromFinishedExternalInteraction(
				orders,
				processed,
				"panic: panic in FinishedExternalInteractionToProcessFinished()",
			)
		}
	}()

	for order := range orders {
		wg.Add(1)
		go func() {
			reserveOrders <- order
			wg.Done()
		}()
		actions.FinishedExternalInteractionToProcessFinished()
		order.OrderStates = append(order.OrderStates, model.ProcessFinished)
		processed <- model.OrderProcessFinished{
			OrderFinishedExternalInteraction: order,
			OrderStates:                      order.OrderStates,
			Error:                            order.Error,
		}
	}
}

func makeProcessFinishedFromInit(
	orders <-chan model.OrderInitialized,
	processed chan model.OrderProcessFinished,
	str string) {
	for order := range orders {
		fei := model.OrderProcessStarted{
			OrderInitialized: order,
			OrderStates:      order.OrderStates,
			Error:            order.Error,
		}
		processed <- model.OrderProcessFinished{
			OrderFinishedExternalInteraction: model.OrderFinishedExternalInteraction{
				OrderProcessStarted: fei,
				OrderStates:         fei.OrderStates,
				Error:               fei.Error,
			},
			OrderStates: order.OrderStates,
			Error:       errors.New(str),
		}
	}
}

func makeProcessFinishedFromStarted(
	orders <-chan model.OrderProcessStarted, processed chan model.OrderProcessFinished, str string) {
	for order := range orders {
		processed <- model.OrderProcessFinished{
			OrderFinishedExternalInteraction: model.OrderFinishedExternalInteraction{
				OrderProcessStarted: order,
				OrderStates:         order.OrderStates,
				Error:               order.Error,
			},
			OrderStates: order.OrderStates,
			Error:       errors.New(str),
		}
	}
}

func makeProcessFinishedFromFinishedExternalInteraction(
	orders <-chan model.OrderFinishedExternalInteraction,
	processed chan model.OrderProcessFinished,
	str string,
) {
	for order := range orders {
		processed <- model.OrderProcessFinished{
			OrderFinishedExternalInteraction: order,
			OrderStates:                      order.OrderStates,
			Error:                            errors.New(str),
		}
	}
}
