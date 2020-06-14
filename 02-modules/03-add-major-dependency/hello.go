package hello

//new...
import (
	"rsc.io/quote"
	quoteV3 "rsc.io/quote/v3"
)

//Hello is a...
func Hello() string {
	//return "Hello, world."
	return quote.Hello()
}

//Proverb is a ...
func Proverb() string {
	return quoteV3.Concurrency()
}
