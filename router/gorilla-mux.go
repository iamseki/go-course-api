package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type gorillaMux struct {
	router *mux.Router
}

func NewGorillaMuxRouter() Router {
	return &gorillaMux{router: mux.NewRouter()}
}

func (r *gorillaMux) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	r.router.HandleFunc(uri, f).Methods("GET")
}

func (r *gorillaMux) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	r.router.HandleFunc(uri, f).Methods("POST")
}

func (r *gorillaMux) SERVE(port string) {
	fmt.Printf("Gorilla Mux HTTP server running on port %v", port)
	http.ListenAndServe(port, r.router)
}
