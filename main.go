package main

func main() {
	cakes := 20
	apples := 25

	boxCount, cakeCount, appleCount := GetItemsPerBox(cakes, apples)
	println("cake:", cakes, "apple:", apples)
	println("Total box:", boxCount, "cake per box:", cakeCount, "apple per box:", appleCount)
}

// CountBox returns the number of boxes needed to pack the given number of cakes and apples.
func CountBox(cakes int, apples int) int {
	smaller := cakes
	greater := apples

	// if apples smaller than cakes
	// swap the value
	if apples < smaller {
		tmp := smaller
		smaller = apples
		greater = tmp
	}

	for i := smaller; i > 1; i-- {
		if smaller%i == 0 && greater%i == 0 {
			return i
		}
	}

	return 1
}

// GetItemsPerBox returns the number of boxes and cakes & apples per box
func GetItemsPerBox(cakes, apples int) (box, cakeCount, appleCount int) {
	box = CountBox(cakes, apples)
	cakeCount = cakes / box
	appleCount = apples / box

	return
}
