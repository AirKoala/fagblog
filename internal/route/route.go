package route

import (
	"net/http"

	"github.com/airkoala/fagblog/internal/fagblog"
	"github.com/airkoala/fagblog/internal/middleware"
)

type Data struct {
	Context *fagblog.Context
}

type Route struct {
	Pattern     string
	Handler     http.HandlerFunc
	Middlewares []middleware.Middleware
}
