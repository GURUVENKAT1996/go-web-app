package main

import (
    "log"
    "net/http"
    "time"
)

func homePage(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "static/home.html")
}

func coursePage(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "static/courses.html")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "static/about.html")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "static/contact.html")
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/home", homePage)
    mux.HandleFunc("/courses", coursePage)
    mux.HandleFunc("/about", aboutPage)
    mux.HandleFunc("/contact", contactPage)

    server := &http.Server{
        Addr:         "0.0.0.0:8080",
        Handler:      mux,
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  120 * time.Second,
    }

    log.Println("Starting server on :8080")
    log.Fatal(server.ListenAndServe())
}
``