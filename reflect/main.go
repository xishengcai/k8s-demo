package main

import (
	"fmt"
	"reflect"

)

func main() {
	i := 20
	_ = reflect.TypeOf(i)

	a := make(map[string]int)

	delete (a, "a")
	fmt.Println(a)
}
