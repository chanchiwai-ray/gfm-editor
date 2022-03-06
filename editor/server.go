package editor

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//go:embed static/index.html
var fs embed.FS
var templates = template.Must(template.ParseFS(fs, "static/index.html"))

func SimpleHandler(file *File, path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			err := templates.ExecuteTemplate(w, "index.html", file)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		case "POST":
			content := r.FormValue("content")
			file.Content = []byte(content)
			err := file.Save()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			log.Printf("Saved changes to %s", path)
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}
}

func Serve(file *File, path string, port string) {
	http.Handle("/", SimpleHandler(file, path))
	log.Printf("Listening on http://localhost:%s ...", port)
	log.Fatal(http.ListenAndServe(string(fmt.Sprintf(":%s", port)), nil))
}
