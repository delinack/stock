package server

import (
	"net/http"
	"storage/internal/domain"
	"time"

	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json"
	"github.com/rs/zerolog/log"
)

// Server .
type Server struct {
	HTTPServer http.Server
}

// NewServer constructor for server
func NewServer(handlers *domain.Domain, port string) *Server {
	var srv Server

	server := rpc.NewServer()

	server.RegisterCodec(json.NewCodec(), "application/json")
	err := server.RegisterService(handlers.StockDomain, "Stock")
	if err != nil {
		log.Error().Err(err).Msgf("cannot register service: Stock")
	}
	err = server.RegisterService(handlers.ItemDomain, "Item")
	if err != nil {
		log.Error().Err(err).Msgf("cannot register service: Item")
	}

	srv.HTTPServer = http.Server{
		Handler:           server,
		Addr:              ":" + port,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       5 * time.Second,
	}

	//srv.Echo = echo.New()
	//srv.Echo.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
	//	LogLatency:  true,
	//	LogRemoteIP: true,
	//	LogURI:      true,
	//	LogMethod:   true,
	//	LogError:    true,
	//	LogValuesFunc: func(_ echo.Context, v middleware.RequestLoggerValues) error {
	//		log.Info().Str("comp", "server").
	//			Str("URI", v.URI).Str("remote_ip", v.RemoteIP).
	//			Str("method", v.Method).Dur("latency", v.Latency).
	//			Interface("params", v.FormValues).
	//			Send()
	//		return nil
	//	},
	//}))

	//r := mux.NewRouter()
	//r.Handle("/", server)

	return &srv
}
