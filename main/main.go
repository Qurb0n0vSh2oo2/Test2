package main

import (
	"fmt"
	"test/pkg/bank/card"
	"test/pkg/bank/types"
)

func main() {
	result := types.Card{
		Activity: true,
		Balance:  10000,
	}

	fmt.Println(card.Card(&result))
}
