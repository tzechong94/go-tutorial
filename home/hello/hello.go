package main

// declare a main package. package is a way to group functions, and it's made up of all the files in the same directory

import (
	"fmt"

	"log"

	"example.com/greetings"
)

// import fucnctions for formatting texts and prnting to console

func main() {
	    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Gladys", "Sam", "Darren"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}

//func is function
// When you ran go mod tidy, it located and downloaded the rsc.io/quote module that contains the package you imported. By default, it downloaded the latest version -- v1.5.2.

