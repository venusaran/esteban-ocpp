package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Beep-Technologies/beepbeep3-ocpp/api/rpc"
	chargepoint "github.com/Beep-Technologies/beepbeep3-ocpp/internal/service/charge_point"
	statusnotification "github.com/Beep-Technologies/beepbeep3-ocpp/internal/service/status_notification"
	transaction "github.com/Beep-Technologies/beepbeep3-ocpp/internal/service/transaction"
	"github.com/Beep-Technologies/beepbeep3-ocpp/pkg/constants"
)

type TransactionAPI struct {
	chargepointService        *chargepoint.Service
	statusnotificationService *statusnotification.Service
	transactionService        *transaction.Service
}

func NewTransactionAPI(cpS *chargepoint.Service, snS *statusnotification.Service, ts *transaction.Service) *TransactionAPI {
	return &TransactionAPI{
		chargepointService:        cpS,
		statusnotificationService: snS,
		transactionService:        ts,
	}
}

// GetTransactions gets all the transactions
func (api *TransactionAPI) GetAllTransactions(c *gin.Context) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, constants.CtxKey("gin"), c)

	limit := c.Param("limit")
	offset := c.Param("offset")

	iLimit, _ := strconv.Atoi(limit)
	iOffset, _ := strconv.Atoi(offset)

	req := &rpc.GetTransactionsReq{
		Limit:  int32(iLimit),
		Offset: int32(iOffset),
	}

	resp, err := api.transactionService.GetTransactions(ctx, req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success",
		"data":    resp,
	})
}

// GetTransactions gets all the transactions
func (api *TransactionAPI) GetSingleTransaction(c *gin.Context) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, constants.CtxKey("gin"), c)

	id := c.Param("id")

	iID, _ := strconv.Atoi(id)

	req := &rpc.GetTransactionReq{
		Id: int32(iID),
	}

	resp, err := api.transactionService.GetSingleTransaction(ctx, req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success",
		"data":    resp,
	})
}

// GetChargePointTransactions gets all the transactions of a charge point
func (api *TransactionAPI) GetChargePointTransactions(c *gin.Context) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, constants.CtxKey("gin"), c)

	cpId := c.Param("charge_point_id")

	cpIdInt, _ := strconv.Atoi(cpId)
	resp, err := api.transactionService.GetChargePointTransactions(ctx, int32(cpIdInt))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success",
		"data":    resp,
	})
}
