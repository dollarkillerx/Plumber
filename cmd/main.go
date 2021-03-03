package main

import (
	"flag"
	"log"
	"strings"

	"github.com/dollarkillerx/plumber/internal/config"
	"github.com/dollarkillerx/plumber/internal/scheduler"
)

var configFileName = flag.String("cfn", "config", "name of configs file")
var configFilePath = flag.String("cfp", "./configs/", "path of configs file")

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	flag.Parse()
	err := config.InitConfiguration(*configFileName, strings.Split(*configFilePath, ","), &config.CONF)
	if err != nil {
		panic(err)
	}

	ser := scheduler.New(*config.CONF)
	if err := ser.ListenAndServe(); err != nil {
		log.Fatalf("%+v", err)
	}
}
