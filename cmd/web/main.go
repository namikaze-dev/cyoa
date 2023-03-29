package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/namikaze-dev/cyoa/views"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("cyoa: missing required json story store argument")
		os.Exit(1)
	}

	store, err := NewStoryStore(openFile(args[1]))
	if err != nil {
		HandleFatalError(err)
	}

	handler := Handler{
		Store: store,
	}

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.FS(views.Files))
	mux.Handle("/static/", fileServer)
	mux.Handle("/", &handler)

	err = http.ListenAndServe(":8000", mux)
	HandleFatalError(err)
}

func openFile(fn string) *os.File {
	f, err := os.Open(fn)
	if err != nil {
		HandleFatalError(err)
	}
	return f
}

func HandleFatalError(err error) {
	fmt.Println("cyoa:", err)
	os.Exit(1)
}
