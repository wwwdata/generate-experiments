package main

import "fmt"

type Pill int

// defining constants of the new type Pill
// iota counts from 0 and increases in steps by 1 (in case you didn't know this feature)
//go:generate stringer -type=Pill
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
	n := NewPill
	fmt.Printf("Taking a %s pill and a %s pill\n", p, n)
}
