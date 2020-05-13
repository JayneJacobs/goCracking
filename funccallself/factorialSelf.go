package funccallself

// Factorial is a function that calls itself
func Factorial(n int)  int {
	if n == 0 {
		return 1
	}

	n = n * Factorial(n-1)
	// fmt.Println(n)
	return n
}

// LenString returns the lenth of an input string
func LenString(instr string) int {
	return len(instr)
}