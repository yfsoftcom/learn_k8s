package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"log"
	"net/http"
	"os"
)

var (
	host      string
	mode      string
	redisHost string
	redisPort string
	client    *redis.Client
)

func init() {
	log.SetPrefix("[ Cmode ]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)

	var err error
	host, err = os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	mode = os.Getenv("MODE")
	if mode == "" {
		mode = "DEFAULT"
	}

	redisHost = os.Getenv("REDIS_HOST")
	if redisHost == "" {
		redisHost = "172.17.0.8"
	}

	redisPort = os.Getenv("REDIS_PORT")
	if redisPort == "" {
		redisPort = "6379"
	}
	newClient()
}

func newClient() {
	client = redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort, //
		Password: "",                          // no password set
		DB:       0,                           // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	err := client.Incr("counter").Err()
	if err != nil {
		panic(err)
	}
	val, err := client.Get("counter").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("counter", val)

	v := r.URL.Query()
	a := v.Get("foo")
	fmt.Println("Request hit", a)
	fmt.Fprintf(w, "Welcome you , %s ! From  %s , Mode %s, Counter: %s", a, host, mode, val)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ding")
}

func main() {
	http.HandleFunc("/foo", fooHandler)
	http.HandleFunc("/health", healthHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
