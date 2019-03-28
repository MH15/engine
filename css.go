package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

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
	value   string
	id      string
	class   []string
	tagName string
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
	keyword   string
	number    float32
	color     Color
}

// Color holds an RGB color value
type Color struct {
	r uint8
	b uint8
	g uint8
	a uint8
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

// ParseCSS a whole CSS stylesheet
func ParseCSS(r io.Reader) Stylesheet {
	buf := bufio.NewReader(r)
	return Stylesheet{rules: parseRules(buf)}
}

// Parse a list of rule sets, separated by optional whitespace.
func parseRules(buf *bufio.Reader) []Rule {
	rules := []Rule{}

	for {
		char, size, err := buf.ReadRune()
		buf.UnreadRune()
		if err != nil {
			Use(string(char))
			Use(string(size))
			if err == io.EOF {
				fmt.Println("EOF reached.")
				break
			} else {
				log.Fatal(err)
			}
		} else {
			buf.UnreadRune()
			// TODO: consumeWhitespace(buf)
			consumeWhitespace(buf)

			fmt.Println("~~ Parsing Rule ~~")
			rules = append(rules, parseRule(buf))
		}

	}
	return rules
}

// Parse a rule set: `<selectors> { <declarations> }`.
func parseRule(buf *bufio.Reader) Rule {
	// fmt.Println("char before error: " + string(nextChar(buf)))

	return Rule{selectors: parseSelectors(buf), declarations: parseDeclarations(buf)}
}

// Parse a comma-separated list of selectors.
func parseSelectors(buf *bufio.Reader) []Selector {
	selectors := []Selector{}
	for {

		fmt.Println("~~~~ Parsing Selector ~~")
		selectors = append(selectors, parseSimpleSelector(buf))
		consumeWhitespace(buf)

		if nextChar(buf) == ',' {
			buf.ReadRune()
			consumeWhitespace(buf)
		}
		if nextChar(buf) == '{' {
			break
		}
	}
	return selectors
	// TODO: specificity
}

// Parse one simple selector, e.g.: `type#id.class1.class2.class3`
func parseSimpleSelector(buf *bufio.Reader) Selector {
	selector := Selector{}
	for {
		char := nextChar(buf)
		if char == '#' {
			fmt.Println("~~~~~~ Parsed # ~~")
			buf.ReadRune()
			selector.id = parseIdentifier(buf)
		} else if unicode.IsLetter(char) || unicode.IsNumber(char) {

			selector.tagName = parseIdentifier(buf)
			fmt.Println("tagName: '" + selector.tagName + "'")

		} else if char == '.' {
			fmt.Println("~~~~~~ Parsed . ~~")
			buf.ReadRune()
			selector.class = append(selector.class, parseIdentifier(buf))
		} else if char == '*' {
			fmt.Println("~~~~~~ Parsed * ~~")
			buf.ReadRune()
		} else if validIdentifierChar(char) {
			// fmt.Println("~~~~~~ Parsed tag name ~~")
			// selector.tagName = parseIdentifier(buf)
			// fmt.Println("~~~~~~ '" + selector.tagName + "' ~~")
		} else {
			break
		}

		break

		// if char, size, err := buf.ReadRune(); err != nil {
		// 	Use(string(size))
		// 	// Use(int(char))
		// 	if err == io.EOF {
		// 		fmt.Println("EOF reached.")
		// 		break
		// 	} else {
		// 		log.Fatal(err)
		// 	}
		// } else {
		// 	buf.UnreadRune()
		// 	switch char {
		// 	case '#':
		// 		buf.ReadRune()
		// 		selector.id = parseIdentifier(buf)
		// 		fmt.Println("FOUND id: " + selector.id)
		// 	case '.':
		// 		buf.ReadRune()
		// 		selector.class = append(selector.class, parseIdentifier(buf))
		// 		fmt.Println("FOUND class: " + string(strings.Join(selector.class, ", ")))
		// 	case '*':
		// 		buf.ReadRune()
		// 		fmt.Println("universal selector")
		// 	default:
		// 		selector.tagName = parseIdentifier(buf)
		// 		// fmt.Println("FOUND tagName: " + selector.tagName)
		// 	}
		// }
	}
	return Selector{}
}

// Parse a list of declarations enclosed in `{ ... }`.
func parseDeclarations(buf *bufio.Reader) []Declaration {
	declarations := []Declaration{}
	for {
		consumeWhitespace(buf)

		if nextChar(buf) == '}' {
			buf.ReadRune()
			break
		}
		declarations = append(declarations, parseDeclaration(buf))
	}
	return declarations
}

// Parse one `<property>: <value>;` declaration.
func parseDeclaration(buf *bufio.Reader) Declaration {
	declaration := Declaration{}

	fmt.Println("found: '" + string(nextChar(buf)) + "'")
	buf.ReadRune()
	fmt.Println("found: '" + string(nextChar(buf)) + "'")
	panic("ok")

	fmt.Println("loop test b")
	propertyName := parseIdentifier(buf)
	value := ""
	consumeWhitespace(buf)
	if nextChar(buf) == ':' {
		consumeWhitespace(buf)
		value = parseIdentifier(buf) // TODO: make parseValue
		consumeWhitespace(buf)
		if nextChar(buf) == ';' {
			declaration.name = propertyName
			declaration.value = Value{keyword: value}
		}
	}

	return declaration
}

// ConsumeCondition is a wrapper type for a function used as Comparators in Java
// are used to fulfill a condition requirement in the function consumeWhile()
type ConsumeCondition func(rune) bool

// Parse a property name or keyword
func parseIdentifier(buf *bufio.Reader) string {
	// fmt.Println("loop test a")
	// return consumeWhile(buf, validIdentifierChar)
	result := ""

	for {
		if char, size, err := buf.ReadRune(); err != nil {
			Use(string(size))
			if err == io.EOF {
				fmt.Println("EOFFFFF")
				break
			} else {
				log.Fatal(err)
			}
		} else {
			// || char == '_' || char == '-'
			if unicode.IsLetter(char) || unicode.IsNumber(char) {
				// log.Println("YEET")
				result += string(char)
			} else {
				// buf.UnreadRune()
				// fmt.Println("what")
				break
			}
		}
	}

	buf.UnreadRune()

	// fmt.Println("string: " + result)
	return result
}

// Consume and discard zero or more whitespace characters.
func consumeWhitespace(buf *bufio.Reader) {
	// consumeWhile(buf, isWhitespace)
	// fmt.Println("loop test a")
	// return consumeWhile(buf, validIdentifierChar)

	for {
		if unicode.IsSpace(nextChar(buf)) || unicode.IsControl(nextChar(buf)) {
			buf.ReadByte()
		} else {
			break
		}

	}

	// result := ""
	// for {
	// 	if char, size, err := buf.ReadRune(); err != nil {
	// 		Use(string(size))
	// 		if err == io.EOF {
	// 			fmt.Println("EOFFFFF")
	// 			break
	// 		} else {
	// 			log.Fatal(err)
	// 		}
	// 	} else {
	// 		// fmt.Println("char: " + string(char))
	// 		// buf.UnreadRune()
	// 		if isWhitespace(char) {
	// 			result += string(char)
	// 		} else {
	// 			// buf.UnreadRune()
	// 			// fmt.Println("what")
	// 			break
	// 		}
	// 	}
	// }

	// fmt.Println("whitespace: '" + result + "'")
	// return result
}

// Consume characters until `test` returns false.
func consumeWhile(buf *bufio.Reader, condition ConsumeCondition) string {
	result := ""
	for {
		if char, size, err := buf.ReadRune(); err != nil {
			Use(string(size))
			if err == io.EOF {
				fmt.Println("EOFFFFF")
				break
			} else {
				log.Fatal(err)
			}
		} else {
			// fmt.Println("char: " + string(char))
			// buf.UnreadRune()
			if !condition(char) {
				buf.UnreadRune()
				fmt.Println("CONDITION NOT MET")
				break
			} else {
				result += string(char)
			}
		}
	}
	return result
}

func validIdentifierChar(r rune) bool {
	return !(unicode.IsLetter(r) || unicode.IsNumber(r) || r == '_' || r == '-')
}

// isWhitespace is a function tells you if a rune is a whitespace character or not; it accounts for all 26 UNICODE whitespace characters.
// @author: https://github.com/reiver/go-whitespace
func isWhitespace(r rune) bool {

	result := false

	switch r {
	case
		'\u0009', // horizontal tab
		'\u000A', // line feed
		'\u000B', // vertical tab
		'\u000C', // form feed
		'\u000D', // carriage return
		'\u0020', // space
		'\u0085', // next line
		'\u00A0', // no-break space
		'\u1680', // ogham space mark
		'\u180E', // mongolian vowel separator
		'\u2000', // en quad
		'\u2001', // em quad
		'\u2002', // en space
		'\u2003', // em space
		'\u2004', // three-per-em space
		'\u2005', // four-per-em space
		'\u2006', // six-per-em space
		'\u2007', // figure space
		'\u2008', // punctuation space
		'\u2009', // thin space
		'\u200A', // hair space
		'\u2028', // line separator
		'\u2029', // paragraph separator
		'\u202F', // narrow no-break space
		'\u205F', // medium mathematical space
		'\u3000': // ideographic space
		result = true
	default:
		result = false
	}

	return result
}

func nextChar(buf *bufio.Reader) rune {
	char, size, err := buf.ReadRune()
	// fmt.Println("reading char: " + string(char))
	check(err)
	Use(string(size))
	buf.UnreadRune()

	return char
}

func nextCharAdv(buf *bufio.Reader) (rune, error) {
	char, size, err := buf.ReadRune()
	// fmt.Println("r char: " + string(char))
	// check(err)
	Use(string(size))
	buf.UnreadRune()

	return char, err
}

func Use(s string) {
}
