package main

import (
	"fmt"

	"github.com/GoTotus/gototus/totus"
)

func main() {
	t, err := totus.NewTotus("", "", "")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	bal, err := t.Balance()
	if err != nil {
		fmt.Printf("Error fetching balance: %v\n", err)
		return
	}

	fmt.Printf("Account Balance: %.2f Credits\n", bal.Credits)
	fmt.Println(bal)
}
