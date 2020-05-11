package mathfunctions

import "fmt"

// MyGetAverage gets the average by building a struct
func MyGetAverage() {
var innum, totn, sum, i  float64

// var numslc[200] int

fmt.Println("Enter a the amount of numbers you have : \n")
fmt.Scanln(&totn)

for i = 0; i <= totn-1; i++ {
	fmt.Println("Enter the input number: ")
	fmt.Scanln(&innum)
	// numslc[i]=innum
	// sum += numslc[i]
	sum += innum

}
avg := sum/totn
fmt.Printf("The Average is : %v\n", avg)

}

