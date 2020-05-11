package reversesort

import (
	"fmt"
	"sort"
)

// RunrevSort uses sort package
func RunrevSort()  {
	inslc := []int{1,2,3,7,6,5,4,3}
	sort.Sort(sort.Reverse(sort.IntSlice(inslc)))
	fmt.Printf("Sorted slice %v \n", inslc)
}
