package server

import (
	"net/http"
	"github.com/gorilla/mux"
)

type Endpoint struct {
	Path string
	Verb string
	Handler http.Handler
}

func Start(endpoints ...Endpoint) error {
	r := mux.NewRouter()
	for _, v := range endpoints {
		r.Handle(v.Path, v.Handler).Methods(v.Verb)
	}
	http.Handle("/", r)
	err := http.ListenAndServe(":8080", nil)
	return err
}

