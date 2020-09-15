package main

import (
	"log"
	"net/http"
	"os"
)

var targetHost = os.Getenv("TARGETHOST")
var excludeQuery = os.Getenv("EXCLUDEQUERY")
var targetPath = os.Getenv("TARGETPATH")
var httpRedirect = os.Getenv("HTTPREDIRECT")
var tempRedirect = os.Getenv("TEMPREDIRECT")

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
	if len(tempRedirect) == 0 {
		http.Redirect(w, req, target, http.StatusTemporaryRedirect)
	} else {
		http.Redirect(w, req, target, http.StatusPermanentRedirect)
	}
}

func main() {
	if err := http.ListenAndServe(":8080", http.HandlerFunc(redirect)); err != http.ErrServerClosed {
		log.Fatalf("ListenAndServe(): %s", err)
	}
}
