package main

import (
	"go-playground/pkg/display"
	"go-playground/pkg/msg"
)

func main() {
	msg.Hi()

	display.Display("Hello from display")

	msg.Exciting("Exciting message")
}
