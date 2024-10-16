package server

import (
	"fmt"
	"html/template"
	"net/http"

	"ascii-art-web/functions"
)

type ParsedFiles struct {
	Temple01      *template.Template
	ErrorTemplate *template.Template
}

const maxInputTextLength = 500

func (vars *ParsedFiles) Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		vars.ErrorTemplate.Execute(w, "MethodNotAllowed")
		return
	}

	if r.URL.Path != "/" {

		w.WriteHeader(http.StatusNotFound)
		vars.ErrorTemplate.Execute(w, "NOT FOUND")
		return
	}
	err := vars.Temple01.Execute(w, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		vars.ErrorTemplate.Execute(w, "Internal Server Error")
		return
	}
}

func ReadBannerTemplate(banner string) ([]string, error, bool) {
	switch banner {
	case "standard", "shadow", "thinkertoy":
		return functions.ReadFile("banners/" + banner + ".txt")
	default:
		return nil, fmt.Errorf("error: 400 invalid banner choice: %s", banner), false
	}
}

func (vars *ParsedFiles) SubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		vars.ErrorTemplate.Execute(w, "Method Not Allowed")
		return
	}

	// extract the input text and ther banner fromthe request
	inputText := r.FormValue("inputText")
	banner := r.FormValue("choice")

	if len(inputText) > maxInputTextLength {
		w.WriteHeader(http.StatusBadRequest)
		vars.ErrorTemplate.Execute(w, "input text exceeds 500 characters")
		return
	}

	if len(inputText) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		vars.ErrorTemplate.Execute(w, "Enter a text")
		return
	}
	templ, err, status := ReadBannerTemplate(banner)
	if err != nil {
		if status {
			w.WriteHeader(http.StatusInternalServerError)
			vars.ErrorTemplate.Execute(w, "Internal Server Error")
			return
		} else {
			w.WriteHeader(http.StatusNotFound)
			vars.ErrorTemplate.Execute(w, " BANNER NOT FOUND")
			return
		}
	}

	if len(templ) != 856 {
		w.WriteHeader(http.StatusInternalServerError)
		vars.ErrorTemplate.Execute(w, "Internal Server Error")
		return
	}
	treatedText, err := functions.TraitmentData(templ, inputText)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		vars.ErrorTemplate.Execute(w, "Our ascii program  do not suport non-ascii printable characters")
		return
	}
	err = vars.Temple01.Execute(w, treatedText)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		vars.ErrorTemplate.Execute(w, "Internal Server Error")
		return
	}
}

func (vars *ParsedFiles) CssHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		vars.ErrorTemplate.Execute(w, "Method Not Allowed")
		return
	}
	http.ServeFile(w, r, "css/style.css")
}

func (vars *ParsedFiles) CssErrHundle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		vars.ErrorTemplate.Execute(w, "Method Not Allowed")
		return
	}
	http.ServeFile(w, r, "css/error.css")
}
