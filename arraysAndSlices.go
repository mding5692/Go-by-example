package main

import "fmt"

func main() {

  var array [10]int // makes an array with 10 indexes of integers
  var strArr [3]string // makes an array with 3 indexes of strings
  
  fmt.Println(array) // prints out whole array
  array[3] = 100 // sets third index of array to 100
  
  // another way to declare arrays
  arr := [2]int{1,3} // fills out content of arrays with declaration
  
  fmt.Println(len(array)) // prints out length of array
  
  // multidimensional array in go
  var multiArray [2][2]int
  
  // fills out the above array
  for i:= 0; i < 2; i++ {
    for j:=0; j<2; j++ {
      multiArray[i][j] = 2
    }
  }
  
  /** Go also supports Slices **/
  /* Slices are only typed by the elements they contain (think strings or ints) and don't need to specify length
  unlike arrays */
  
  // you gotta use the make keyword to build a slice of non-zero length
  slice := make([]int, 4) // builds a slice of 4 indices and takes integers
  // notice the above is like calling a make function which takes two parameters, one of the slice/type and another of length
  sliceWithoutLength := []string // creates empty slice
  declaredString := []string{"g", "h", "i"} // creates slice with already declared values
  
  // you have more functionality with slices than arrays like the append
  slice = append(slice, 2) // note that append is builtin and returns a slice type back so you need a variable to reference the return value
  
  // making a multidimensional slice 
  multiSlice := make([][]int, 3) // length of inner slice can vary
  
  // filling above slice in 
      for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            multiSlice[i][j] = i + j
        }
    }
    
    // still looks similar when printed out
    fmt.Println(multiSlice)
  }
