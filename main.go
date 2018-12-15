package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/Snoopy1964/webapp/controller"
	"github.com/Snoopy1964/webapp/middleware"
)

func main() {
	// http.ListenAndServe(":8000", http.FileServer(http.Dir("webapp/public")))
	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	log.Println("Working Directory: ", dir)

	templates := populateTemplates()
	controller.Startup(templates)
	http.ListenAndServe(":8000", &middleware.TimeoutMiddleware{new(middleware.GzipMiddleware)})
}

func populateTemplates() map[string]*template.Template {
	log.Println("Processing template files ...")
	result := make(map[string]*template.Template)
	const basePath = "src/github.com/snoopy1964/webapp/templates"
	layout := template.Must(template.ParseFiles(basePath + "/_layout.html"))
	template.Must(layout.ParseFiles(basePath+"/_header.html", basePath+"/_footer.html"))

	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Failed to open template blocks folder: " + err.Error())
	}

	files, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to load content files of content directory: " + err.Error())
	}
	for _, file := range files {

		log.Printf("Processing Template file: %v", file.Name())

		f, err := os.Open(basePath + "/content/" + file.Name())
		if err != nil {
			panic("Failed to open template '" + file.Name() + "'")
		}
		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic("Failed to read content of file: " + file.Name())
		}
		f.Close()

		tmpl := template.Must(layout.Clone())
		_, err = tmpl.Parse(string(content))
		if err != nil {
			panic("Failed to parse contents of '" + file.Name() + "' as template")
		}
		result[file.Name()] = tmpl
	}
	log.Println("Processing template files finished!")

	return result
}
