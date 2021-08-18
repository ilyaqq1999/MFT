package models
//results, counttoshowint, activepageint, pagesarr, pages

type Shops struct {
	Results [] Result `json:"Results"`
	Counttoshowint int `json:"Counttoshowint"`
	Activepageint int	`json:"Activepageint"`
	Pagesarr [] int	`json:"Pagesarr"`
	Pages int `json:"Pages"`
}

