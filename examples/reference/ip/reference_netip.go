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
	ref := t.Reference()

	fmt.Println("Your Public NetIP ...")
	result, err := ref.NetIP()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println(result)

	fmt.Println("Cloudflare 1.1.1.1 ...")
	result, err = ref.NetIP4("1.1.1.1")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println(result)

	fmt.Println("Cloudflare ip6 for previous 1.1.1.1: 2606:4700:4700::1111 ...")
	result, err = ref.NetIP6("2606:4700:4700::1111")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println(result)
}
