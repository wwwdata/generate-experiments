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
	input := []byte(`{"ID": "1234","Size":"XL","Color":"Red"}`)
	err := json.Unmarshal(input, &s)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n\n", s)
	fmt.Printf("as string %s\n\n", s)

	// marshaling works as well
	// awesome, we now have static compile time checked types!
	n :=
		Tshirt{
			ID:    "1337",
			Size:  XXL,
			Color: Blue,
		}

	json, err := json.Marshal(n)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(json))

}
