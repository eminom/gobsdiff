// +build ignore
package main

import (
	"fmt"

	"./bsdiff"
)

func printArray(a []int) {
	for i := 0; i < len(a); i++ {
		fmt.Printf("%v ", a[i])
	}
	fmt.Println()
}

func main() {
	instr := "abcbcb111" //
	fmt.Printf("try <%v>\n", instr)
	o := bsdiff.SortOutString1(instr)
	printArray(o)
	for _, v := range o {
		fmt.Printf("%v\n", instr[v:])
	}
}
