package main

import (
	"context"
	"flag"
	"log"

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

	err = pipeline.Do(ctx)
	if nil != err {
		log.Fatal(err)
	}
}
