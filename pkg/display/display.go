package display

import "fmt"

func Display(msg string) { // capital letter => exported
	fmt.Println(msg)
}

func hello() { // not exported
	fmt.Println("hello")
}
