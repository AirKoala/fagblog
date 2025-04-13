package route

import (
	"github.com/airkoala/fagblog/internal/fagblog"
	"net/http"
)

func Static(context *fagblog.Context, config *fagblog.Config) Route {
	fs := http.FileServer(http.Dir(config.StaticDir))
	return Route{
		Pattern: "GET /static/",
		Handler: http.StripPrefix("/static/", fs).ServeHTTP,
	}
}
