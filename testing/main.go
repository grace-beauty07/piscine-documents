package main

import (
	"fmt"
	"piscine_go"
)

func main() {

	link := &piscine_go.List{}

	piscine_go.ListPushFront(link, "Hello")
	piscine_go.ListPushFront(link, "man")
	piscine_go.ListPushFront(link, "how are you")

	it := link.Head
	for it != nil {
		fmt.Print(it.Data, " ")
		it = it.Next
	}
	fmt.Println()
}
