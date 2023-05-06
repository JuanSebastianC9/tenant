package main

import (
    "html/template"
    "net/http"
)

func main() {
    http.HandleFunc("/", mainHandler)
    http.HandleFunc("/info", infoHandler)
    http.HandleFunc("/contact", contactHandler)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
    http.ListenAndServe(":8080", nil)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("templates/main.html")
    t.Execute(w, nil)
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("templates/info.html")
    t.Execute(w, nil)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("templates/contact.html")
    t.Execute(w, nil)
}