package main

import "fmt"

type Pill int

// defining constants of the new type Pill
// iota counts from 0 and increases in steps by 1 (in case you didn't know this feature)
const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
	NewPill
	Acetaminophen = Paracetamol // Acetaminophen is the same as Paracetamol
)

func main() {
	var p Pill = Ibuprofen
	fmt.Printf("Taking a %s pill\n", p)
}
