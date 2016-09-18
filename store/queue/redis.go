package queue

import (
	"encoding/json"
	"net/http"

	"github.com/andrzejewsky/go-queue/queues"
	"github.com/andrzejewsky/tracker/server"
)

// RedisQueue queue based on redis
type RedisQueue struct {
	queue queues.Queue
}

// NewRedisQueue new instance of queue
func NewRedisQueue(config map[string]string) *RedisQueue {

	queue, _ := queues.GetQueue("redis", config)

	return &RedisQueue{queue}
}

// StartSave start saving
func (q *RedisQueue) StartSave(requestBus chan *http.Request) {
	for request := range requestBus {
		jsonString, _ := json.Marshal(server.CreateByRequest(request))
		q.queue.Push(string(jsonString))
	}
}
