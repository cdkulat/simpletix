package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fs))

	mux.HandleFunc("/", app.home)
	// mux.HandleFunc("/t/view", app.ticketView)
	// mux.HandleFunc("t/write", app.ticketWrite)

	return mux
}
