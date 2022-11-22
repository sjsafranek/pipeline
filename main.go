package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/uuid"
)

var (
	config_file string
)

func init() {
	flag.StringVar(&config_file, "config", "", "Config file")
	flag.Parse()
}

func main() {

	// Get specified pipeline
	pipeline := Pipeline{}
	err := pipeline.Load(config_file)
	if nil != err {
		log.Fatal(err)
	}

	// Get context
	ctx := context.Background()
	ctx = context.WithValue(ctx, "id", uuid.New().String())
	ctx = context.WithValue(ctx, "start_time", time.Now())
	ctx, cancel := context.WithCancel(ctx)

	// Wait for signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	defer func() {
		signal.Stop(c)
		cancel()
	}()
	go func() {
		select {
		case <-c:
			log.Println("Shutdown")
			cancel()
		case <-ctx.Done():
		}
	}()

	// Kick off pipeline
	err = pipeline.Do(ctx)
	if nil != err {
		log.Fatal(err)
	}
}
