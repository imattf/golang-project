package main

import (
	"fmt"

	"github.com/imattf/quotes"
)

func main() {

	// for i, v := range sayit.Favs() {
	// 	fmt.Println(i, v)
	// }

	for i, v := range quotes.Favs() {
		fmt.Println(i, v)
	}

}
