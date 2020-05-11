package stringprnt

import "fmt"

// MyltiLineStr two types of multilin string tic marks and \n
func MyltiLineStr()  {
	multln1 := `firstline
	secondlin
	thirdline
	fourthline`

	multiline2 := "firstline \n  secondline \n thirdline \n"

		fmt.Println("First style", multln1)
		fmt.Println("Second style", multiline2)

		multiline3 := "firstline \n" + "secondline \n" + "thirdline \n"

		fmt.Println("Third style\n", multiline3)

}

// UseScanfln various uses of Scanf Scan and Sscanf
func UseScanfln()  {
	fstint := 0
	scdint := 0
	thrdint := 0

	fmt.Sscan("222\n333", &fstint, &scdint)
	fmt.Printf("first = %d \n second = %d \n", fstint, scdint)
	

	fmt.Sscanf("(999,888)", "(%d,%d)", &fstint, &scdint)
	fmt.Printf("firstnow = %d \n secondnow = %d\n", fstint, scdint)


	fmt.Sscanln("666 667 111\n", &fstint, &scdint, &thrdint)
	fmt.Printf("firstthd = %d \n secondthd = %d\n thirdint = %d \n", fstint, scdint, thrdint)
}