//before making code change do...
// go mod init github.com/imattf/golang-project/02-modules/05-hands-on-exercises/02-hand-on
//now make code changes

package main

import (
	"fmt"

	"rsc.io/quote/v3"
)

func main() {
	fmt.Println(quote.GlassV3())
}
