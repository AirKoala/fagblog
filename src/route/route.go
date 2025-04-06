package route

import (
	"github.com/airkoala/fagblog/src/middleware"
	"net/http"
)

type Route struct {
	Pattern     string
	Handler     http.HandlerFunc
	Middlewares []middleware.Middleware
}
