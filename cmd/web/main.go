package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/namikaze-dev/cyoa/views"
)

type Options struct {
	Source, Addr string
}

func main() {
	options := parseFlags()
	store, err := NewStoryStore(openFile(options.Source))
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

	err = http.ListenAndServe(options.Addr, mux)
	HandleFatalError(err)
}

func parseFlags() Options {
	var options Options

	flag.StringVar(&options.Addr, "addr", ":8000", "http server address")
	flag.StringVar(&options.Source, "json", "", "json file containing story data")
	flag.Parse()

	return options
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
