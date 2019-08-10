package main

import (
	"log"
	"net/http"
	"os"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func redirect(w http.ResponseWriter, req *http.Request) {
	target := "https://" + getEnv("TARGETHOST", req.Host) + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}
	log.Printf("redirect to: %s", target)
	http.Redirect(w, req, target, http.StatusPermanentRedirect)
}

func main() {
	if err := http.ListenAndServe(":80", http.HandlerFunc(redirect)); err != http.ErrServerClosed {
		log.Fatalf("ListenAndServe(): %s", err)
	}
}
