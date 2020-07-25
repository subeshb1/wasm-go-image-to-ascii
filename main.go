package main

import (
	_ "image/jpeg"
	_ "image/png"
	"syscall/js"
)

func printMessage(this js.Value, inputs []js.Value) interface{} {
	message := inputs[0].String()

	document := js.Global().Get("document")
	p := document.Call("createElement", "p")
	p.Set("innerHTML", message)
	return document.Get("body").Call("appendChild", p)
}

func main() {

	c := make(chan bool)
	// Create convert options
	// convertOptions := convert.DefaultOptions
	// convertOptions.FixedWidth = 100
	// convertOptions.FixedHeight = 40

	// // Create the image converter
	// converter := convert.NewImageConverter()
	js.Global().Set("printMessage", js.FuncOf(printMessage))
	<-c
}
