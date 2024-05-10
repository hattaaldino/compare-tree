package tree

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *Tree, ch chan int) {

	if t.Left != nil {
		Walk(t.Left, ch)
	}

	ch <- t.Value

	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

/// Compare Tree 1 (t1) and Tree 2 (t2) to see if they have the same list of value
func Same(t1, t2 *Tree) bool {
	const treeSize = 10

	// Used to control channel and reference value in Tree 1 (ti)
	var walkControl int = 0

	// Since we only compare its value, then we just need to store value of Tree 1 (t1)
	var t1Values [treeSize]int

	// Walk through Tree 1 (t1)
	chanT1 := make(chan int)

	go Walk(t1, chanT1)

	for v1 := range chanT1 {
		t1Values[walkControl] = v1

		walkControl++

		if walkControl == treeSize {
			close(chanT1)
		}
	}

	// Walk through Tree 2 (t2)
	chanT2 := make(chan int)

	go Walk(t2, chanT2)

	walkControl = 0
	for v2 := range chanT2 {

		if v2 != t1Values[walkControl] {
			// If any value in Tree 2 (t2) is different than Tree 1 (t1), it can be concluded that both tree are not same
			return false
		}

		walkControl++

		if walkControl == treeSize {
			close(chanT2)
		}
	}

	return true
}
