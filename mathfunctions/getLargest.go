package mathfunctions

import "fmt"
// GetLargestNum compares three inputs
func GetLargestNum()  {
fmt.Println("Please enter 3 numbers: ")
var first int
var second int
var third int
fmt.Scanln(&first)
fmt.Scanln(&second)
fmt.Scanln(&third)

if first >= second && first >= third{
	fmt.Println("The largest number is : ", first)
}
if second >= first && second >= third {
	fmt.Println("The lafgest number is: ", second)
}

if third >= first && third >= second {
	fmt.Println("The largest number is: ", third)
}

}

// CheckPalindrome uses modulus to compare 
func CheckPalindrome() {
	var inx, rev int
	// var rev int  = 0
	
	fmt.Println("\nEnter a number you think may be a palindrome: ")
	fmt.Scanln(&inx)
	originx := inx //to print the original number

	
	for {
		inxmod10 := inx % 10
		fmt.Println("inxmod10 = inx % 10: ", inxmod10)
		rev = rev * 10 + inxmod10
		fmt.Println("rev = inxmod10 + 10: ", rev)
		inx /= 10
		fmt.Println("inx /= 10: ", inx)
		
		if inx == 0 {
			break
		}
	}
	if originx == rev {
		fmt.Printf("The input number %v is a palindrome\n", originx)
		return
	}
	fmt.Printf("The input number %v is not a palindrome\n", originx)
}

// GetAverage from am input array
func GetAverage() {
	var tot, num, avg, ndx int
	var  inar [4]int
	fmt.Println("Enter the number of items in teh list\n")
	fmt.Scanln(&ndx)

	for i := 0; i < ndx; i++ {
		fmt.Println("\nEnter an input number: ")
		fmt.Scanln(&num)
		inar[i]=num
		tot += inar[i]
	}

	avg = tot/ndx
	fmt.Printf("The average is:  %v\n", avg)
}
