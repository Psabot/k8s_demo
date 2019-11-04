package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "3001"
	}

	rand.Seed(time.Now().UnixNano())
	min := 10
	max := 30
	randN := rand.Intn(max-min+1) + min

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Morpheus, my ID is: %v\n", randN)
	})
	http.ListenAndServe(":"+PORT, nil)
}
