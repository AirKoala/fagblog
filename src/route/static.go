package route

import (
	"github.com/airkoala/fagblog/src/fagblog"
	"net/http"
)

func Static(context *fagblog.Context, config *fagblog.Config) Route {
	fs := http.FileServer(http.Dir(config.ContentDir))
	return Route{
		Pattern: "GET /static/",
		Handler: http.StripPrefix("/static/", fs).ServeHTTP,
	}
}
