package main

import (
	"encoding/json"
	"fmt"
)

// Tshirt that can be bought
type Tshirt struct {
	ID    string
	Size  string // needs to validate against S, L, M, XL, XXL
	Color string // needs to validate against Red, Green, Blue
}

func main() {
	var s Tshirt
	input := []byte(`{"ID": "1234","Size":"Blubb","Color":"Rainbow"}`)
	err := json.Unmarshal(input, &s)
	if err != nil {
		panic(err)
	}

	// manually calling some sort of validator that needs to be maintained extra
	err = validate(s)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v", s)
}

func validate(t Tshirt) error {
	if t.Size != "S" && t.Size != "L" && t.Size != "M" && t.Size != "XL" && t.Size != "XXL" {
		return fmt.Errorf("Invalid T-Shirt size %s", t.Size)
	}

	if t.Color != "Red" && t.Color != "Green" && t.Color != "Blue" {
		return fmt.Errorf("Invalid T-shirt color %s", t.Color)
	}

	return nil
}
