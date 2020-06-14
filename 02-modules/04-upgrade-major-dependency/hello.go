package hello

//new...
import (
	// remove the previous version
	//"rsc.io/quote"
	"rsc.io/quote/v3"
)

//Hello is a...
func Hello() string {
	//return "Hello, world."
	//return quote.Hello()
	return quote.HelloV3()
}

//Proverb is a ...
func Proverb() string {
	return quote.Concurrency()
}
