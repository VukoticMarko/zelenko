package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Router interface {
	GET(uri string, f func(w http.ResponseWriter, r *http.Request))
	POST(uri string, f func(w http.ResponseWriter, r *http.Request))
	DELETE(uri string, f func(w http.ResponseWriter, r *http.Request))
	SERVE(port string)
}

type muxRouter struct {
	handler http.Handler
	router  *mux.Router
}

func NewMuxRouter() Router {

	r := mux.NewRouter()
	// CORS Middleware
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"*"},
	})

	return &muxRouter{
		handler: c.Handler(r),
		router:  r,
	}
}

func (mr *muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	mr.router.HandleFunc(uri, f).Methods("GET")
}

func (mr *muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	mr.router.HandleFunc(uri, f).Methods("POST")
}

func (mr *muxRouter) DELETE(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	mr.router.HandleFunc(uri, f).Methods("DELETE")
}

func (mr *muxRouter) SERVE(port string) {
	fmt.Printf("Mux http server running on port %v\n", port)
	http.ListenAndServe(port, mr.handler)
}
