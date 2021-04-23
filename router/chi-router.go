package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type chiRouter struct {
	router *chi.Mux
}

func NewChiRouter() Router {
	return &chiRouter{router: chi.NewRouter()}
}

func (r *chiRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	r.router.Get(uri, f)
}

func (r *chiRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	r.router.Post(uri, f)
}

func (r *chiRouter) SERVE(port string) {
	fmt.Printf("Chi HTTP server running on port %v", port)
	http.ListenAndServe(port, r.router)
}
