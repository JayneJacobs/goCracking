package stringmethods

import (
	"fmt"
	"strings"
)

// ConvertCase Upper and Lower
func ConvertCase()  {
	fmt.Println(strings.Split("Jayne-Jacobs", "-"))
	b := fmt.Sprintln(strings.ToLower("base32"))
	fmt.Printf(b)
	b = fmt.Sprintln(strings.ToUpper(b))
	fmt.Printf(b)
}


// LenString returns the lenth of an input string
func LenString(instr string) int {
	return len(instr)
}