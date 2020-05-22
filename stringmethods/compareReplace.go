package stringmethods

import (
	"fmt"
	"strings"
)

// CompareReplace Containes Suffix Join Replace
func CompareReplace()  {
	// fmt.Println(strings.Contains("Jayne", "J"))
	fmt.Println(strings.Compare("Jayne", "Jayne"))
	fmt.Println(strings.Compare("Jayne", "jaYne"))
	fmt.Println(strings.Compare("Jayne", "Jayne"))
	// fmt.Println(strings.HasPrefix("Jayne", "Jay"))

	// fmt.Println(strings.HasSuffix("Jayne", "ne"))
	// fmt.Println(strings.Index("Jayne", "ne"))
	// fmt.Println(strings.Index("Jayne", "J"))
	// fmt.Println(strings.Join([]string{"Jayne", "Jacobs"}, " - "))
	// fmt.Println(strings.Join([]string{"Jayne", "Jacobs", "Jim"}, " - "))

	// fmt.Println(strings.Repeat("Jayne", 3))
	// fmt.Println(strings.Replace("jay", "j", "J", 1))
	// a := "jayne jacobs"
	// a = fmt.Sprintln(strings.Replace(a, "j", "J", 2))
	// a = fmt.Sprintln(strings.Replace(a, "bb", "b", 1))
	// fmt.Printf(a)
}