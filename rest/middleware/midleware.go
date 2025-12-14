package middlewares

import "ecommerce/config"

type Middlewares struct {
	cnf *config.Config
}

func NewMiddleware(cnf *config.Config) *Middlewares {
	return &Middlewares{
		cnf: cnf,
	}
}
