package factorials

import "fmt"

// n *(n-1)! a number that is the same number*the factorial of itself minus1.
//Written as N!
//4! = 4 *3 *2 *1



// MyFactorial uses a for loop to calculae the factorial of a number and returns the factorial as a uint64
func MyFactorial(innum int) uint64 {
	var fact uint64 = 1
	fmt.Println("Something")

	for i := innum; i >= 1; i-- {
		ui := uint64(i)
		fmt.Println("uint(innum): ",ui)
		fact *= ui
		fmt.Println("The factorial is: ", fact)
	}
	return fact
}

// RunMyFactorial requests the number
func RunMyFactorial()  {
	var inn int
	fmt.Println("Enter an even number greater than 0: ")
	fmt.Scanln(&inn)
	outfact := MyFactorial(inn)
	fmt.Printf("This is the factorial of %v: %v \n", inn, outfact)
}