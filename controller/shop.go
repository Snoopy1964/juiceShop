package controller

import (
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/Snoopy1964/webapp/model"
	"github.com/Snoopy1964/webapp/viewmodel"
)

type shop struct {
	shopTemplate     *template.Template
	categoryTemplate *template.Template
}

func (s shop) registerRoutes() {
	http.HandleFunc("/shop", s.handleShop)
	http.HandleFunc("/shop/", s.handleShop)
}

func (s shop) handleShop(w http.ResponseWriter, r *http.Request) {
	categoryPattern, _ := regexp.Compile(`/shop/(\d+)`)
	matches := categoryPattern.FindStringSubmatch(r.URL.Path)

	log.Printf("handleShop URLPath: %v", r.URL.Path)
	log.Printf("handleShop matches: %v", matches)

	if len(matches) > 1 {
		categoryID, _ := strconv.Atoi(matches[1]) // warum matches[1] und nicht [0]?
		/*
			Die regex enthält "()" was einer Gruppe entspricht. matches enthält dann
			den kompletten String, also bei /shop/123 den String als 1. Element,
			und den Ausdruck in () - (\d+) als 2. Element
			-> hier also 123
			-> Ausprobieren auf https://regex101.com/
		*/
		s.handleCategory(w, r, categoryID)
	} else {

		categories := model.GetCategories()

		// log.Println(categories)

		vm := viewmodel.NewShop(categories)
		w.Header().Add("Content-Type", "text/html")
		s.shopTemplate.Execute(w, vm)
	}
}

func (s shop) handleCategory(w http.ResponseWriter, r *http.Request, categoryID int) {
	log.Printf("handleCategory - categoryID = %v\n", categoryID)
	products := model.GetProductsForCategory(categoryID)
	vm := viewmodel.NewShopDetail(products)
	s.categoryTemplate.Execute(w, vm)
}
