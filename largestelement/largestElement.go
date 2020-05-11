package largestelement

import "fmt"

// LargestElArray is here
func LargestElArray()  {
	var myArray []int
	innum := 0
	t := 1
	fmt.Printf("Please enter the number of values: \n")
	fmt.Scanln(&t)
	for i := 1; i <= t; i++ {
		fmt.Println("Enter value ",i)
		fmt.Scanln(&innum) 
		myArray = append(myArray, innum)
	}
	for x := 1; x <= t-1; x++ {
		fmt.Println(myArray[1], "compared to ", myArray[x])
		if (myArray[1] < myArray[x]) {
			myArray[1] = myArray[x]
		}
	}
	fmt.Println("Greatest value: ", myArray[1])
}