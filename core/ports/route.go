package ports

import "net/http"

type Route struct {
	URI    string
	Method string
	Func   func(w http.ResponseWriter, r *http.Request)
}
