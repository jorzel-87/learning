package main

import (
	"fmt"
)

func example(dy, dx int) []uint8 {

    p := make([]uint8, dy)

    for i := 0; i < dy; i++ {
	    p[i] = uint8(dx)
    }
return p    
}

func main() {
	fmt.Println(example(6, 5))
}

