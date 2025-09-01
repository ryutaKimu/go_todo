package main

import (
	"context"
	"errors"
	"github.com/ryutaKimu/go_todo/api/router"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	r := router.Router()
	srv := &http.Server{
		Addr:    ":8000",
		Handler: r,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		log.Printf("server starting at port %s", srv.Addr)
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Printf("server error: %v", err)
		}
		log.Println("server stopped serving new connections")
	}()

	<-ctx.Done()
	log.Println("received signal interrupt")

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("server shutdown error: %v", err)
	}

	log.Println("server graceful shutdown complete")
}
