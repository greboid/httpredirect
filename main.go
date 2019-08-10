package main

import (
	"log"
	"net/http"
	"os"
)

var targetHost, targetHostError = os.LookupEnv("TARGETHOST")
var excludeQuery, excludeQueryError = os.LookupEnv("EXCLUDEQUERY")
var targetPath, pathError = os.LookupEnv("TARGETPATH")
var targethttp, targethttpError = os.LookupEnv("TARGETHTTP")

func redirect(w http.ResponseWriter, req *http.Request) {
	var target string
	//Check if http or https
	if !targethttpError {
		target += "https://"
	} else {
		target += "http://"
	}
	//Check what host to point at, defaulting to the requested host
	if !targetHostError {
		target += req.Host
	} else {
		target += targetHost
	}
	//Check whether to redirect to a fixed path, or existing
	if !pathError {
		target += req.URL.Path
	} else {
		target += targetPath
	}
	//Should we include the query parameters?
	if !excludeQueryError {
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
