package server

import (
	"fmt"
	"html/template"
	"net/http"

	"ascii-art-web/functions"
)

type PageData struct {
	Message string
}

const maxInputTextLength = 500

var temple01 *template.Template

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Error 404: NOT FOUND", http.StatusNotFound)
		return
	}
	tmpl, err := template.ParseFiles("index.html")
	temple01 = tmpl
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ParseForm(r *http.Request) (string, string, error) {
	err := r.ParseForm()
	if err != nil {
		return "", "", err
	}
	inputText := r.FormValue("inputText")
	if len(inputText) > maxInputTextLength {
		return "", "", fmt.Errorf("input text exceeds %d characters", maxInputTextLength)
	}
	banner := r.FormValue("choice")
	return inputText, banner, nil
}

func ReadBannerTemplate(banner string) ([]string, error) {
	switch banner {
	case "standard", "shadow", "thinkertoy":
		return functions.ReadFile("banners/" + banner + ".txt"), nil
	default:
		return nil, fmt.Errorf("error: 300 invalid banner choice: %s", banner)
	}
}

func TreatData(templ []string, inputText string) string {
	return functions.TraitmentData(templ, inputText)
}

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Error 405: Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	inputText, banner, err := ParseForm(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	templ, err := ReadBannerTemplate(banner)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	treatedText := TreatData(templ, inputText)

	test := temple01
	err = test.Execute(w, PageData{Message: treatedText})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
