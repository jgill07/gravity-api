package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jgill07/gravity-api/internal/dto"
	"github.com/jgill07/gravity-api/internal/service"
)

func healthz(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
	})
}

func createTransaction(svc *service.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var model dto.TransactionIn
		if err := ctx.ShouldBindJSON(&model); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := svc.CreateTransaction(ctx, model); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, gin.H{})
	}
}

func getTransactions(svc *service.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		transactions, err := svc.GetTransactions(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, transactions)
	}
}
