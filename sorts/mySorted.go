package sorts

import (
	"fmt"
	"sort"
)

// MySortsInt sorts then searches ints
func MySortsInt() {
	var a []int

	t := 1
	fmt.Println("Enter the number of elements: ")
	fmt.Scanln(&t)

	num := 0
	for i := 1; i <= t; i++ {
		fmt.Println("Enter the number: \n")
		fmt.Scanln(&num)
		a = append(a, num)
	}

	sorted := sort.IntsAreSorted(a)
	fmt.Printf("\nInts sorted: %v\n", sorted)

	if !sorted {
		sort.Ints(a)
	}

	fmt.Println("Sorted array : ", a)
	fmt.Println(sort.SearchInts(a, 3))

}

// MySortsStrings sorts input strings in a slice
func MySortsStrings()  {
	var a []string
	t := 1
	fmt.Println("Enter the number of items: ")
	fmt.Scanln(&t)
	instr := ""
	for i := 1; i <= t; i++ {
		fmt.Println("Enter the string", 1)
		fmt.Scanln(&instr)
		a = append(a, instr)
	}
	isSorted := sort.StringsAreSorted(a)
	if !isSorted {
		sort.Strings(a)
	}
	var schstr string
	fmt.Println("Enter a search string: ")
	fmt.Scanln(&schstr)
	fmt.Println("Check if strings are sorted: ", isSorted)
	fmt.Println("The Sorted strings", a)
	fmt.Printf("Searching for string %v in %v index is %v: \n",  schstr, a,  sort.SearchStrings(a, schstr))
}