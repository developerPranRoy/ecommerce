package middlewares

import (
	"net/http"
)

type Middleware func(next http.Handler) http.Handler

type Manager struct {
	globle []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globle: make([]Middleware, 0),
	}
}

func (mngr *Manager) Use(mws ...Middleware) {
	mngr.globle = append(mngr.globle, mws...)
}

func (mngr *Manager) With(handler http.Handler, middlewares ...Middleware) http.Handler {
	next := handler
	for i := len(middlewares) - 1; i >= 0; i-- {
		next = middlewares[i](next)
	}
	for i := len(mngr.globle) - 1; i >= 0; i-- {
		next = mngr.globle[i](next)
	}
	return next
}
