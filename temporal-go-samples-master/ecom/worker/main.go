package main

import (
	"go.temporal.io/temporal/client"
	"go.temporal.io/temporal/worker"
	"go.uber.org/zap"

	"github.com/temporalio/temporal-go-samples/ecom"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	// The client and worker are heavyweight objects that should be created once per process.
	c, err := client.NewClient(client.Options{
		HostPort: client.DefaultHostPort,
		Logger:   logger,
	})
	if err != nil {
		logger.Fatal("Unable to create client", zap.Error(err))
	}
	defer c.Close()

	w := worker.New(c, "ecom", worker.Options{})

	w.RegisterWorkflow(ecom.EcomWorkflow)
	w.RegisterActivity(ecom.GetUserActivity)
	w.RegisterActivity(ecom.GetProductActivitiy)
	w.RegisterActivity(ecom.DiscountActivitiy)
	w.RegisterActivity(ecom.PaymentActivitiy)

	err = w.Run()
	if err != nil {
		logger.Fatal("Unable to start worker", zap.Error(err))
	}
}
