package viewmodel

import "github.com/Snoopy1964/webapp/model"

// ShopDetail structure
type ShopDetail struct {
	Title    string
	Active   string
	Products []Product
}

func NewShopDetail(products []model.Product) ShopDetail {
	result := ShopDetail{
		Title:    "Lemonade Stand Supply - Juice Shop",
		Active:   "shop",
		Products: []Product{},
	}

	for _, p := range products {
		result.Products = append(result.Products, product2VM(p))
	}
	return result
}
