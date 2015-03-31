package main

import (
	"fmt"
)

func main() {

	a := []int{1, 1, 2, 3}
	b := []int{3, 4, 5, 3, 5}

	aOnly, both, bOnly := diffs(a, b)

	fmt.Println(aOnly, both, bOnly)
}

func diffs(a, b []int) (aOnly, both, bOnly []int) {

	aMap := make(map[int]int)

	for _, aElem := range a {
		aMap[aElem] = aMap[aElem] + 1
	}

	for _, bElem := range b {
		aCount, ok := aMap[bElem]
		if !ok {
			bOnly = append(bOnly, bElem)
		} else {
			both = append(both, bElem)
			count := aCount - 1
			if count > 0 {
				aMap[bElem] = count
			} else {
				delete(aMap, bElem)
			}
		}
	}

	for aElem, count := range aMap {
		for count > 0 {
			aOnly = append(aOnly, aElem)
			count = count - 1
		}
	}
	return
}
