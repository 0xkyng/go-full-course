package main

import (
	"strings"
)

type list []*product

func (l list) String() string {
	if len(l) == 0 {
		return "Sorry. We're waiting for delivery 🚚.\n"
		
	}
    var str strings.Builder
	for _, p := range l {
		// p.print()
		// fmt.Printf("* %s\n", p)
		str.WriteString("* ")
		str.WriteString(p.String())
		str.WriteRune('\n')

	}
	return str.String()
}

func (l list) discount(ratio float64) {
	for _, p := range l {
		p.discount(ratio)
	}
}

// Len isthe number of elements in the collection
func (l list) Len() int {
	return len(l)
}

// Less reports whether the element with 
// index i should sort before the element with index j
func (l list) Less(i, j int) bool {
	return l[i].Title < l[j].Title
}

// Swap swaps the element with indexes i and j
func (l list) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

