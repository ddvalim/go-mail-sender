package routes

import (
	"github.com/ddvalim/go-mail-sender/cmd"
	"github.com/ddvalim/go-mail-sender/core/ports"
	"net/http"
)

var AuthRoutes = []ports.Route{
	{
		URI:    "/auth",
		Method: http.MethodPost,
		Func:   cmd.NewAuthHandler().Auth,
	},
	{
		URI:    "/callback",
		Method: http.MethodGet,
		Func:   cmd.NewAuthHandler().Callback,
	},
}
