package store

import (
	"net/http"

	"github.com/andrzejewsky/tracker/store/queue"
)

// Store interface for each stores requests
type Store interface {
	StartSave(requestBus chan *http.Request)
}

// SelectStoreByName create a store by name
func SelectStoreByName(name string, config map[string]string) Store {
	switch name {
	case "redis_queue":
		return queue.NewRedisQueue(
			map[string]string{
				"host": config["host"],
				"port": config["port"],
			})
	}

	return nil
}
