package app

import (
	grpcapp "grpc/auth/internal/app/grpc"
	"grpc/auth/internal/services/auth"
	"grpc/auth/internal/storage/postgres"
	"log/slog"
	"time"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(log *slog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {

	storage, err := postgres.New(storagePath)
	if err != nil {
		panic(err)
	}

	authService := auth.New(log, storage, storage, storage, tokenTTL)

	grpcApp := grpcapp.New(log, grpcPort, authService)
	return &App{
		GRPCSrv: grpcApp,
	}

}
