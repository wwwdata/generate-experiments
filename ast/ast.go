package main

// how to parse consts of a type?
func main() {
	// example go file that we want to parse all the Unicorn constants out of
	src := `
		package foo

		type Unicorn int

		const (
			Pink Unicorn = iota
			Fluffy
			Rainbow
		)
	`
}
