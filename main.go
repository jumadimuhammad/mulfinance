package main

import (
	"log"
	"mulfinance/config"
	authHandler "mulfinance/pkg/auth/handler"
	fileHandler "mulfinance/pkg/file/handler"

	debiturHandler "mulfinance/pkg/debitur/handler"
	debiturRepository "mulfinance/pkg/debitur/repository"
	debiturUsecase "mulfinance/pkg/debitur/usecase"

	limitHandler "mulfinance/pkg/limit/handler"
	limitRepository "mulfinance/pkg/limit/repository"
	limitUsecase "mulfinance/pkg/limit/usecase"

	transactionHandler "mulfinance/pkg/transaction/handler"
	transactionRepository "mulfinance/pkg/transaction/repository"
	transactionUsecase "mulfinance/pkg/transaction/usecase"

	"github.com/joho/godotenv"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	router := config.RouterInitialize()
	db := config.NewDBConnection()

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(authHandler.Claims)
		},
		SigningKey: []byte("secret"),
	}

	auth := router.Group("/auth")
	authHandler := authHandler.NewHandler()
	authHandler.Routes(auth)

	api := router.Group("/api")
	v1 := api.Group("/v1")
	v1.Use(echojwt.WithConfig(config))

	fileHandler.NewHandler().Routes(api)

	debiturRepo := debiturRepository.NewRepository(db)
	limitRepo := limitRepository.NewRepository(db)
	transactionRepo := transactionRepository.NewRepository()

	debiturUsecase := debiturUsecase.NewUsecase(debiturRepo, limitRepo, transactionRepo, db)
	limitUsecase := limitUsecase.NewUsecase(limitRepo)
	transactionUsecase := transactionUsecase.NewUsecase(transactionRepo, debiturRepo, db)

	debiturHandler := debiturHandler.NewHandler(debiturUsecase)
	limitHandler := limitHandler.NewHandler(limitUsecase)
	transactionHandler := transactionHandler.NewHandler(transactionUsecase)

	debiturHandler.Routes(v1)
	limitHandler.Routes(v1)
	transactionHandler.Routes(v1)

	log.Fatal(router.Start(":8000"))
}
