package main

import "fmt"

type Part string

func lineContents(parts []Part) {
	fmt.Println(" === Assembly line ===")

	for i, element := range parts {
		fmt.Println(i, element)
	}

	fmt.Println(" =====================")
	fmt.Println()
}

func main() {
	assemblyLine := make([]Part, 3)

	assemblyLine = []Part{"Part one", "Part two", "Part three"}
	lineContents(assemblyLine)

	assemblyLine = append(assemblyLine, "Part four", "Part five")
	lineContents(assemblyLine)

	assemblyLine = assemblyLine[3:]
	lineContents(assemblyLine)

}
