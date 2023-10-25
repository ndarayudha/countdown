package font

import "strings"

// Define ASCII art representations of various characters.
var colon = `
..
#.
..
#.
..
`

var zero = `
######.
#....#.
#....#.
#....#.
######.
`

var one = `
.....#.
.....#.
.....#.
.....#.
.....#.
`

var two = `
######.
.....#.
######.
#......
######.
`

var three = `
######.
.....#.
...###.
.....#.
######.
`

var four = `
#......
#......
#...#..
######.
....#..
`

var five = `
######.
#......
######.
.....#.
######.
`

var six = `
######.
#......
######.
#....#.
######.
`

var seven = `
######.
.....#.
.....#.
.....#.
.....#.
`

var height = `
######.
#....#.
######.
#....#.
######.
`

var nine = `
######.
#....#.
######.
.....#.
######.
`

// smallFonts defines the font used to display characters on the terminal.
var smallFonts = map[rune][][]rune{
	':': asArray(colon),
	'1': asArray(one),
	'2': asArray(two),
	'3': asArray(three),
	'4': asArray(four),
	'5': asArray(five),
	'6': asArray(six),
	'7': asArray(seven),
	'8': asArray(height),
	'9': asArray(nine),
	'0': asArray(zero),
}

// asArray converts a string representation of characters to a 2D slice of runes.
func asArray(chars string) [][]rune {
	result := [][]rune{}
	line := []rune{}
	str := strings.TrimPrefix(chars, "\n")

	for _, c := range str {
		if c == '\n' {
			result = append(result, line)
			line = []rune{}
		} else {
			line = append(line, c)
		}
	}
	return result
}
