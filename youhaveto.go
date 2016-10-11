package main

import (
	"net/http"
	"os"

	"github.com/sparkymat/resty/router"
	"github.com/sparkymat/resty/verb"
	"github.com/sparkymat/youhaveto/controller"
)

func main() {
	r := router.New()

	r.Resource([]string{"recommendations"}, controller.Recommendation{}).Only(verb.Create, verb.Index)

	r.EnableDebug()
	r.PrintRoutes(os.Stdout)
	r.HandleRoot()
	http.ListenAndServe(":4000", nil)
}
