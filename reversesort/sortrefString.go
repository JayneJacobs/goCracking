package reversesort

import (
	"fmt"
	"sort"
)

// RunStringRefSort runs reversesort for strings
func RunStringRefSort()  {
	inarray := []string{"Jayne","Jim","Scott","Rich","Elvis"}

	sort.Sort(sort.Reverse(sort.StringSlice(inarray)))
	fmt.Printf("This is the reverse sorted slice: %v\n",inarray)
	sort.Sort(sort.StringSlice(inarray))
	fmt.Printf("This is the  sorted slice: %v\n",inarray)
}