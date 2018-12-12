package viewmodel

type StandLocator struct {
	Title  string
	Active string
}

func NewStandLocator() StandLocator {
	result := StandLocator{
		Active: "standlocator",
		Title:  "Lemonade Stand Supply - Stand Locator RE",
	}
	return result
}
