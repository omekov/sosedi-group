package api

import (
	"context"
	"flag"
	"fmt"
	"log"
	gohttp "net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/omekov/sosedi-group/internal/api/delivery/http"
	"github.com/omekov/sosedi-group/internal/counter"
	"github.com/omekov/sosedi-group/pkg/redis"
)

var flagPort = flag.String("port", "80", "specify the port")

func Run() {
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	redisClient := redis.NewClient("localhost:6379", "")
	defer redisClient.Close()

	if err := redisClient.Ping(ctx).Err(); err != nil {
		log.Fatal(fmt.Errorf("redis not connection: %s", err))
	}

	counterService := counter.NewService(redisClient)

	handler := http.NewHandler(counterService)

	httpServer := NewServer(handler.Init())

	go func() {
		if err := httpServer.ListenAndServe(); err != gohttp.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("failed to stop server: %v", err)
	}

}

func NewServer(handler gohttp.Handler) gohttp.Server {
	return gohttp.Server{
		Addr:           fmt.Sprintf(":%s", *flagPort),
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
