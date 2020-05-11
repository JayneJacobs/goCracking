package converttype

import (
	"fmt"
	"reflect"
	"strconv"
)

// ConvertStringInt uses the reflect package
func ConvertStrings()  {

	a := "21.123438782734"
	b, _ :=  strconv.ParseBool(a)
	fmt.Println(reflect.TypeOf(b))

	flt, _ := strconv.ParseFloat(a, 64)

	stint, _ := strconv.ParseInt(a, 10, 64)

	fmt.Printf("a string: %T, bool: %T as float64: %T. int: %T\n", a, b, flt, stint)

}

// ConvertInts uses strconf.Itoa .FormatUint and FormatInt
func ConvertInts()  {
	myint := 10
	myint1 := -10000
	int1 := strconv.Itoa(myint)
	fmt.Printf("myint: %T %v converted to: %T  %v\n ", myint, myint, int1, int1)

	int2 := strconv.FormatUint(uint64(myint1), 10)
	fmt.Printf("my int1: %T %v converted to: %T  %v\n ", int2, int2, int1, int1)

	int3 := strconv.FormatInt(int64(myint1), 10)
	fmt.Printf("my int: %T %v converted to: %T  %v\n ", myint, myint, int3, int3)

}