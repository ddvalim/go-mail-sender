package routes

import (
	"github.com/ddvalim/go-mail-sender/cmd"
	"github.com/ddvalim/go-mail-sender/core/ports"
	"net/http"
)

var MailRoutes = []ports.Route{
	{
		URI:    "/send",
		Method: http.MethodPost,
		Func:   cmd.NewHandler().Send,
	},
}
