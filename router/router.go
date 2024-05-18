package router

import (
	"github.com/ddvalim/go-mail-sender/core/ports"
	"github.com/ddvalim/go-mail-sender/router/routes"
	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	return configRouter(mux.NewRouter())
}

func configRouter(router *mux.Router) *mux.Router {
	var r []ports.Route

	r = append(r, routes.MailRoutes...)

	for _, route := range r {
		router.HandleFunc(route.URI, route.Func).Methods(route.Method)
	}

	return router
}
