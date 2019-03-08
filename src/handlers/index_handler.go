package handlers

import (
	"log"
	"net/http"
)

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	str := "Sample Message"
	if err := h.Template.ExecuteTemplate(w, "/views/index.tmpl", str); err != nil {
		log.Fatal(err)
	}
}
