package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
}

// Method in APIServer to run the application
func (s *APIServer) Run() error {
	router := mux.NewRouter()

	// Could add a prefix to maintain versioning in Routing
	// subRouter := router.PathPrefix("/api/v1").Subrouter()

	//create a new service and register it to the router or subRouter if exists
	receiptService := NewHandler()
	receiptService.RegisterRoutes(router)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}
