package hashtable

import "fmt"

// MakeHashTable 
func MakeHashTable()  {
	cx := make(map[int]string)

	cx[1] = "Jayne"
	cx[2] = "Jim"
	cx[3] = "Rich"
	for ix, jx  := range cx {
		fmt.Printf("key: %d value: %s \n", ix, jx)
	}
}