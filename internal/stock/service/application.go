package service

import (
	"context"

	"github.com/miRceLzeS/gorder-v2/common/metrics"
	"github.com/miRceLzeS/gorder-v2/stock/adapters"
	"github.com/miRceLzeS/gorder-v2/stock/app"
	"github.com/miRceLzeS/gorder-v2/stock/app/query"
	"github.com/miRceLzeS/gorder-v2/stock/infrastructure/integration"
	"github.com/miRceLzeS/gorder-v2/stock/infrastructure/persistent"
	"github.com/sirupsen/logrus"
)

func NewApplication(_ context.Context) app.Application {
	//stockRepo := adapters.NewMemoryStockRepository()
	db := persistent.NewMySQL()
	stockRepo := adapters.NewMySQLStockRepository(db)
	logger := logrus.NewEntry(logrus.StandardLogger())
	stripeAPI := integration.NewStripeAPI()
	metricsClient := metrics.TodoMetrics{}
	return app.Application{
		Commands: app.Commands{},
		Queries: app.Queries{
			CheckIfItemsInStock: query.NewCheckIfItemsInStockHandler(stockRepo, stripeAPI, logger, metricsClient),
			GetItems:            query.NewGetItemsHandler(stockRepo, logger, metricsClient),
		},
	}
}
