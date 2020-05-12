package readatleast

import (
	"fmt"
	"io"
	"strings"
)


type Mystruct struct {
	int1 int
	int2 int
	str1 string
}

// MyNewReader ranges through a struct for values to create a Reader, a byte array and ReadAtlLeast buffer
func MyNewReader()  {
fmt.Println("This")
mysrcslc := []Mystruct {
	   {
			1,
			2,
			"This",
		},
		{
			3,
			4,
			"That",
		},
	}

	for _, myarr := range mysrcslc {
		ip := strings.NewReader(myarr.str1)
		bf := make([]byte, myarr.int1)
		br, err := io.ReadAtLeast(ip, bf, myarr.int1)
		fmt.Printf("Lent(bf) = %d, Min = %d, \nnumofbytesread = %d, err = %s, %s, bf=%s\n", myarr.int1, myarr.int2, br, err, myarr.str1, bf )
	}
}
