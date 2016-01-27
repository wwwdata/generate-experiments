package main

import (
	"encoding/json"
	"fmt"
)

// Tshirt that can be bought
type Tshirt struct {
	ID    string
	Size  ShirtSize
	Color Color
}

//go:generate jsonenums -type=ShirtSize
type ShirtSize int

const (
	S ShirtSize = iota
	M
	L
	XL
	XXL
)

//go:generate jsonenums -type=Color
type Color int

const (
	Red Color = iota
	Green
	Blue
)

func main() {
	var s Tshirt
	input := []byte(`{"ID": "1234","Size":"Blubb","Color":"Rainbow"}`)
	err := json.Unmarshal(input, &s)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v", s)
}
