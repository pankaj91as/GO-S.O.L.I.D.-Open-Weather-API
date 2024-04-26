package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	"github.com/op/go-logging"
)

var Log = logging.MustGetLogger("rest")

var (
	Wait         time.Duration
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
	IdleTimeout  time.Duration
	DBhost       string
	DBport       int
	DBusername   string
	DBpassword   string
	DBname       string
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Command Line Option To Set Server Gracefuls Shutdown Timeout
	flag.DurationVar(&Wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.DurationVar(&WriteTimeout, "write-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.DurationVar(&ReadTimeout, "read-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.DurationVar(&IdleTimeout, "ideal-timeout", time.Second*60, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")

	// Command Line Option To Set DB Credentials
	flag.StringVar(&DBhost, "db-host", "0.0.0.0", "database host domain/ip - e.g. localhost or 0.0.0.0")
	flag.IntVar(&DBport, "db-port", 3306, "database port number - e.g. 3306")
	flag.StringVar(&DBusername, "db-username", "root", "database user name - e.g. admin or root")
	flag.StringVar(&DBpassword, "db-password", "password", "database user secret/password")
	flag.StringVar(&DBname, "database", "open_weather", "database user secret/password")
	flag.Parse()

	// Implement Server
	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: WriteTimeout,
		ReadTimeout:  ReadTimeout,
		IdleTimeout:  IdleTimeout,
		Handler:      Router().InitRouter(),
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
	ctx, cancel := context.WithTimeout(context.Background(), Wait)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}
