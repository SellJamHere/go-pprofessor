package pprofessor

import (
	"log"
	"net/http"
	_ "net/http/pprof" // pprof injects profiling routes
)

// Serve function serves
func Serve(port string) {
	go func() {
		log.Fatal(http.ListenAndServe(port, nil))
	}()
}
