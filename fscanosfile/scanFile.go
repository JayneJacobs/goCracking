package fscanosfile

import (
	"fmt"
	"os"
)

var (
	myint int
	mystring string
	myfloat float32
	mybool bool
	myuint uint64
)

// ScanFile retrieves values from a file and prints them
func ScanFile()  {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	fmt.Fscan(file, &myint,&myfloat,&myint,&mystring,&myuint)
	fmt.Printf("types struct %d,%f,%d,%s,%d, \n", myint,myfloat,myint,mystring,myuint)
	fmt.Fscanln(file, &myint,&myfloat,&myint,&mystring,&myuint)
	fmt.Printf("types struct %d,%f,%d,%s,%d, \n", myint,myfloat,myint,mystring,myuint)
	fmt.Fscanf(file, "%d\n%f\n%d\n%s %d, \n",&myint,&myfloat,&myint,&mystring,&myuint)
	fmt.Printf("types struct %d,%f,%d,%s,%d, \n", myint,myfloat,myint,mystring,myuint)


}