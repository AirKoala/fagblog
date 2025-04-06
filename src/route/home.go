package route

import (
	"github.com/airkoala/fagblog/src/fagblog"
	"log"
	"net/http"
)

func Home(context *fagblog.Context) Route {
	return Route{
		Pattern: "GET /{$}",
		Handler: func(w http.ResponseWriter, r *http.Request) {

			err := context.Templates.ExecuteTemplate(w, "home.html", context)

			if err != nil {
				log.Printf("Error executing template: %v\n", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}}
}
