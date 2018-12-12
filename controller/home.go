package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/Snoopy1964/webapp/viewmodel"
)

type home struct {
	homeTemplate         *template.Template
	loginTemplate        *template.Template
	standLocatorTemplate *template.Template
}

func (h home) registerRoutes() {
	http.HandleFunc("/", h.handleHome)
	http.HandleFunc("/home", h.handleHome)
	http.HandleFunc("/login", h.handleLogin)
	http.HandleFunc("/stand-locator", h.handleStandLocator)
}

func (h home) handleHome(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewHome()
	h.homeTemplate.Execute(w, vm)
}

func (h home) handleStandLocator(w http.ResponseWriter, r *http.Request) {
	log.Println("handleSTandLocator called: ", r.URL, r.RequestURI)
	vm := viewmodel.NewStandLocator()
	h.homeTemplate.Execute(w, vm)
}

func (h home) handleLogin(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewLogin()

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Println(fmt.Errorf("Error logging in: %v", err))
		}
		email := r.Form.Get("email")
		password := r.Form.Get("password")
		if email == "ralf@ehret-family.com" && password == "12345" {
			http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
		} else {
			vm.Email = "please enter valid user"
			vm.Password = "Please enter valid password"
		}
	}

	h.loginTemplate.Execute(w, vm)
}
