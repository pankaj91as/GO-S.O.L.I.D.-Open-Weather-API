package main

import (
	"context"
	"controller"
	"flag"
	"fmt"
	"log"
	"logger"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func main(){
    // Command Line Option To Set Server Gracefuls Shutdown Timeout
	var (
        wait time.Duration
        WriteTimeout time.Duration
        ReadTimeout time.Duration
        IdleTimeout time.Duration
    )
    flag.DurationVar(&wait, "graceful-timeout", time.Second * 15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
    flag.DurationVar(&WriteTimeout, "write-timeout", time.Second * 15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
    flag.DurationVar(&ReadTimeout, "read-timeout", time.Second * 15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
    flag.DurationVar(&IdleTimeout, "ideal-timeout", time.Second * 60, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
    flag.Parse()

    // Imlement Router
	r := mux.NewRouter()
	r.Use(logger.LoggingMiddleware)
	r.HandleFunc("/", controller.LandingPage)
	r.HandleFunc("/weather/{location}", controller.FetchWeather).Methods(http.MethodGet)
	r.HandleFunc("/weather/current/{location}", controller.FetchCurrentWeather).Methods(http.MethodGet)

    // Implement Server
	srv := &http.Server{
        Addr:           "0.0.0.0:8080",
        WriteTimeout:   WriteTimeout,
        ReadTimeout:    ReadTimeout,
        IdleTimeout:    IdleTimeout,
        Handler: 		r,
    }

    // Run our server in a goroutine so that it doesn't block.
    go func() {
        fmt.Println("Server starting on ", srv.Addr)
        if err := srv.ListenAndServe(); err != nil {
            log.Println(err)
        }
    }()

    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    <-c
    ctx, cancel := context.WithTimeout(context.Background(), wait)
    defer cancel()
    srv.Shutdown(ctx)
    log.Println("shutting down")
    os.Exit(0)
}
