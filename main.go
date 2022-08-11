package main

import (
	"flag"
	"log"
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
	err = pipeline.Do()
	if nil != err {
		log.Fatal(err)
	}
}
