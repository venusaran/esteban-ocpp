package usertransaction

import (
	"context"

	"gorm.io/gorm"

	"github.com/Beep-Technologies/beepbeep3-ocpp/api/rpc"
	"github.com/Beep-Technologies/beepbeep3-ocpp/internal/models"
	chargepoint "github.com/Beep-Technologies/beepbeep3-ocpp/internal/repository/charge_point"
	transaction "github.com/Beep-Technologies/beepbeep3-ocpp/internal/repository/transaction"
	user "github.com/Beep-Technologies/beepbeep3-ocpp/internal/repository/user"
	userTransaction "github.com/Beep-Technologies/beepbeep3-ocpp/internal/repository/user_transaction"
	"github.com/Beep-Technologies/beepbeep3-ocpp/pkg/logger"
)

const RFC3339Milli = "2006-01-02T15:04:05.000Z07:00"

type Service struct {
	db              *gorm.DB
	transaction     transaction.BaseRepo
	chargePoint     chargepoint.BaseRepo
	user            user.BaseRepo
	userTransaction userTransaction.BaseRepo
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		db:              db,
		chargePoint:     chargepoint.NewBaseRepo(db),
		transaction:     transaction.NewBaseRepo(db),
		user:            user.NewBaseRepo(db),
		userTransaction: userTransaction.NewBaseRepo(db),
	}
}

func (srv Service) CreateUserTransaction(ctx context.Context, req *rpc.CreateUserTransactionReq) (resp *rpc.CreateUserTransactionResp, err error) {
	ut := models.UserTransaction{
		UserID:        req.Transaction.UserId,
		TransactionID: req.Transaction.TransactionId,
	}

	err = srv.userTransaction.CreateUserTransaction(&ut)
	if err != nil {
		logger.Errorw("UserTransaction creation failed", "err", err.Error())
		return nil, err
	}

	logger.Infow("UserTransaction Created Successfully!")

	resp = &rpc.CreateUserTransactionResp{
		IsSuccess: true,
	}

	return resp, nil
}

func (srv Service) UpdateUserTransaction(ctx context.Context, req *rpc.UpdateUserTransactionReq) (resp *rpc.UpdateUserTransactionResp, err error) {
	ut := models.UserTransaction{
		ID:            req.Transaction.Id,
		UserID:        req.Transaction.UserId,
		TransactionID: req.Transaction.TransactionId,
	}

	err = srv.userTransaction.UpdateUserTransaction(&ut)
	if err != nil {
		logger.Errorw("UserTransaction update failed", "err", err.Error())
		return nil, err
	}

	logger.Infow("UserTransaction Updated Successfully!")

	resp = &rpc.UpdateUserTransactionResp{
		IsSuccess: true,
	}

	return resp, nil
}

func (srv Service) GetAllUserTransaction(ctx context.Context, req *rpc.GetUserTransactionReq) (resp *rpc.GetUserTransactionResp, err error) {
	txs, err := srv.userTransaction.GetUserTransactionByUserId(req.UserId)
	if err != nil {
		logger.Errorw("UserTransaction fetching failed", "err", err.Error())
		return nil, err
	}

	logger.Infow("Get UserTransaction Success!")

	tss := []*rpc.UserTransaction{}
	for _, t := range txs {
		ts := rpc.UserTransaction{
			Id:            t.ID,
			UserId:        t.UserID,
			TransactionId: t.TransactionID,
		}
		tss = append(tss, &ts)
	}

	resp = &rpc.GetUserTransactionResp{
		Transactions: tss,
	}

	return resp, nil
}
