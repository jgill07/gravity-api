package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jgill07/gravity-api/internal/log"
	"github.com/jgill07/gravity-api/internal/models"
	"github.com/jgill07/gravity-api/internal/service"
)

func getPluralServiceName(svcType string) string {
	switch svcType {
	case string(models.TransactionIncome):
		return "incomes"
	case string(models.TransactionExpense):
		return "expenses"
	default:
		log.Error("Unknown service type")
		panic("Unknown service type")
	}
}

func SetupRouter(svc *service.Service) *gin.Engine {
	router := initRouter()
	router.GET("/healthz", healthz)

	tRouter := router.Group(fmt.Sprintf("/%s", getPluralServiceName(svc.Config.ApiConfig.Service)))
	tRouter.POST("", createTransaction(svc))
	tRouter.GET("", getTransactions(svc))

	return router
}

func initRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.ContextWithFallback = true
	router.Use(logger)
	return router
}
