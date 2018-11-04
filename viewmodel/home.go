package viewmodel

// Base is the base data object for the  templates
type Home struct {
	Title  string
	Active string
}

// NewBase is the contructor function of the Base structure
func NewHome() Home {
	return Home{
		Title:  "Lemonade Stand Supply",
		Active: "home",
	}
}
