package main //(a package is a way to group functions, and it's made up of all the files in the same directory

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	println("hello world")
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
