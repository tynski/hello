package main

import ( 
    "fmt"

    "github.com/tynski/greetings"
)


func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Barti")
	fmt.Println(message)
}
