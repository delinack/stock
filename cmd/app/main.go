package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/delinack/stock/internal/config"
	"github.com/delinack/stock/internal/domain"
	"github.com/delinack/stock/internal/pkg/logger"
	"github.com/delinack/stock/internal/pkg/server"
	"github.com/delinack/stock/internal/pkg/service"
	"github.com/delinack/stock/internal/pkg/storage"
	"github.com/delinack/stock/internal/pkg/storage/item_storage"
	"github.com/delinack/stock/internal/pkg/storage/stock_storage"
	"github.com/rs/zerolog/log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	applicationConfig := config.MustParseConfig()
	log.Logger = logger.NewLogger(applicationConfig.LoggerConfig)
	log.Info().Str("comp", "main").Msg("starting application...")

	db, err := storage.NewPGConnection(ctx, applicationConfig.PGConfig)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot open DB connection")
	}

	log.Info().Msg("connection to DB was established")
	defer func() {
		err = db.Close()
		if err != nil {
			log.Error().Err(err).Msg("error while closing DB connection")
		}
		log.Info().Msg("connection to DB was closed")
	}()

	newTemplateRepository := stock_storage.NewStockRepository(db)
	newParameterRepository := item_storage.NewItemRepository(db)
	store := storage.NewStorage(newTemplateRepository, newParameterRepository)
	services := service.NewService(store)
	handlers := domain.NewDomain(services)

	srv := server.NewServer(handlers, applicationConfig.HTTPServerPort)

	log.Info().Str("port", applicationConfig.HTTPServerPort).Msg("starting an http srv...")
	go func() {
		if err = srv.HTTPServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Panic().Err(err).Msg("cannot start an http srv")
		}
	}()

	killSignal := make(chan os.Signal, 1)
	signal.Notify(killSignal, syscall.SIGINT, syscall.SIGTERM)
	<-killSignal
	log.Info().Msg("graceful shutdown. This can take a while...")

	gracefulCtx, gracefulCancel := context.WithTimeout(ctx, 30*time.Second)
	defer gracefulCancel()

	if err = srv.HTTPServer.Shutdown(gracefulCtx); err != nil {
		log.Fatal().Err(err).Msg("cannot shutdown srv")
	}
}
