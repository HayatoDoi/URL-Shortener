package handlers

import (
	"html/template"

	"github.com/HayatoDoi/shortener/src/views"
)

type Handler struct {
	Template *template.Template
}

func New() *Handler {
	handler := new(Handler)
	handler.Template = template.New("view")
	return handler
}

func (h *Handler) AddTemplate(path string) error {
	file, err := views.Assets.Open(path)
	defer file.Close()
	if err != nil {
		return err
	}
	buf := make([]byte, 1024)
	n, err := file.Read(buf)
	if err != nil {
		return err
	}
	h.Template.New(path).Parse(string(buf[:n]))
	return nil
}
