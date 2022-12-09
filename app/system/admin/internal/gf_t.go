package main

import (
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
)

type T struct {
	s []int
	a int
}

func main() {
	t := T{a: 1, s: []int{}}
	g.Dump(t.s == nil)

	fmt.Println(8) //0000 1000 -> 0111 0111
	// 0111 0111  等于 7

	fmt.Println(24 &^ 8)
	fmt.Println(0 << 3)
}
