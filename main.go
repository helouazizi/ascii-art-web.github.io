package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"ascii-art-web/server"
)

func main() {
	Temple01, err := template.ParseFiles("template/index.html")
	if err != nil {
		log.Fatal(err)
	}
	ErrorTemplate, err := template.ParseFiles("template/error.html")
	if err != nil {
		log.Fatal(err)
	}

	Pointer := &server.ParsedFiles{
		Temple01:      Temple01,
		ErrorTemplate: ErrorTemplate,
	}

	http.HandleFunc("/", Pointer.Home)
	http.HandleFunc("/css/style.css", Pointer.CssHandler)
	http.HandleFunc("/ascii-art", Pointer.SubmitHandler)
	http.HandleFunc("/css/error.css", Pointer.CssErrHundle)

	fmt.Println("Server is running on port 8080", ">>> http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
