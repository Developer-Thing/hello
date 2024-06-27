package hello

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hello, %v. Welcome !", name)
	return message
}

func Hi(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome !", name)
	return message
}