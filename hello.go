package hellolib

import "fmt"

func Greeting() string {
	return fmt.Sprintln("Hello, World!")
}

func GreetingV2() string {
	return fmt.Sprintln("Hello, World2!")
}