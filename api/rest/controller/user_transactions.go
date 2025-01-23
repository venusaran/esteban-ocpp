package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Beep-Technologies/beepbeep3-ocpp/api/rpc"
	userTx "github.com/Beep-Technologies/beepbeep3-ocpp/internal/service/user_transaction"
	"github.com/Beep-Technologies/beepbeep3-ocpp/pkg/constants"
	"github.com/Beep-Technologies/beepbeep3-ocpp/pkg/logger"
)

type UserTxAPI struct {
	userTxService *userTx.Service
}

func NewUserTxAPI(utS *userTx.Service) *UserTxAPI {
	return &UserTxAPI{
		userTxService: utS,
	}
}

func (api *UserTxAPI) CreateUserTx(c *gin.Context) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, constants.CtxKey("gin"), c)

	req := &rpc.CreateUserTransactionReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	resp, err := api.userTxService.CreateUserTransaction(ctx, req)
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

func (api *UserTxAPI) UpdateUserTx(c *gin.Context) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, constants.CtxKey("gin"), c)

	txId := c.Param("user_tx_td")
	req := &rpc.UpdateUserTransactionReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	iTxID, _ := strconv.ParseInt(txId, 10, 32)
	if int32(iTxID) != req.Transaction.Id {
		logger.Error("user transaction id mismatch in queryparam and body")
		c.Error(err).SetType(gin.ErrorTypeAny)
		return
	}

	resp, err := api.userTxService.UpdateUserTransaction(ctx, req)
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

func (api *UserTxAPI) GetAllUserTxs(c *gin.Context) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, constants.CtxKey("gin"), c)

	req := &rpc.GetUserTransactionReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	resp, err := api.userTxService.GetAllUserTransaction(ctx, req)
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
