package factorials

import "fmt"

// var factin int
// var fact uint64 = 1
 
func factorialLoop(number int) uint64 {
	var fact uint64 
	fact = 1
	if number < 0 {
		fmt.Println("The number was less than 0")
		return 0
	}
	for i := 1; i <= number; i++ {
		 fact *= uint64(i)
		 fmt.Printf("now %v\n", fact)
	}
	return fact
}

// RunFactorialLoop requests an input int and runs Factorial Loop 
func RunFactorialLoop(){
	var factin int
	fmt.Println("Enter a positive integer greater than 0: ")
	fmt.Scan(&factin)
	factout := factorialLoop(factin)
	fmt.Printf("\nThe Factorial of %v is : %v\n", factin, factout)
}