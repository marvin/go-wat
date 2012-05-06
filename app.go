package main

// imports
import (
	"html/template"
	"net/http"
)

func start(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, nil)
}

func main() {
	// serve functions
	http.HandleFunc("/", start)
	// serve static
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./public/js"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./public/css"))))
	// listen
	http.ListenAndServe(":8080", nil)
}
