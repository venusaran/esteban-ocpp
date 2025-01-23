package user

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"github.com/Beep-Technologies/beepbeep3-ocpp/api/rpc"
	"github.com/Beep-Technologies/beepbeep3-ocpp/internal/models"
	chargepoint "github.com/Beep-Technologies/beepbeep3-ocpp/internal/repository/charge_point"
	transaction "github.com/Beep-Technologies/beepbeep3-ocpp/internal/repository/transaction"
	user "github.com/Beep-Technologies/beepbeep3-ocpp/internal/repository/user"
	"github.com/Beep-Technologies/beepbeep3-ocpp/pkg/logger"
	"github.com/google/uuid"
)

const RFC3339Milli = "2006-01-02T15:04:05.000Z07:00"

type Service struct {
	db          *gorm.DB
	transaction transaction.BaseRepo
	chargePoint chargepoint.BaseRepo
	user        user.BaseRepo
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		db:          db,
		chargePoint: chargepoint.NewBaseRepo(db),
		transaction: transaction.NewBaseRepo(db),
		user:        user.NewBaseRepo(db),
	}
}

func (srv Service) RegisterUser(ctx context.Context, req *rpc.RegisterUserReq) (resp *rpc.RegisterUserResp, err error) {
	apiKey := uuid.New()

	hashedAPIKey, err := HashAPIKey(apiKey.String())
	if err != nil {
		logger.Errorw("hashing api key failed", "err", err.Error())
		return
	}

	ep, err := EncryptPw(req.Password)
	if err != nil {
		logger.Errorw("encrypting password failed", "err", err.Error())
		return
	}

	user := models.User{
		Username:     req.Username,
		Email:        req.Email,
		HashedAPIKey: hashedAPIKey,
		Password:     ep,
	}

	err = srv.user.CreateUser(&user)
	if err != nil {
		logger.Errorw("User creation failed", "err", err.Error())
		return
	}

	logger.Infow("User Created Successfully!")

	resp = &rpc.RegisterUserResp{
		ApiKey: apiKey.String(),
	}

	return
}

func (srv Service) VerifyUser(ctx context.Context, req *rpc.VerifyUserReq) (resp *rpc.VerifyUserResp, err error) {
	hashedAPIKey, err := HashAPIKey(req.ApiKey)
	if err != nil {
		logger.Errorw("hashing password failed", "err", err.Error())
		return
	}

	// get user
	user, err := srv.user.GetUserByHashedApiKey(hashedAPIKey)
	if err != nil {
		return nil, err
	}

	u := rpc.UserOut{
		Id:       user.ID,
		UserName: user.Username,
		Email:    user.Email,
		IsAdmin:  user.IsAdmin,
	}

	resp = &rpc.VerifyUserResp{
		User: &u,
	}

	return resp, nil
}

func (srv Service) LoginUser(ctx context.Context, req *rpc.LoginReq) (resp *rpc.LoginResp, err error) {
	// get user
	user, err := srv.user.GetUserByUsername(req.Username)
	if err != nil {
		return nil, err
	}
	isSuccess, loggedInUser := VerifyPw(req.Username, req.Password, *user)
	if !isSuccess {
		err = fmt.Errorf("error validating password")
		logger.Errorw("error validating password", "err", err.Error())
		return nil, err
	}

	ou := &rpc.UserOut{
		Id:       loggedInUser.ID,
		UserName: loggedInUser.Username,
		Email:    loggedInUser.Email,
		IsAdmin:  loggedInUser.IsAdmin,
	}

	resp = &rpc.LoginResp{
		User: ou,
	}

	return resp, nil
}
