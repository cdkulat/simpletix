package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// implement later - set to display 5 most recent tickets
	tickets, err := app.tickets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	// loop that initializes the display of each ticket and then dumps it
	for _, ticket := range tickets {
		fmt.Fprint(w, "%+v\n", ticket)
	}

	// initialize filesystem to parse files for templating
	files := []string{
		"./ui/html/base.html",
		"./ui/html/home.html",
		"./ui/html/nav.html",
	}

	// reads the files and stores the templates
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// Executes templates and writes content into base template
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.errorLog.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
}
