package main

import (
	"html/template"
	"net/http"
)

// Define a struct to hold the data for the template
type PageData struct {
	Title string
	Body  string
}

// Handler for the main page (GET /)
func mainPageHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "Welcome to My ASCII Art Server",
		Body:  "Enter your text and choose a banner below!",
	}

	tmpl := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>{{.Title}}</title>
	</head>
	<body>
		<h1>{{.Title}}</h1>
		<p>{{.Body}}</p>
		<form action="/ascii-art" method="POST">
			<label for="text">Text:</label>
			<input type="text" id="text" name="text" required>
			<br>
			<label for="banner">Banner:</label>
			<select id="banner" name="banner">
				<option value="simple">Simple</option>
				<option value="bold">Bold</option>
				<option value="fancy">Fancy</option>
			</select>
			<br>
			<input type="submit" value="Generate ASCII Art">
		</form>
	</body>
	</html>
	`

	t := template.Must(template.New("main").Parse(tmpl))
	t.Execute(w, data)
}

// Handler for the ASCII art endpoint (POST /ascii-art)
func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	// ad smone ref
	text := r.FormValue("text")
	banner := r.FormValue("banner")

	// Here you could add logic to create ASCII art based on the text and banner.
	// For demonstration, we'll just respond with the received values.
	response := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>ASCII Art Result</title>
	</head>
	<body>
		<h1>Your ASCII Art</h1>
		<p>Text: ` + text + `</p>
		<p>Banner: ` + banner + `</p>
		<p>This is where your ASCII art would be displayed!</p>
		<a href="/">Back to main page</a>
	</body>
	</html>
	`

	w.Write([]byte(response))
}

func main() {
	http.HandleFunc("/", mainPageHandler)          // Handle GET /
	http.HandleFunc("/ascii-art", asciiArtHandler) // Handle POST /ascii-art

	http.ListenAndServe(":8080", nil) // Start server on port 8080
}
