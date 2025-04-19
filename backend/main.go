package main

import (
	"log"
	"net/http"
)

func europeServer(errChan chan<- error) {
	server := http.NewServeMux()
	server.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Home Page frmo Europe!"))
	})
	log.Println("Starting Server 1 on port 8080...")
	err := http.ListenAndServe(":8080", server)
	errChan <- err
}

func usaServer(errChan chan<- error) {
	server := http.NewServeMux()
	server.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Home Page frmo USA!"))
	})
	log.Println("Starting Server 2 on port 8081...")
	err := http.ListenAndServe(":8081", server)
	errChan <- err
}

func main() {
	errChan := make(chan error, 2)
	go europeServer(errChan) // Server 1 running on port 8080
	go usaServer(errChan)    // Server 2 running on port 8081

	select {
	case err := <-errChan:
		if err != nil {
			log.Printf("Server 1 encountered an error: %v\n", err)
		}
	case err := <-errChan:
		if err != nil {
			log.Printf("Server 2 encountered an error: %v\n", err)
		}
	}
}
