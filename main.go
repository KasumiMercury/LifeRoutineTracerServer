package main

import (
	"log/slog"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello, World!"))
		if err != nil {
			slog.Error("Failed to write response", "error", err)
			return
		}
		slog.Info("Handled request", "method", r.Method, "url", r.URL)
	})

	slog.Info("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		slog.Error("Failed to start server", "error", err)
		return
	}
}
