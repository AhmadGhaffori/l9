package cards

import (
	"L9/pkg/bank/types"
	"fmt"
)

func ExampleTotal() {
	card := []types.Card{
		{
			Balance: 565,
			Active:  true,
		},
		{
			Balance: 1235,
			Active:  true,
		},
	}
	fmt.Println(card)
}
