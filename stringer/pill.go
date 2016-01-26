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
	Acetaminophen = Paracetamol // Acetaminophen is the same as Paracetamol
)

// Implements the Stringer interface
func (p Pill) String() string {
	switch p {
	case Placebo:
		return "Placebo"
	case Aspirin:
		return "Aspirin"
	case Ibuprofen:
		return "Ibuprofen"
	case Paracetamol: // == Acetaminophen
		return "Paracetamol"
	}
	return fmt.Sprintf("Pill(%d)", p)
}

func main() {
	var p Pill = Ibuprofen
	fmt.Printf("Taking a %s pill\n", p)
}
