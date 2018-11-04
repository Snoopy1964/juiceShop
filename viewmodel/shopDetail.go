package viewmodel

// ShopDetail structure
type ShopDetail struct {
	Title    string
	Active   string
	Products []Product
}

func NewShopDetail() ShopDetail {
	var result ShopDetail
	result.Title = "Lemonade Stand Supply - Juice Shop"
	result.Active = "shop"
	lemonJuice := MakeLemonJuiceProduct()
	appleJuice := MakeAppleJuiceProduct()
	watermelonJuice := MakeWatermelonJuiceProduct()
	kiwiJuice := MakeKiwiJuiceProduct()
	mangosteenJuice := MakeMangosteenJuiceProduct()
	orangeJuice := MakeOrangeJuiceProduct()
	pinappleJuice := MakePineappleJuiceProduct()
	strawberryJuice := MakeStrawberryJuiceProduct()

	result.Products = []Product{lemonJuice, appleJuice, watermelonJuice, kiwiJuice, mangosteenJuice, orangeJuice, pinappleJuice, strawberryJuice}

	return result
}
