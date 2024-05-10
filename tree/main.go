package main

import (
	"fmt"
	"tree/tree"
)

func main() {
	t1ValueRange := 10
	t2ValueRange := 10

	t1 := tree.New(t1ValueRange)
	t2 := tree.New(t2ValueRange)

	var comparationStatement string
	var isSame = tree.Same(t1, t2)

	if isSame {
		comparationStatement = "same"
	} else {
		comparationStatement = "not same"
	}

	fmt.Println("Tree 1 and Tree 2 is", comparationStatement)
}
