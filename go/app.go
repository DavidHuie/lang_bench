package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"net/http"
	"runtime"
	"strconv"
)

const (
	REDIS_KEY       = "test_list"
	REDIS_POOL_SIZE = 20
	PORT            = ":3001"
)

var RedisPool *redis.Pool

func httpHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	value := r.Form["test"][0]
	conn := RedisPool.Get()
	defer conn.Close()
	response, err := conn.Do("lpush", REDIS_KEY, value)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Fprintf(w, strconv.Itoa(int(response.(int64))))
}

func main() {
	http.HandleFunc("/", httpHandler)
	error := http.ListenAndServe(PORT, nil)
	if error != nil {
		panic(error)
	}
}

func newConn() (redis.Conn, error) {
	return redis.Dial("tcp", ":6379")
}

func init() {
	RedisPool = redis.NewPool(newConn, REDIS_POOL_SIZE)
	// Use all CPUs
	runtime.GOMAXPROCS(runtime.NumCPU())
}
