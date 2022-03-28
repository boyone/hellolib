package main

import (
	"fmt"

	"github.com/boyone/hellolib"
	v2 "github.com/boyone/hellolib/v2"
)

func main() {
	fmt.Println(hellolib.Greeting())
	fmt.Println(v2.Greeting())
}
