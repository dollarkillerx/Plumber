package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/dollarkillerx/plumber/internal/config"
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

	fmt.Println(config.CONF)

	marshal, err := json.Marshal(config.CONF)
	if err == nil {
		fmt.Println(string(marshal))
	}
}
