package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"net/http"
	"strconv"
)

const (
	REDIS_KEY = "test_list"
	PORT      = ":3001"
)

var RedisConn redis.Conn

func httpHandler(w http.ResponseWriter, r *http.Request) {
	value := r.Form["test"]
	response, _ := RedisConn.Do("lpush", REDIS_KEY, value)
	fmt.Fprintf(w, strconv.Itoa(int(response.(int64))))
}

func main() {
	http.HandleFunc("/", httpHandler)
	error := http.ListenAndServe(PORT, nil)
	if error != nil {
		panic(error)
	}
}

func init() {
	RedisConn, _ = redis.Dial("tcp", ":6379")
}
