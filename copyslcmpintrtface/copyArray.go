package copyslcmpintrtface

import "fmt"

// CopyArrayToSlice copies an array to a slice
func CopyArrayToSlice()  {
	myarray := [5]string{"Jayne","Jim","Scott","Rich","Elvis"}
	copofArrayslc := make([]string, 5)
	copofArrayslc = myarray[1:4]
	fmt.Println(copofArrayslc)
}