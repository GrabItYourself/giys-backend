package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/GrabItYourself/giys-backend/apigateway/internal/config"
	"github.com/GrabItYourself/giys-backend/apigateway/internal/v1router"
	v1handler "github.com/GrabItYourself/giys-backend/apigateway/internal/v1router/handler"
	"github.com/GrabItYourself/giys-backend/lib/logger"
	"github.com/GrabItYourself/giys-backend/user/pkg/client"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Initialize context
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	// Load config from YAML file
	conf := config.InitConfig()

	// Initialize logger
	logger.InitLogger(&conf.Log)

	// Initialize GRPC client
	userGrpcClient, err := client.NewClient(ctx, conf.Grpc.User.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Panic(err.Error())
	}

	// Initialize fiber app
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	// Create Routing group
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// Handle API v1 routes
	v1Handler := v1handler.NewHandler(userGrpcClient)
	v1Router := v1router.NewRouter(ctx, v1, v1Handler)
	v1Router.InitUserRoute(ctx, "/user")

	// Graceful shutdown for fiber app
	go func() {
		<-ctx.Done()
		logger.Info("Gracefully shutting down...")
		app.Shutdown()
	}()

	// Start the server
	if err := app.Listen(":" + conf.Server.Port); err != nil {
		logger.Panic(errors.Wrap(err, "Failed to start server").Error())
	}

}
