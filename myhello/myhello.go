package hello

import (
	"fmt"

	"rsc.io/quote"
)

func TestHello() {
	fmt.Println("this is my test module")
	fmt.Println(quote.Hello())
}
