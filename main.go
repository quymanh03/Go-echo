package main

import (
	"beginner/config"
	"beginner/handler"
	"beginner/router"

	_ "beginner/docs"

	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description Beginner API Server
// @title Beginner API Server

// @host 127.0.0.1:3000
// @BasePath /api/v1

// @schemes http https
// @produce	application/json
// @consumes application/json

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	r := router.New()
	r.GET("/docs/*", echoSwagger.WrapHandler)
	v1 := r.Group("/api/v1")

	// Connect Database
	envConfig, err := config.LoadConfig()
	if err != nil {
		r.Logger.Fatal("Failed to load config:", err)
	}
	err = config.ConnectDatabase(&envConfig.DB)
	if err != nil {
		r.Logger.Fatal("Failed to connect to database:", err)
	}

	gorm := config.DB()

	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}

	err = dbGorm.Ping()
	if err != nil {
		r.Logger.Fatal("Failed to ping database:", err)
	}
	r.Logger.Info("Database connected")

	h := handler.New(gorm)
	h.Register(v1)
	// Start the server
	r.Logger.Fatal(r.Start(":3000"))
}
