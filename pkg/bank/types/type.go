package types

type Money int

type dirams Money

type Currency string

type Transaction struct {
	Amount Money
	Type   string
	//Time   time.Time
}

type Card struct {
	Name       string
	Pin        string
	Balance    Money
	Currency   Currency
	Type       string
	Activity   bool
}

var LocalCard = Card{
	Balance:  9000_00,
	Currency: Dollars,
	Type:     Gold,
	Activity: Active,
}

const (
	Platinum = "Platinum"
	Gold     = "Gold"
	Silver   = "Silver"
)

const (
	Somoni  = "TJS"
	Rubls   = "RUB"
	Dollars = "USD"
	Euro    = "EUR"
)

const (
	Active   = true
	InActive = false
)
