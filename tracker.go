package main

import (
	"flag"
	"net/http"
	"strings"

	"github.com/andrzejewsky/tracker/server"
	"github.com/andrzejewsky/tracker/store"
)

var httpPort int
var storeName string
var storeConfig string

func init() {
	flag.IntVar(&httpPort, "http-port", 8080, "http port for listening request")
	flag.StringVar(&storeName, "store", "redis_queue", "name of store")
	flag.StringVar(&storeConfig, "store-config", "", "store configuration")
}

func main() {
	flag.Parse()

	requestBus := make(chan *http.Request)

	sem := make(chan bool)

	httpServer := server.NewHTTPServer()
	go httpServer.StartListening(httpPort, requestBus)

	store := store.SelectStoreByName(storeName, getStoreConfig(storeConfig))

	go store.StartSave(requestBus)

	<-sem
}

func getStoreConfig(lineConfig string) map[string]string {

	config := make(map[string]string)

	elements := strings.Split(lineConfig, ",")

	for _, row := range elements {
		keyValue := strings.Split(row, ":")
		config[keyValue[0]] = keyValue[1]
	}

	return config
}
