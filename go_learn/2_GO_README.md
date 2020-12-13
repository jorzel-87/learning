FOR

Go has only one looping construct, the for loop.

The basic for loop has three components separated by semicolons:

    the init statement: executed before the first iteration
    the condition expression: evaluated before every iteration
    the post statement: executed at the end of every iteration

The init statement will often be a short variable declaration, and the variables declared there are visible only in the scope of the for statement.

The loop will stop iterating once the boolean condition evaluates to false.

Note: Unlike other languages like C, Java, or JavaScript there are no parentheses surrounding the three components of the for statement and the braces { } are always required.

package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}


Init and post are optional:


func main() {
	i := 1
	for ; i < 10; {
		i += i
		fmt.Println(i)
	}
}

###############################################

WHILE (really it's just for)

package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
		fmt.Println(sum)
	}
	fmt.Println("Sum is:", sum)
}


(If you omit the loop condition it loops forever)

func main() {
	for {
	}
}

###############################################

IF, ELSE

package main
 
import (
	"fmt"
)
 
func main() {
	x := 100
 
	if x == 100 {
		fmt.Println("It equals:)")
	} else {
		fmt.Println("Not equal:/")
	}
}

###############################################

IF WITH SHORT STATEMENT

Like for, the if statement can start with a short statement to execute before the condition.

Variables declared by the statement are only in scope until the end of the if. 

package main

import "fmt"
import "math"

func add(x int, y int) int {
	return x + y
}

var k float64 = 2

func main() {
	if v := math.Pow(k, k); v < 5 {
	fmt.Println(add(10, 15))
	}
//fmt.Println(v)
}

###############################################

IF, ELSE, SHORT STATMENTS AND VARS

Variables declared inside an if short statement are also available inside any of the else blocks. 


package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

###############################################

SWITCH

shorter way to write a sequence of if - else 

Go's switch is like the one in C, C++, Java, JavaScript, and PHP, except that Go only runs the selected case, not all the cases that follow. In effect, the break statement that is needed at the end of each case in those languages is provided automatically in Go. Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers. 

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

another example:

func main() {
 	

    i := "Przypadek3"
    fmt.Print("Write ", i, " as ")
    switch i {
    case "Przypadek1":
        fmt.Println("one")
    case "Przypadek2":
        fmt.Println("")
    case "Przypadek3":
        fmt.Println("three")}
}


Switch cases evaluate cases from top to bottom, stopping when a case succeeds.

(For example,

switch i {
case 0:
case f():
}

does not call f if i==0.) 


Switch without a condition is the same as switch true.

This construct can be a clean way to write long if-then-else chains. 

package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

###############################################

DEFER

A defer statement defers the execution of a function until the surrounding function returns.

The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns. 

package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

Stacking DEFERS

Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.

To learn more about defer statements read this blog post. 

package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}