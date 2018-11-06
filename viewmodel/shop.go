package viewmodel

import (
	"fmt"

	"github.com/snoopy1964/webapp/model"
)

// Shop structure
type Shop struct {
	Title      string
	Active     string
	Categories []Category
}

// Category structure
type Category struct {
	URL             string
	ImageURL        string
	Title           string
	Description     string
	IsOrientedRight bool
}

// NewShop is the contructor of a Shop
func NewShop(categories []model.Category) Shop {
	result := Shop{
		Title:      "Lemonade Stand Supply - Shop",
		Active:     "shop",
		Categories: make([]Category, len(categories)),
	}

	// log.Println(categories)

	for i := 0; i < len(categories); i++ {
		// log.Printf("Index i = %v\n", i)
		vm := category2VM(categories[i])
		vm.IsOrientedRight = i%2 == 1
		result.Categories[i] = vm
	}
	return result
}

func category2VM(c model.Category) Category {
	return Category{
		URL:         fmt.Sprintf("/shop/%v", c.ID),
		ImageURL:    c.ImageURL,
		Title:       c.Title,
		Description: c.Description,
	}
}
