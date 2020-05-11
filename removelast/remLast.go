package removelast

import "fmt"

// RemoveL removes last element of array using len
func RemoveL()  {
	myArray := []int{5,3,4,1}
	fmt.Println(myArray)

	lstEl := myArray[len(myArray)-1]
	fmt.Printf("The last element is %v \n", lstEl)

	fstEl := myArray[:0]
	fmt.Printf("The first element is %v \n", fstEl)

	removedArray := myArray[:len(myArray)-1]
	fmt.Println(removedArray)
}