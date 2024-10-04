package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/redis/go-redis/v9"
)

func main() {
	// int a = 1;
	// a := 1;
	//

	// int add(int a, int b) { return a + b; }
	// func add(a int, b int) int { return a + b }
	// c := add(1,2)

	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	client := redis.NewClient(&redis.Options{
        Addr:	  redisHost + ":" + redisPort,
        Password: "", // no password set
        DB:		  0,  // use default DB
    })

    app := NewApp(client)

	mux := http.NewServeMux()
	mux.HandleFunc("/users", app.HandleVisits)

	
	if err := http.ListenAndServe("0.0.0.0:8080", mux); err != nil {
		fmt.Printf("Failed to start server %v", err)
	}
}

type App struct {
	RedisClient *redis.Client
}

func (app *App) HandleVisits(w http.ResponseWriter, r *http.Request) {
	result, err := app.RedisClient.Incr(r.Context(), "visits").Result()
	if err != nil {
		fmt.Fprint(os.Stdout, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to fetch from redis"))
		return
	}
	w.Write([]byte(fmt.Sprintf("This is the %d visitor", result)))
}

func NewApp(redisClient *redis.Client) *App {
	app := App{
		RedisClient: redisClient,
	}
	return &app
}
