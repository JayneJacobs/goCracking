package findtype

import (
	"fmt"
	"reflect"
)
// TypeFinder will print the type
func TypeFinder()  {
	t1 := "Jayne"
	t2 := 1
	t3 := []string{"This","that","other"}
	t4 := []int{1,2,3}
	var t5 uint
	t6 :=1.4529834750834
	t7 := map[string]int{"x":1, "y":2}
	fmt.Printf("This is t1 \t%v, %v kind %v: \n", t1, reflect.TypeOf(t1),reflect.ValueOf(t1).Kind())
	fmt.Printf("This is t2 \t%v, %v kind %v: \n", t2, reflect.TypeOf(t2),reflect.ValueOf(t2).Kind())
	fmt.Printf("This is t3 \t%v, %v kind %v: \n", t3, reflect.TypeOf(t3),reflect.ValueOf(t3).Kind())
	fmt.Printf("This is t4 \t%v, %v kind %v: \n", t4, reflect.TypeOf(t4),reflect.ValueOf(t4).Kind())
	fmt.Printf("This is t5 \t%v, %v kind %v: \n", t5, reflect.TypeOf(t5),reflect.ValueOf(t5).Kind())
	fmt.Printf("This is t6 \t%v, %v kind %v: \n", t6, reflect.TypeOf(t6),reflect.ValueOf(t6).Kind())
	fmt.Printf("This is t7 \t%v, %v kind %v: \n", t7, reflect.TypeOf(t7),reflect.ValueOf(t7).Kind())
fmt.Println("######Using %d ###############")
	fmt.Printf("This is t1 %v kind %T: \n", t1, t1,)
	fmt.Printf("This is t2 %v kind %T: \n", t2, t2,)
	fmt.Printf("This is t3 %v kind %T: \n", t3, t3,)
	fmt.Printf("This is t4 %v kind %T: \n", t4, t4,)
	fmt.Printf("This is t5 %v kind %T: \n", t5, t5,)
	fmt.Printf("This is t6 %v kind %T: \n", t6, t6,)
	fmt.Printf("This is t7 %v kind %T: \n", t7, t7,)

}