package stringmeth

import (
	"fmt"
	"strings"
)

// StringContainsAny uses strins.ContainsAny
 func StringContainsAny()  {
	 fmt.Println(strings.ContainsAny("Jayne", "a"))
	 fmt.Println(strings.Contains("Jayne", "o"))
 }

// StringContains uses Contains matches complete
 func StringContains()  {
	 fmt.Println(strings.Contains("Jayne","o"))
	 fmt.Println(strings.Contains("Jayne", "Jayne"))
 }

// StringCount must be in the right order
 func StringCount()  {
	 fmt.Println("Jayne Jacobs", "Ja")
	 fmt.Println("Jayne Jacobs", "bocaJ" )
	 fmt.Println(strings.EqualFold("JAyne","jayne"))
 }