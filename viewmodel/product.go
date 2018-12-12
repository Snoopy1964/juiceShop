package viewmodel

import "github.com/Snoopy1964/webapp/model"

// Product is the home data object for the templates
type Product struct {
	Name             string
	DescriptionLong  string
	DescriptionShort string
	PricePerLiter    float32
	PricePer10Liter  float32
	IsOrganic        bool
	Origin           string
	ImageURL         string
	ID               int
}

func product2VM(p model.Product) Product {

	return Product{
		Name:             p.Name,
		DescriptionShort: p.DescriptionShort,
		DescriptionLong:  p.DescriptionLong,
		PricePerLiter:    p.PricePerLiter,
		PricePer10Liter:  p.PricePer10Liter,
		Origin:           p.Origin,
		IsOrganic:        p.IsOrganic,
		ImageURL:         p.ImageURL,
		ID:               p.ID,
	}
}
