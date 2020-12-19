
Exercise: Slices

Implement Pic. It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers. When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.

The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.

(You need to use a loop to allocate each []uint8 inside the [][]uint8.)

(Use uint8(intValue) to convert between types.)


package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {

    //making matrix

    p := make([][]uint8)

    for i := 0; i < dy; i++ {
        p[i] := make([]uint8, dx)
    }

    //filling up the matrix line by line

    for y := 0; y < dy; y++ {
        for x := 0; x < dx; x++ {
            p[x][y] = uint8(x+y)/2
        }
    }

return p
}


func main() {
pic.Show(Pic)
}
