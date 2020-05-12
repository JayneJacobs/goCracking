package readfull

import (
	"fmt"
	"io"
	"strings"
)

type itemstrct struct {
	in int
	str string
}

// MyReadFull ranges throug a  aslice of structs to create a New reader and byte array for teh io readfull method. 
func MyReadFull()  {
	myrfarray := []itemstrct{
		{3, "Jayne",},
		{2, "Jim",},
	}

	for _, str1 := range myrfarray {
		strrd := strings.NewReader(str1.str)
		bytstr := make([]byte, str1.in)
		bytesR, err := io.ReadFull(strrd, bytstr)
		fmt.Printf("Len(read) %d, numbytesread = %d\n err = %v, str1 = %s\n", str1.in, bytesR,  err,  bytstr)
		
	}
}