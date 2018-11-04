package viewmodel

// Base is the base data object for the  templates
type Base struct {
	Title string
}

// NewBase is the contructor function of the Base structure
func NewBase() Base {
	return Base{
		Title: "Lemonade Stand Supply",
	}
}
