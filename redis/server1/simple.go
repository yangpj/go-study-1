package main

import (
	redis "github.com/dotcloud/go-redis-server"
)

type MyHandler struct {
	values map[string][]byte
}

func (h *MyHandler) GET(key string) ([]byte, error) {
	v := h.values[key]
	return v, nil
}

func (h *MyHandler) SET(key string, value []byte) error {
	h.values[key] = value
	return nil
}

func main() {
	handler, _ := redis.NewAutoHandler(&MyHandler{values: make(map[string][]byte)})
	server := &redis.Server{Handler: handler, Addr: ":6389"}
	server.ListenAndServe()
}
