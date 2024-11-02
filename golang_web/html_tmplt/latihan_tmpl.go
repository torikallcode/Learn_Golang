package main

import (
	"html/template"
	"net/http"
)

type InitialData struct {
	Title, Name string
	Age         int
}

func initialFunc(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("latihan_tmpl.html")
	if err != nil {
		http.Error(w, "Error parsing Template", http.StatusInternalServerError)
	}

	data := InitialData{
		Title: "Latihan HTML Page",
		Name:  "M. Torikal Akbar",
		Age:   20,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error rendring template", http.StatusInternalServerError)
	}

}

func main() {
	http.HandleFunc("/", initialFunc)
	http.ListenAndServe(":8080", nil)
}
