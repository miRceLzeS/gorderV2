package service

import (
	"context"

	grpcClient "github.com/miRceLzeS/gorder-v2/common/client"
	"github.com/miRceLzeS/gorder-v2/common/metrics"
	"github.com/miRceLzeS/gorder-v2/payment/adapters"
	"github.com/miRceLzeS/gorder-v2/payment/app"
	"github.com/miRceLzeS/gorder-v2/payment/app/command"
	"github.com/miRceLzeS/gorder-v2/payment/domain"
	"github.com/miRceLzeS/gorder-v2/payment/infrastructure/processor"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewApplication(ctx context.Context) (app.Application, func()) {
	orderClient, closeOrderClient, err := grpcClient.NewOrderGRPCClient(ctx)
	if err != nil {
		panic(err)
	}
	orderGRPC := adapters.NewOrderGRPC(orderClient)
	//memoryProcessor := processor.NewInmemProcessor()
	stripeProcessor := processor.NewStripeProcessor(viper.GetString("stripe-key"))
	return newApplication(ctx, orderGRPC, stripeProcessor), func() {
		_ = closeOrderClient()
	}
}

func newApplication(_ context.Context, orderGRPC command.OrderService, processor domain.Processor) app.Application {
	logger := logrus.NewEntry(logrus.StandardLogger())
	metricClient := metrics.TodoMetrics{}
	return app.Application{
		Commands: app.Commands{
			CreatePayment: command.NewCreatePaymentHandler(processor, orderGRPC, logger, metricClient),
		},
	}
}
