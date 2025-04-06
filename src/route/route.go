package route

import (
	"net/http"

	"github.com/airkoala/fagblog/src/fagblog"
	"github.com/airkoala/fagblog/src/middleware"
)

type Data struct {
	Context *fagblog.Context
}

type Route struct {
	Pattern     string
	Handler     http.HandlerFunc
	Middlewares []middleware.Middleware
}
