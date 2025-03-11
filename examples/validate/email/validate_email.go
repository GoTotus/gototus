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
	val := t.Validate()

	emails := []string{
		"invalid@gototus.com",
		"sdfsdf@sdfsdfsdfsfs.fdfsfs.fdfsds",
		"temporary@blondmail.com",
		"info@x.com",
	}
	for _, email := range emails {
		result, err := val.Email(email, totus.CheckLevelL4Dbs)
		if err != nil {
			fmt.Printf("Error validating %s: %v\n", email, err)
			continue
		}
		fmt.Printf("email %s: good email? %v; with score: %d/100\n", email, result.IsValid(), result.Score())
		fmt.Println(result)
	}
}
