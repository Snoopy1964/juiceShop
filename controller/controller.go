package controller

import (
	"html/template"
	"net/http"
)

var (
	homeController home
	shopController shop
)

func Startup(templates map[string]*template.Template) {

	homeController.homeTemplate = templates["home.html"]
	homeController.loginTemplate = templates["login.html"]
	homeController.standLocatorTemplate = templates["stand_locator.html"]
	homeController.registerRoutes()

	shopController.shopTemplate = templates["shop.html"]
	shopController.registerRoutes()

	http.Handle("/img/", http.FileServer(http.Dir("web/public")))
	http.Handle("/css/", http.FileServer(http.Dir("web/public")))
}
