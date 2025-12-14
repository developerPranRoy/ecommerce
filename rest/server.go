package rest

import (
	"ecommerce/config"
	"ecommerce/rest/handlers/cart"
	"ecommerce/rest/handlers/products"
	"ecommerce/rest/handlers/reviews"
	"ecommerce/rest/handlers/users"
	middlewares "ecommerce/rest/middleware"
	"fmt"
	"net/http"
	"strconv"
)

type Server struct {
	cnf            *config.Config
	producthandler *products.Handler
	userHandler    *users.Handler
	reviewHandler  *reviews.Handler
	cartHandler    *cart.Handler
}

func NewServer(
	cnf *config.Config,
	producthandler *products.Handler,
	userHandler *users.Handler,
	reviewHandler *reviews.Handler,
	cartHandler *cart.Handler,

) *Server {
	return &Server{
		cnf:            cnf,
		producthandler: producthandler,
		userHandler:    userHandler,
		reviewHandler:  reviewHandler,
		cartHandler:    cartHandler,
	}
}

func (server *Server) Start() {

	manager := middlewares.NewManager()
	manager.Use(
		middlewares.CorsPreflight,
		middlewares.Logger,
	)
	mux := http.NewServeMux()
	server.producthandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)
	server.reviewHandler.RegisterRoutes(mux, manager)
	server.cartHandler.RegisterRoutes(mux, manager)

	address := ":" + strconv.Itoa(server.cnf.HttpPort)
	fmt.Println("✅ Server started at ", address)
	err := http.ListenAndServe(address, mux)
	if err != nil {
		fmt.Println("Error to starting server", err)

	}

}
