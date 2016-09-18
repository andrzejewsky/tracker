[![Go Report Card](https://goreportcard.com/badge/github.com/andrzejewsky/tracker)](https://goreportcard.com/report/github.com/andrzejewsky/tracker)

An example of analytics tracker which will save http requests into defined store.
In the production environment it should be a something with very fast saving,
in order to relieve http tracking (like queue, fast database etc.).

Example usage:

```
./tracker -http-port=8080 -store="redis_queue" -store-config="host:localhost,port:6379"
```
