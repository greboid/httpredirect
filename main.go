package main

import (
	"log"
	"net/http"
	"os"
)

var targetHost = os.Getenv("HOST")
var excludeQuery = os.Getenv("EXCLUDEQUERY")
var targetPath = os.Getenv("PATH")
var httpRedirect = os.Getenv("HTTPREDIRECT")

func redirect(w http.ResponseWriter, req *http.Request) {
	var target string
	//Check if http or https
	if len(httpRedirect) == 0 {
		target += "https://"
	} else {
		target += "http://"
	}
	//Check what host to point at, defaulting to the requested host
	if len(targetHost) == 0 {
		target += req.Host
	} else {
		target += targetHost
	}
	//Check whether to redirect to a fixed path, or existing
	if len(targetPath) == 0 {
		target += req.URL.Path
	} else {
		target += targetPath
	}
	//Should we include the query parameters?
	if len(excludeQuery) == 0 {
		if len(req.URL.RawQuery) > 0 {
			target += "?" + req.URL.RawQuery
		}
	}
	log.Printf("redirect to: %s", target)
	http.Redirect(w, req, target, http.StatusPermanentRedirect)
}

func main() {
	if err := http.ListenAndServe(":1234", http.HandlerFunc(redirect)); err != http.ErrServerClosed {
		log.Fatalf("ListenAndServe(): %s", err)
	}
}
