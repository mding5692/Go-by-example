
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
  
  // reminds me of Javascript variables
  var str string = "This is a string variable and this is the number "
  
  // and yep, those are data types behind the variable name
  var num int = 1
  
  // boolean seems the same
  var truth = true
  
  // printing it out now
  fmt.Println(str + num)
  
  // without any value binded to a int variable, it will be 0
  var initNum int
  if (initNum == 0 ) { fmt.Println(truth) }
  // ^ did a decision up there
  
  shortHand := "short way of declaring variable"
  // equivalent to saying:
  // var shortHand string = "short way ..."
  
}

/*
Use this to compile:
go run valuesAndVariables.go
*/
