package main

import (
	"fmt"
	"go-server/functions"
	"html/template"
	"net/http"
	"strings"
)

type PageData struct {
	Message string
}

func home(w http.ResponseWriter, r *http.Request) {

	// pease the html template
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "eror", http.StatusInternalServerError)
		return
	}
	// now  execute the template
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
	}
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Parse the form
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}

		// Get the input text
		inputText := r.FormValue("inputText")
		slice := functions.ReadFile("./banners/standard.txt")
		joinedString := strings.Join(slice, " ")
		test := []byte(joinedString)
		outout := functions.TraitmentData(test, inputText, false)
		// Prepare the message to display
		message := outout
		fmt.Println(message)

		// Render the home page with the message
		tmpl, err := template.ParseFiles("index.html", "style.css")
		if err != nil {
			http.Error(w, "Unable to load template", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, PageData{Message: message})
		if err != nil {
			http.Error(w, "Unable to execute template", http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/t", submitHandler) // Handle POST /submit
	fmt.Println("srever is running i  port 8080 ", ">>> http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
