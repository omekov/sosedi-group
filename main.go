package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

var flagPort = flag.String("port", "80", "specify the port")

func main() {
	flag.Parse()

	httpServer := http.Server{
		Addr:           fmt.Sprintf(":%s", *flagPort),
		Handler:        GetHandler(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	httpServer.ListenAndServe()
}

func GetHandler() *http.ServeMux {
	httpMux := http.NewServeMux()
	httpMux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("hello"))
	})
	return httpMux
}
