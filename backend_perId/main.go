package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	
)

var payloads map[int][]byte
var once sync.Once

func main() {
	generatePayloads()

	http.HandleFunc("/publiccache/", func(w http.ResponseWriter, r *http.Request) {
		PublicCacheHandler(w, r)
	})

	// Start server
	fmt.Println("Server is listening on port 8081...")
	http.ListenAndServe(":8081", nil)
}

func PublicCacheHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/publiccache/"):])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	payload, ok := payloads[id]
	if !ok {
		http.Error(w, "ID not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Cache-Control", "public, max-age=60")
	w.Write(payload)
	additionalContent := []byte(" cached as public for ID " + strconv.Itoa(id))
	w.Write(additionalContent)
}

func generatePayloads() {
	once.Do(func() {
		payloads = make(map[int][]byte)
		for i := 1; i <= 10; i++ {
			payload := make([]byte, 100)
			for j := range payload {
				payload[j] = 'x'
			}
			payloads[i] = payload
		}
	})
}
