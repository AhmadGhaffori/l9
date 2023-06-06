package main

import (
	"L9/pkg/bank/cards"
	"L9/pkg/bank/types"
	"fmt"
)

func main() {
	card := []types.Card{
		{
			Balance: 5667,
			Active:  true,
		},
		{
			Balance: -67,
			Active:  true,
		},
		{
			Balance: 2,
			Active:  false,
		},
	}
	chose := cards.Total(card)
	fmt.Println(chose)
}
