package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/handlers/cart"
	"ecommerce/rest/handlers/products"
	"ecommerce/rest/handlers/reviews"
	"ecommerce/rest/handlers/users"
	middlewares "ecommerce/rest/middleware"
	"ecommerce/user"
	"fmt"
	"os"
)

func Serve() {
	cnf := config.GetConfig()

	dbCon, err := db.NewConnection(&cnf.DB)
	if err != nil {
		fmt.Println("Db connection Error", err)
		os.Exit(1)
	}

	err = db.MigrateDB(dbCon, "./migrations")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//repos
	productRepo := repo.NewProductRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)
	reviewRepo := repo.NewReviewRepo(dbCon)
	cartRepo := repo.NewCartRepo(dbCon)

	//domains

	usrServie := user.NewService(userRepo)

	//middlewares
	middleware := middlewares.NewMiddleware(&cnf)

	//handlers
	producthandler := products.NewHandler(middleware, productRepo)
	userHandler := users.NewHandler(cnf, usrServie)
	reviewHandler := reviews.NewHandler(reviewRepo)
	cartHandler := cart.NewCartHandler(cartRepo, middleware)

	serverobj := rest.NewServer(
		&cnf,
		producthandler,
		userHandler,
		reviewHandler,
		cartHandler,
	)
	serverobj.Start()
}
