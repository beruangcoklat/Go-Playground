package main

import (
	"fmt"
	"net/http"

	"github.com/beruangcoklat/Go-Playground/config"
	"github.com/beruangcoklat/Go-Playground/internal/handler"
	"github.com/gorilla/mux"
)

func router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.Index)
	r.HandleFunc("/code", handler.RunCode)
	return r
}

func main() {
	fmt.Println("listen ", config.PORT)
	http.ListenAndServe(":"+config.PORT, router())
}
