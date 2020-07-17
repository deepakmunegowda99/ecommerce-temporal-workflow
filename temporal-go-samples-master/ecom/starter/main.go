package main

import (
	"context"
	"strconv"

	"go.temporal.io/temporal/client"
	"go.uber.org/zap"

	"github.com/temporalio/temporal-go-samples/ecom"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	// The client is a heavyweight object that should be created once per process.
	c, err := client.NewClient(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		logger.Fatal("Unable to create client", zap.Error(err))
	}
	defer c.Close()

	for i := 0; i < 1; i++ {

		workflowOptions := client.StartWorkflowOptions{
			ID:        "ecom_" + strconv.Itoa(i),
			TaskQueue: "ecom",
		}

		var payload ecom.Payload
		payload.CustomerID = "user_" + strconv.Itoa(i)
		payload.Products = []string{"prod_" + strconv.Itoa(i+1), "prod_" + strconv.Itoa(i+12), "prod_" + strconv.Itoa(i+13)}
		payload.Offers = []string{"coupon_" + strconv.Itoa(i+1), "coupon_" + strconv.Itoa(i+2)}
		payload.PaymentMethod = "CCV"

		_, err = c.ExecuteWorkflow(context.Background(), workflowOptions, ecom.EcomWorkflow, &payload)
		if err != nil {
			logger.Fatal("Unable to execute workflow", zap.Error(err))
		}

	}

	// logger.Info("Started workflow", zap.String("WorkflowID", we.GetID()), zap.String("RunID", we.GetRunID()))

}
