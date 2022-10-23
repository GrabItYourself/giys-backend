package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/GrabItYourself/giys-backend/apigateway/internal/config"
	"github.com/GrabItYourself/giys-backend/apigateway/internal/v1router"
	v1handler "github.com/GrabItYourself/giys-backend/apigateway/internal/v1router/handler"
	authclient "github.com/GrabItYourself/giys-backend/auth/pkg/client"
	"github.com/GrabItYourself/giys-backend/lib/logger"
	shopclient "github.com/GrabItYourself/giys-backend/shop/pkg/shopclient"
	userclient "github.com/GrabItYourself/giys-backend/user/pkg/client"
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
	userGrpcClient, userGrpcConn, err := userclient.NewClient(ctx, conf.Grpc.User.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Panic(errors.Wrap(err, "Failed to initialize user GRPC client").Error())
	}
	defer func() {
		logger.Info("Closing user GRPC connection...")
		if err := userGrpcConn.Close(); err != nil {
			logger.Panic(errors.Wrap(err, "Failed to close user GRPC connection").Error())
		}
	}()
	authGrpcClient, authGrpcConn, err := authclient.NewClient(ctx, conf.Grpc.Auth.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Panic(errors.Wrap(err, "Failed to initialize auth GRPC client").Error())
	}
	defer func() {
		logger.Info("Closing auth GRPC connection...")
		if err := authGrpcConn.Close(); err != nil {
			logger.Panic(errors.Wrap(err, "Failed to close auth GRPC connection").Error())
		}
	}()
	shopGrpcClient, shopGrpcConn, err := shopclient.NewClient(ctx, conf.Grpc.Auth.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Panic(errors.Wrap(err, "Failed to initialize shop GRPC client").Error())
	}
	defer func() {
		logger.Info("Closing shop GRPC connection...")
		if err := shopGrpcConn.Close(); err != nil {
			logger.Panic(errors.Wrap(err, "Failed to close shop GRPC connection").Error())
		}
	}()

	grpcClients := &v1handler.GrpcClients{
		User: userGrpcClient,
		Auth: authGrpcClient,
		Shop: shopGrpcClient,
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
	v1Handler := v1handler.NewHandler(grpcClients, &conf.OAuth)
	v1Router := v1router.NewRouter(ctx, v1, v1Handler)
	v1Router.InitUserRoutes("/user")
	v1Router.InitAuthRoutes("/auth")
	v1Router.InitShopRoutes("/shop")

	// Graceful shutdown for fiber app
	go func() {
		<-ctx.Done()
		logger.Info("Gracefully shutting down...")
		err := app.Shutdown()
		if err != nil {
			logger.Error(errors.Wrap(err, "Failed to shutdown fiber app").Error())
		}
	}()

	// Start the server
	if err := app.Listen(":" + conf.Server.Port); err != nil {
		logger.Panic(errors.Wrap(err, "Server returned with error").Error())
	}
}
