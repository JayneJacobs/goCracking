package concatslices

import "fmt"

// ConcatonateSl concatonates three slices 
func ConcatonateSl()  {
	Sl1 := []int{1,2,3,4}
	Sl2 := []int{5,6,7,8}
	Sl3 := []int{9,10,11}

	Sl1 = append(Sl1, Sl2...)
	Sl1 = append(Sl1, Sl3...)

	fmt.Printf("This is the contents of Sl1 %v \n This is capacity of Sl1 %v\n", Sl1, cap(Sl1))
}