package app

import (
	"html/template"
	"log"
	"net/http"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/templates/index.html"))
	tmpl.Execute(w, nil)
}

func handleBoardIndex(w http.ResponseWriter, r *http.Request) {
	board := r.PathValue("board")
	tmpl := template.Must(template.ParseFiles("web/templates/board.html"))
	data := map[string]string{
		"Board": board,
	}
	tmpl.Execute(w, data)
}

func handleBoardCatalog(w http.ResponseWriter, r *http.Request) {
	board := r.PathValue("board")
	tmpl := template.Must(template.ParseFiles("web/templates/catalog.html"))
	data := map[string]string{
		"Board": board,
	}
	tmpl.Execute(w, data)
}

func handleThread(w http.ResponseWriter, r *http.Request) {
	board := r.PathValue("board")
	threadId := r.PathValue("threadId")
	tmpl := template.Must(template.ParseFiles("web/templates/thread.html"))
	data := map[string]string{
		"Board":    board,
		"ThreadId": threadId,
	}
	tmpl.Execute(w, data)
}

func Start() {
	http.HandleFunc("GET /", handleIndex)
	http.HandleFunc("GET /{board}/", handleBoardIndex)
	http.HandleFunc("GET /{board}/catalog", handleBoardCatalog)
	http.HandleFunc("GET /{board}/{threadId}", handleThread)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
