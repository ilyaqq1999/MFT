package models

type Shop struct {
	ID          int
	Merchant    int
	Type        int
	Name        string
	Address     string
	Phone       string
	ContactName string
	Email       string
	SearchName  string
	Blocked     bool
}
