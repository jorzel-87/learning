
FUNCTION VALUES

 Functions are values too. They can be passed around just like other values.

Function values may be used as function arguments and return values. 

package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

Another example:

package main

import "fmt"

func add(a, b int) int64 {
	return int64(a + b)
}

func main() {
	result := add(1, 5)
	fmt.Println(result)
}

###############################################

CLOSURES

This function intSeq returns another function, which we define anonymously in the body of intSeq. The returned function closes over the variable i to form a closure.

package main

import "fmt"


func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}


func main() {

    nextInt := intSeq()

    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    newInts := intSeq()
    fmt.Println(newInts())
}

###############################################
