package main

import (
	"fmt"
	"math"
)

var isqrtCache = make(map[int16]int16)

func isqrt(n int16) int16 {
        x := isqrtCache[n]
        if x == 0 {
		fmt.Println("call Sqrt")
                x = int16(math.Sqrt(float64(n)))
                isqrtCache[n] = x
        }

	fmt.Printf("isqrt %v %v\n", n, x)
        return x
}

func main() {
	for i := int16(1); i < 10; i++ {
		fmt.Printf("%v\n", isqrt(i))
	}
	for i := int16(1); i < 10; i++ {
		fmt.Printf("%v\n", isqrt(i))
	}
}
