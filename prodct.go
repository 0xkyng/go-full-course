package main

import (
	"fmt"
	"strconv"
	"time"
)

type product struct {
	title string
	price money
	released interface{}
}

func (p *product) print() {
	fmt.Printf("%s: %s (%s)\n", p.title, p.price.string(), format(p.released))
}

func (p *product) discount(ratio float64) {
	p.price *= money(1 - ratio)
}

func format(v interface{}) string {
	var t int

	switch v := v.(type) {
	case int:
		// book{title: "moby dick", price: 10, published: 118281600},
		t = v
	case string:
		// book{title: "odyssey", price: 15, published: "733622400"},
		t, _ = strconv.Atoi(v)
	default:
		// book{title: "hobbit", price: 25},
		return "unknown"
	}

	// Mon Jan 2 15:04:05 -0700 MST 2006
	const layout = "2006/01"

	u := time.Unix(int64(t), 0)
	return u.Format(layout)
}