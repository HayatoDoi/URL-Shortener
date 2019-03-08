package main

import (
	"net/http"
	"net/http/cgi"

	"github.com/HayatoDoi/shortener/src/handlers"
)

//go:generate go-assets-builder -p views -o views/views.go views/

func main() {
	handler := handlers.New()
	handler.AddTemplate("/views/index.tmpl")

	http.HandleFunc("/", handler.Index)
	cgi.Serve(nil)

}
