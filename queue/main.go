package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"net/http"
)

func main() {
	http.HandleFunc("/payments", paymentsHandler)
	http.ListenAndServe(":8080", nil)
}

func paymentsHandler(w http.ResponseWriter, req *http.Request) {

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.TODO()

	buf := new(bytes.Buffer)
	// This validate
	buf.ReadFrom(req.Body)

	paymentDetails := buf.String()

	err := redisClient.RPush(ctx, "payments", paymentDetails).Err()

	if err != nil {
		fmt.Fprintf(w, err.Error()+"\r\n")
	} else {
		fmt.Fprintf(w, "Payment details accepted successfully\r\n")
	}

}
