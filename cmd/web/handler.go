package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/namikaze-dev/cyoa/views"
	"github.com/namikaze-dev/cyoa/internal/store"
)

type Handler struct {
	Store map[string]store.Story
}

const defaultArcName = "intro"

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	arcName := r.URL.Path
	if arcName == "/" {
		arcName = defaultArcName
	} else {
		arcName = strings.TrimLeft(arcName, "/")
	}

	story, ok := h.Store[arcName]
	if !ok {
		http.Error(w, fmt.Sprintf("story with name: %q does not exist", arcName), http.StatusBadRequest)
		return
	}

	h.render(w, http.StatusOK, "index", &story)
}

func (h *Handler) render(w http.ResponseWriter, status int, page string, data *store.Story) {
	ts, err := template.ParseFS(views.Files, fmt.Sprintf("html/%v.html", page))
	if err != nil {
		fmt.Println("cyoa:", fmt.Errorf("the template %s does not exist", page))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	buf := &bytes.Buffer{}
	err = ts.Execute(buf, data)
	if err != nil {
		fmt.Println("cyoa:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	buf.WriteTo(w)
}
