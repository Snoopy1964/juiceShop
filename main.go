package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/Snoopy1964/webapp/middleware"
	"github.com/Snoopy1964/webapp/model"

	"github.com/Snoopy1964/webapp/controller"

	_ "github.com/lib/pq"
)

func main() {
	// http.ListenAndServe(":8000", http.FileServer(http.Dir("webapp/public")))
	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	log.Println("Working Directory: ", dir)

	templates := populateTemplates()
	db := connectToDatabase()
	log.Printf("Connect to database: %v", db)
	defer db.Close()
	controller.Startup(templates)
	//http.ListenAndServeTLS(":8000", "cert.pem", "key.pem", &middleware.TimeoutMiddleware{new(middleware.GzipMiddleware)})
	http.ListenAndServe(":8000", &middleware.TimeoutMiddleware{new(middleware.GzipMiddleware)})
}

func connectToDatabase() *sql.DB {
	db, err := sql.Open("postgres", "postgres://user:password@localhost/accounts?sslmode=disable")
	if err != nil {
		log.Fatalln(fmt.Errorf("Unable to connect to database: %v", err))
	}
	model.SetDatabase(db)
	return db
}

func populateTemplates() map[string]*template.Template {
	// log.Println("Processing template files ...")
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

		// log.Printf("Processing Template file: %v", file.Name())

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
	// log.Println("Processing template files finished!")

	return result
}
