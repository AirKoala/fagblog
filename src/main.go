package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/airkoala/fagblog/src/middleware"
	"github.com/airkoala/fagblog/src/route"
	"github.com/airkoala/fagblog/src/fagblog"
)

func main() {
	// TODO: Load config from file
	config := fagblog.Config{
		Port:        8000,
		TemplateDir: "templates",
		ContentDir:  "content",
		StaticDir:   "static",
	}
	blogMetadata, err := fagblog.BlogMetadataFromToml(config.ContentDir + "/meta.toml")

	if err != nil {
		log.Fatalf("Error loading blog metadata: %v", err)
	}

	mux := http.NewServeMux()

	tmpl, err := template.ParseGlob(config.TemplateDir + "/*.html")
	if err != nil {
		log.Fatalf("Error parsing templates: %v", err)
	}

	context := fagblog.Context{
		BlogMetadata: blogMetadata,
		Templates: tmpl,
	}

	handle(mux, route.Home(&context))
	handle(mux, route.Static(&context, &config))

	log.Printf("Starting server on :%d", config.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", config.Port), mux)

	log.Println("Server stopped: ", err)
}

func handle(s *http.ServeMux, route route.Route) {
	middlewares := route.Middlewares
	if route.Method != "" {
		middlewares = append(middlewares, middleware.Method(route.Method))
	}

	middlewares = append(middlewares, middleware.Logging())

	log.Printf("Registering %s %s\n", route.Method, route.Pattern)

	s.HandleFunc(route.Pattern, middleware.Chain(route.Handler, middlewares...))
}
