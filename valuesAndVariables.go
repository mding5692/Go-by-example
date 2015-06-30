
package main

// don't forget the input/output library
import "fmt"

// nor your main function
func main() {
  
  fmt.Println("this is a number:", 2+2) 
  // would print out "this is a number: 4"
  
  fmt.Println("Strings" + "can" + "be" + "concatenated" + "together")
  
  // booleans still the same
  fmt.Println(true && false)
  fmt.Println(true || false)
  fmt.Println(!true)
  
  
  
}

/*
Use this to compile:
go run valuesAndVariables.go
*/
