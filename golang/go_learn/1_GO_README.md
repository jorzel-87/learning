PACKAGES

Every Go program is made up of packages.

Programs start running in package main.

This program is using the packages with import paths "fmt" and "math/rand".

By convention, the package name is the same as the last element of the import path. For instance, the "math/rand" package comprises files that begin with the statement package rand.

###############################################

IMPORTS

import (
	"fmt"
	"math"
)

=

import "fmt"
import "math"

###############################################

EXPORTED NAMES

In Go, a name is exported if it begins with a capital letter. For example, Pizza is an exported name, as is Pi, which is exported from the math package.

pizza and pi do not start with a capital letter, so they are not exported.

When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.

###############################################

FUNCTION

A function can take zero or more arguments.

In this example, add takes two parameters of type int.

Notice that the type comes after the variable name.

func add(x int, y int) int {
	return x + y
}

https://blog.golang.org/declaration-syntax

When two or more consecutive named function parameters share a type, you can omit the type from all but the last.

func add(x, y int) int {
	return x + y
}

###############################################

MULTIPLE RESULTS

A function can return any number of results.

The swap function returns two strings.

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

###############################################

READING A TYPE OF VAR

Type of var:

func main() {
	z := 3
	fmt.Printf("z = %T\n", z) 
}

###############################################

VARIABLES WITH INITIALIZERS

A var declaration can include initializers, one per variable.

If an initializer is present, the type can be omitted; the variable will take the type of the initializer.

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(c, python, java)
	fmt.Printf("c, python, java = %T, %T, %T", c, python, java)
}

###############################################

NAMED RETURN VALUES

Go's return values may be named. If so, they are treated as variables defined at the top of the function.

These names should be used to document the meaning of the return values.

A return statement without arguments returns the named return values. This is known as a "naked" return.

Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions.

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

###############################################

VARIABLES

The var statement declares a list of variables; as in function argument lists, the type is last.

A var statement can be at package or function level. We see both in this example.

package main

import "fmt"

var c, python, java bool

func main() {
	var i int
//	var i int = 4
	fmt.Println(i, c, python, java)
}

###############################################

SHORT VARIABLES DECLARATIONS

Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.

Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

k := 3

VS

var k int = 3

###############################################

TYPES

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128

Example:

package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

###############################################

ZERO VALUES

Variables declared without an explicit initial value are given their zero value.

The zero value is:

0 for numeric types,
false for the boolean type, and
"" (the empty string) for strings.

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

###############################################

TYPE CONVERTIONS

var i int = 42
var f float64 = float64(i)
var u uint = uint(f)

or

i := 42
f := float64(i)
u := uint(f)

###############################################

TYPE INFERENCE

func main() {
	v := 42 // change me!
	fmt.Printf("v is of type %T\n", v)
}

###############################################

CONSTANTS

High precision values.

Constants are declared like variables, but with the const keyword.

Constants can be character, string, boolean, or numeric values.

Constants cannot be declared using the := syntax.

package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "??"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
	fmt.Printf("Truth is of type %T\n", Truth)
	fmt.Printf("Pi is of type %T\n", Pi)
}

https://gobyexample.com/constants
