package main

import (
	"fmt"

	"github.com/jamm3e3333/quiz-app/config"
	"github.com/jamm3e3333/quiz-app/grpc"
	"github.com/jamm3e3333/quiz-app/logger"
	"github.com/jamm3e3333/quiz-app/shutdown"
)

func main() {
	ctx := shutdown.SetupShutdownContext()
	cfg, err := config.NewParseConfigForENV()
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	lg := logger.New(logger.ParseLevel(cfg.Logger.Level), cfg.Logger.ShouldUseDevelMode)

	grpcServer := grpc.NewServer(
		lg,
		uint32(cfg.GRPCServer.Port),
		cfg.GRPCServer.ShouldUseReflection,
		grpc.NewLogInterceptor(lg),
	)
	RegisterModule(grpcServer)

	errChan := grpcServer.Run()
	lg.Info(fmt.Sprintf("Server is running on port %d", cfg.GRPCServer.Port))

	select {
	case <-ctx.Done():
		lg.Info("Shutting down server...")
	case err := <-errChan:
		lg.Error("Error running server: %v", err)
		shutdown.SignalShutdown()
	}
}
