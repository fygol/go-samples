/* 

Pointer is a variable which stores the address of another variable.
& operator - yields the address of a variable.
* operator - retrieves the variable that the pointer refers to.
There is no pointer arithmetic.

*/

package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 10

	fmt.Println(a) // 10

	var p *int
	fmt.Println(p) // <nil>
	
	// & - address operator
	p = &a // returns *int (pointer to int); the same as var p *int = &a
	fmt.Println(p) // 0xc42000a2f0
	fmt.Println(reflect.TypeOf(p)) // *int

	*p = 11 // dereference pointer variable, the same as a = 11
	fmt.Println(a) // 11
}
