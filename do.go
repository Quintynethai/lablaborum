package main

import (
    "fmt"
)

// Person struct represents a person with a name and age.
type Person struct {
    Name string
    Age  int
}

func main() {
    // Initializing a Person object
    person := Person{
        Name: "Alice",
        Age:  30,
    }

    // Printing out the details of the person
    fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
}
