package main

import (
	"fmt"
	"net/http"

	"ascii-art-web.github.io/server"
)

func main() {
	http.HandleFunc("/", server.Home)
	http.HandleFunc("/ascii-art", server.SubmitHandler)
	fmt.Println("Server is running on port 8080", ">>> http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
