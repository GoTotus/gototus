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

	auth, err := t.Authenticated()
	if err != nil {
		fmt.Printf("Error checking authentication: %v\n", err)
		return
	}

	fmt.Printf("Authenticated: %v\n", auth.Authenticated)
	fmt.Printf("Description: %s\n", auth.Description)
	fmt.Printf("Token (masked): %s\n", auth.TokenSM)
	fmt.Println(auth)
}
