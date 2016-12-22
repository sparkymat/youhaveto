package main

import (
	"io"
	"net/http"
	"os"

	"github.com/sparkymat/reactor"
	"github.com/sparkymat/resty/router"
	"github.com/sparkymat/resty/verb"
	"github.com/sparkymat/youhaveto/controller"
)

func main() {
	r := router.New()

	r.Resource([]string{"recommendations"}, controller.Recommendation{}).Only(verb.Create, verb.Index)

	r.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		app := reactor.New("YouHaveTo")
		app.MapJavascriptFolder("public/js", "js")
		app.MapCssFolder("public/css", "css")

		app.AddCustomJavascriptLink("/js/lib/jquery-3.1.1.min.js")
		app.AddCustomJavascriptLink("/js/lib/bootstrap.js")
		app.AddCustomCssLink("/css/bootstrap.min.css")

		io.WriteString(response, app.Html().String())
	})

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	r.EnableDebug()
	r.PrintRoutes(os.Stdout)
	r.HandleRoot()
	http.ListenAndServe(":4000", nil)
}
