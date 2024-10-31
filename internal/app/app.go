package app

import (
	"log/slog"
	grpcapp "sso/internal/app/grpc"
	"sso/internal/config"
	"sso/internal/services/auth"
	"sso/internal/storage"
	"sso/internal/storage/postgres"
	"time"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	config config.Config,
	tokenTTL time.Duration,
) *App {
	// TODO: init storage (storage)
	db, err := postgres.New(config.DB, log)
	if err != nil {
		log.Error("failed to init storage", slog.String("err", err.Error()))
		return nil
	}

	storageInit, err := storage.NewStorage(db, log)
	if err != nil {
		log.Error("failed to init storage", slog.String("err", err.Error()))
		return nil
	}

	// TODO: init auth service (auth)
	authService := auth.New(log, storageInit, tokenTTL)

	grpcApp := grpcapp.New(log, authService, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
