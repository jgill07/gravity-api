package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jgill07/gravity-api/internal/log"
)

type ServiceHandler func() error
type ShutdownHandler func(ctx context.Context) error

func setup(router http.Handler, port int) (ServiceHandler, ShutdownHandler) {
	srv := &http.Server{
		Addr:    fmt.Sprint(":", port),
		Handler: router,
	}
	svcHandler := func() error {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.WithError(err).Fatal("API server's http server failed to listen and serve")
			return err
		}
		return nil
	}
	shutdownHandler := func(ctx context.Context) error {
		if err := srv.Shutdown(ctx); err != nil {
			log.WithError(err).Error("service failed to shut down")
		}
		return nil
	}
	return svcHandler, shutdownHandler
}

func Run(router http.Handler, port int) {
	svcHandler, shutdownHandler := setup(router, port)
	go func() {
		if err := svcHandler(); err != nil {
			log.WithError(err).Fatal("API server failed to start")
		}
		log.Info("API server stopped")
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// Wait for interrupt signal to gracefully shutdown the server with
	<-quit
	sdCtx, cancel := context.WithTimeout(context.Background(), time.Second*25)
	defer cancel()
	if err := shutdownHandler(sdCtx); err != nil {
		log.WithError(err).Error("service failed to shut down")
	}
	log.Info("service stopped running")
}
