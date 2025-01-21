package main

import (
	"fmt"
	"net/http"
)

// Handler for delivering static files such as CSS
func staticFileServer() http.Handler {
	fileServer := http.FileServer(http.Dir("D:/projects/turing/21-01-2025/528377/turn2/modelA/static"))
	return http.StripPrefix("/static/", fileServer)
}

func generateHTML(title, message string) string {
	// Reference to the external CSS file in the HTML
	return fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>%s</title>
			<link rel="stylesheet" href="/static/style.css">
		</head>
		<body>
			<h1>%s</h1>
			<p>%s</p>
		</body>
		</html>`, title, title, message)
}

func handler(w http.ResponseWriter, r *http.Request) {
	html := generateHTML("Welcome Page", "This is a dynamically generated HTML page with external CSS.")
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, html)
}

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/static/", staticFileServer())
	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
