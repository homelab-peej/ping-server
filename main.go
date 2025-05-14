package main

import (
	"fmt"
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ok")
}

func readyzHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ready")
}

func startzHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "started")
}

func main() {
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/healthz", healthzHandler)
	http.HandleFunc("/readyz", readyzHandler)
	http.HandleFunc("/startz", startzHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Failed to start server:", err)
		return
	}
}
