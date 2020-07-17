package ecom

import (
	"strconv"
	"time"

	"go.temporal.io/temporal/workflow"
	"go.uber.org/zap"
)

var (
	userServiceServerHostPort = "http://localhost:4000"
	productServiceHostPort    = "http://localhost:4001"
	offerServiceHostPort      = "http://localhost:4002"
	paymentHostPort           = "http://localhost:4003"
)

func EcomWorkflow(ctx workflow.Context, payload *Payload) (*Output, error) {
	// step 1, Get user details
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx1 := workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)

	var customer User
	err := workflow.ExecuteActivity(ctx1, GetUserActivity, payload.CustomerID).Get(ctx1, &customer)
	if err != nil {
		logger.Error("Failed to get user", zap.Error(err))
		return nil, err
	}

	var products []Product

	for i := 0; i < len(payload.Products); i++ {
		future, settable := workflow.NewFuture(ctx1)
		workflow.Go(ctx1, func(ctx workflow.Context) {
			defer logger.Info("Second goroutine completed.")

			var result Product
			err := workflow.ExecuteActivity(ctx, GetProductActivitiy, payload.Products[i]).Get(ctx, &result)
			settable.Set(result, err)
		})

		var prodResult Product
		err = future.Get(ctx1, &prodResult)
		if err != nil {
			return nil, err
		}
		products = append(products, prodResult)
	}

	for i := 0; i < len(payload.Offers); i++ {
		future, settable := workflow.NewFuture(ctx1)
		workflow.Go(ctx1, func(ctx workflow.Context) {
			defer logger.Info("Second goroutine completed.")

			var result Price
			err := workflow.ExecuteActivity(ctx, DiscountActivitiy, strconv.Itoa(products[i].Cost)).Get(ctx, &result)
			settable.Set(result, err)
		})

		var disount Price
		err = future.Get(ctx1, &disount)
		if err != nil {
			return nil, err
		}
		products[i].Discount = disount.DiscountedPrice
	}

	ao = workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx2 := workflow.WithActivityOptions(ctx1, ao)

	var paymentResult Pay
	err = workflow.ExecuteActivity(ctx2, PaymentActivitiy).Get(ctx2, &paymentResult)
	if err != nil {
		logger.Error("Failed to get payment", zap.Error(err))
		return nil, err
	}

	finalOut := &Output{
		Customer: customer,
		Products: products,
		Payment:  paymentResult.Status,
		Shipping: paymentResult.Status,
		ID:       paymentResult.ID,
	}

	logger.Info("Ecommerce Workflow completed.")
	return finalOut, nil
}
