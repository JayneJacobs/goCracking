package copyslcmpintrtface

import "fmt"

// CopySlices ...
func CopySlices(inslc []int) []int {
	outslc := make([]int, 2)
	copy(outslc, inslc)
	return outslc

}



// Copies a map that is passed and returns it. 
// func CopyMap(inmp map[string]int)  map[string]int {
// 	outmap := map[string]int{}
// 	copy(outmap, inmp)
// }

// UseCopyExample creates a slice and a map to copy using the copy function
// the dst slc has to have a declared capacity. It will copy the number of elements for its capacity
func UseCopyExample()  {
	fmt.Println("#####Slices#####")
	myinslc := []int{3,4,6,7,8}
	outslc := make([]int, 6)
	myinslc = append(myinslc, 1)
	numcopied := copy(outslc, myinslc)
	myretslc := CopySlices(outslc)
	fmt.Println(outslc, numcopied)
	fmt.Println(myinslc)
	fmt.Println(myretslc)
fmt.Println("#####Maps#####")

}
// CopyMap takes a map and returns a map
func CopyMap()  {
	IsFemale := map[string]bool{"Jayne":true, "Jim":false}
	fmt.Printf("is Jayne female %v Jime?  %v \n", IsFemale["Jayne"], IsFemale["Jim"])
	CPofIsFemale := make(map[string]bool)
	for name, value := range IsFemale {
		CPofIsFemale[name] = value
		fmt.Printf("Is %v female? %v \n",name, value)
	}
	fmt.Printf("here is the copy:  %v\n", CPofIsFemale)
}