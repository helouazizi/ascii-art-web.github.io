package main

import (
	"fmt"
	"html/template"
	"net/http"

	"server/functions"
)

type PageData struct {
	Message string
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Eror 404: NOT FOUND", http.StatusNotFound)
		return
	}
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
		if len(inputText) > 500 {
			http.Error(w, "Eror : Bad Request ==> text overfflow", http.StatusBadRequest)
			return
		}

		banner := r.FormValue("choice")
		templwwwate := []string{}

		switch banner {
		case "standard":
			templwwwate = functions.ReadFile("banners/" + banner + ".txt")
		case "shadow":
			templwwwate = functions.ReadFile("banners/" + banner + ".txt")
		case "thinkertoy":
			templwwwate = functions.ReadFile("banners/" + banner + ".txt")

		default:
			http.Error(w, "Eror : Bad Request", http.StatusBadRequest)
			return

		}

		tretedtext := functions.TraitmentData(templwwwate, inputText)

		// Prepare the message to display
		message := /*"he " + inputText + "! " + "you are shosen " + banner*/ tretedtext

		// Render the home page with the message
		tmpl, err := template.ParseFiles("index.html")
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
	http.HandleFunc("/ascii-art", submitHandler) // Handle POST /submit
	fmt.Println("srever is running i  port 8080 ", ">>> http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
