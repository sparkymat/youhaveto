package controller

import "net/http"

type Recommendation struct {
}

func (r Recommendation) Index(response http.ResponseWriter, request *http.Request, parmas map[string][]string) {
	response.Write([]byte("Hello, world"))
}
