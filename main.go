package main

import (
	"context"
	"flag"
	"log"
	"os"

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
	pipeline := Pipeline{}
	err := pipeline.Load(config_file)
	if nil != err {
		log.Fatal(err)
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, "id", uuid.New().String())
	ctx, cancel := context.WithCancel(ctx)

	// Wait for signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	defer func() {
		signal.Stop(c)
		cancel()
	}()
	go func() {
		select {
		case <-c:
			cancel()
		case <-ctx.Done():
		}
	}()

	err = pipeline.Do(ctx)
	if nil != err {
		log.Fatal(err)
	}
}
