package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/kasariks/api_golang/service/cart"
	"github.com/kasariks/api_golang/service/order"
	"github.com/kasariks/api_golang/service/product"
	"github.com/kasariks/api_golang/service/user"
)

type APIServer struct {
	addr   string
	db     *sql.DB
	router *http.ServeMux
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr:   addr,
		db:     db,
		router: http.NewServeMux(),
	}
}

func (s *APIServer) Run() error {
	s.registerServicesRoutes()

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, s.router)
}

func (s *APIServer) registerServicesRoutes() {
	userStore := user.NewStore(s.db)
	userHandler := user.NewHanlder(userStore)
	userHandler.RegisterRoutes(s.router)

	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore)
	productHandler.RegisterRoutes(s.router)

	orderStore := order.NewStore(s.db)
	cartStore := cart.NewHandler(orderStore, productStore, userStore)
	cartStore.RegisterRoutes(s.router)
}
