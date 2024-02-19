package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"stock/internal/config"
	"stock/internal/domain"
	"stock/internal/pkg/logger"
	"stock/internal/pkg/server"
	"stock/internal/pkg/service"
	"stock/internal/pkg/storage"
	"stock/internal/pkg/storage/item_storage"
	"stock/internal/pkg/storage/stock_storage"

	"syscall"
	"time"

	"github.com/rs/zerolog/log"
)

// insert into items (name, size, quantity, created_at) values ('asd', 'S', 30, now()), ('qwe', 'M', 13, now()), ('zxc', 'M', 9, now()), ('jhg', 'L', 65, now()), ('iop', 'S', 34, now());

// insert into stocks (name, is_available, created_at) values ('1', true, now()), ('2', false, now()), ('3', true, now()), ('4', true, now()), ('5', true, now());

// insert into items_stocks (stock_id, item_id, quantity, created_at) values ((select id from stocks where name = '1'), (select id from items where name = 'asd'), 15, now()), ((select id from stocks where name = '3'), (select id from items where name = 'asd'), 15, now()), ((select id from stocks where name = '3'), (select id from items where name = 'zxc'), 9, now());

// insert into items (name, size, quantity, created_at) values ('dress', 'S', 10, now());
// insert into stocks (name, is_available, created_at) values ('dress_stock', true, now());
// insert into items_stocks (stock_id, item_id, quantity, created_at) values ((select id from stocks where name = 'dress_stock'), (select id from items where name = 'dress'), 10, now());
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
