package types

type Money int
type Currency string

const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

type PAN string

type Card struct {
	ID         int
	PAN        PAN
	MinBalance Money
	Balance    Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
}

type Payment struct {
	Type    string
	Number  PAN
	Balance Money
}
