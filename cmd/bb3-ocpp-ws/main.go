package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Beep-Technologies/beepbeep3-ocpp/api/rest/controller"
	"github.com/Beep-Technologies/beepbeep3-ocpp/api/rest/router"
	ocpp16cs "github.com/Beep-Technologies/beepbeep3-ocpp/internal/ocpp_16/central_system"
	ocppserver "github.com/Beep-Technologies/beepbeep3-ocpp/internal/ocpp_server"
	applicationsrv "github.com/Beep-Technologies/beepbeep3-ocpp/internal/service/application"
	chargepointsrv "github.com/Beep-Technologies/beepbeep3-ocpp/internal/service/charge_point"
	operationsrv "github.com/Beep-Technologies/beepbeep3-ocpp/internal/service/operation"
	statusnotificationsrv "github.com/Beep-Technologies/beepbeep3-ocpp/internal/service/status_notification"
	transactionsrv "github.com/Beep-Technologies/beepbeep3-ocpp/internal/service/transaction"
	usersrv "github.com/Beep-Technologies/beepbeep3-ocpp/internal/service/user"
	userTxsrv "github.com/Beep-Technologies/beepbeep3-ocpp/internal/service/user_transaction"

	"github.com/Beep-Technologies/beepbeep3-ocpp/pkg/logger"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

// @title BB3 OCPP API
// @version 2.0
// @description Service to interface with OCPP-compliant charge points

// @contact.name Lowen
// @contact.email lowen@beepbeep.tech

// @host ocpp-dev.chargegowhere.sg:8060
// @schemes http

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// load from .env file. doesn't matter if an error is returned
	godotenv.Load()

	// set up database
	dbhost := os.Getenv("DB_HOST")
	dbname := os.Getenv("DB_NAME")
	dbuser := os.Getenv("DB_USER")
	dbpassword := os.Getenv("DB_PASSWORD")
	dbsslmode := os.Getenv("DB_SSLMODE")

	dbString := fmt.Sprintf("host=%s port=5432 user=%s dbname=%s password=%s", dbhost, dbuser, dbname, dbpassword)
	if dbsslmode != "" {
		dbString += " sslmode=" + dbsslmode
	}

	ORM, err := gorm.Open(
		postgres.New(postgres.Config{
			DSN:                  dbString,
			PreferSimpleProtocol: true,
		}),
		&gorm.Config{
			Logger: gormLogger.Default.LogMode(gormLogger.Silent),
		})
	if err != nil {
		fmt.Println(err)
	}

	sqlDB, _ := ORM.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	var l *zap.Logger
	logger.LogInit()
	if os.Getenv("ENVIRONMENT") == "production" {
		gin.SetMode(gin.ReleaseMode)
		l, _ = zap.NewProduction()
	} else {
		l, _ = zap.NewDevelopment()
	}
	defer l.Sync()

	applicationService := applicationsrv.NewService(ORM)
	chargePointService := chargepointsrv.NewService(ORM)
	transactionService := transactionsrv.NewService(ORM)
	statusNotificationService := statusnotificationsrv.NewService(ORM)
	userService := usersrv.NewService(ORM)
	userTxService := userTxsrv.NewService(ORM)

	ocpp16CentralSystem := ocpp16cs.NewOCPP16CentralSystem(
		l,
		applicationService,
		chargePointService,
		transactionService,
		statusNotificationService,
	)

	operationService := operationsrv.NewService(
		ORM,
		chargePointService,
		transactionService,
		ocpp16CentralSystem,
	)

	ocppWebSocketServer := ocppserver.NewOCPPWebSocketServer(l, ocpp16CentralSystem)
	chargePointController := controller.NewChargePointsAPI(chargePointService, statusNotificationService)
	operationController := controller.NewOperationsAPI(
		operationService,
		statusNotificationService,
	)

	transactionController := controller.NewTransactionAPI(
		chargePointService, statusNotificationService, transactionService,
	)

	userController := controller.NewUserAPI(
		userService,
	)

	userTxController := controller.NewUserTxAPI(
		userTxService,
	)

	rt := router.NewRouter(
		ocppWebSocketServer,
		chargePointController,
		operationController,
		transactionController,
		userController,
		userTxController,
	)

	// set up gin
	r := gin.New()
	// use ginzap middleware
	gl := l.With(zap.String("source", "gin"))
	r.Use(ginzap.Ginzap(gl, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(l, true))
	// apply router
	rt.Apply(r)

	host := "0.0.0.0"
	port := 8060
	if hostPortStr := os.Getenv("HOST_PORT"); hostPortStr != "" {
		port, _ = strconv.Atoi(hostPortStr)
	}

	addr := fmt.Sprintf("%s:%d", host, port)

	s := &http.Server{
		Addr:         addr,
		Handler:      r,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}

	l.Info("Starting server at " + addr + "...")
	if err := s.ListenAndServe(); err != nil {
		l.Error("Server error", zap.String("error", err.Error()))
	}
}
