package main

type LineCallback func(line string)

func lineIterator(lines []string, callback LineCallback) {
	for i := 0; i < len(lines); i++ {
		callback(lines[i])
	}
}

func main() {
	lines := []string{
		"There are",
		"68 letters,",
		"five digits,",
		"12 spaces,",
		"and 4 punctuation marks in these lines of text!",
	}

	letters := 0
	numbers := 0
	punctuation := 0
	spaces := 0

	lineFunc := func(line string) {
		for _, r := range line {
			if unicode.IsLetter(r) {
				letters++
			}
			if unicode.IsDigit(r) {
				numbers++
			}
			if unicode.IsPunct(r) {
				punctuation++
			}
			if unicode.IsSpace(r) {
				spaces++
			}
		}
	}

	lineIterator(lines, lineFunc)

	fmt.Println(letters, "letters")
	fmt.Println(numbers, "numbers")
	fmt.Println(punctuation, "punctuation marks")
	fmt.Println(spaces, "spaces")
}
