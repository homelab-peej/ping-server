package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	totalRequests = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "ping_server_total_http_requests",
		Help: "Total number of HTTP requests.",
	}, []string{"method", "endpoint"})

	totalPongs = promauto.NewCounter(prometheus.CounterOpts{
		Name: "ping_server_total_pongs",
		Help: "Total number of pong responses.",
	})
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	defer totalRequests.WithLabelValues(r.Method, "/ping").Inc()
	defer totalPongs.Inc()
	fmt.Fprintln(w, "pong")
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	defer totalRequests.WithLabelValues(r.Method, "/healthz").Inc()
	fmt.Fprintln(w, "ok")
}

func readyzHandler(w http.ResponseWriter, r *http.Request) {
	defer totalRequests.WithLabelValues(r.Method, "/readyz").Inc()
	fmt.Fprintln(w, "ready")
}

func startzHandler(w http.ResponseWriter, r *http.Request) {
	defer totalRequests.WithLabelValues(r.Method, "/startz").Inc()
	fmt.Fprintln(w, "started")
}

func main() {
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/healthz", healthzHandler)
	http.HandleFunc("/readyz", readyzHandler)
	http.HandleFunc("/startz", startzHandler)

	http.Handle("/metrics", promhttp.HandlerFor(
		prometheus.DefaultGatherer, promhttp.HandlerOpts{
			EnableOpenMetrics: true,
		},
	))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Failed to start server:", err)
		return
	}
}
