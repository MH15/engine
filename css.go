package main

// Stylesheet is everything
type Stylesheet struct {
	rules []Rule
}

// Rule contains selectors and declarations
type Rule struct {
	selectors    []Selector
	declarations []Declaration
}

// Selector stores CSS Selectors
type Selector struct {
	selectorType SimpleSelector
	value        string
}

// SimpleSelector defined as single selector, no chains
type SimpleSelector uint8

const (
	// tagName : example "div"
	tagName = iota
	id
	class
)

// Declaration is one style line in CSS
type Declaration struct {
	name  string
	value Value
}

// Value holds the value of the style line
type Value struct {
	valueType Values
}

// Values const enum is better than just using strings
type Values uint8

const (
	// Keyword is for text
	Keyword Values = iota
	// Number is for numbers
	Number
	// ColorValue is for colors
	ColorValue
)
