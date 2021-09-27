package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	s := &http.Server{
		Addr: ":8081",
		Handler: &handler{
			version: os.Getenv("VERSION"),
		},
	}
	log.Fatal(s.ListenAndServe())
}

type handler struct {
	version string
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	respHeaders := w.Header()
	for k, v := range r.Header {
		respHeaders[k] = v
	}
	respHeaders.Set("VERSION", h.version)
	httpStaus := http.StatusOK
	if r.URL.Path != "/healthz" {
		httpStaus = http.StatusNotFound
	}
	w.WriteHeader(httpStaus)
	log.Printf("Request from %s, HTTP Code %d", r.RemoteAddr, httpStaus)
	w.Write([]byte(fmt.Sprintf(`{"status":"%s"}`, http.StatusText(httpStaus))))
}
