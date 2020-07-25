package main

import (
	_ "image/jpeg"
	_ "image/png"
	"syscall/js"

	"github.com/subeshb1/wasm-go-image-to-ascii/convert"
)

func printMessage(this js.Value, inputs []js.Value) interface{} {
	message := inputs[0].String()

	document := js.Global().Get("document")
	p := document.Call("createElement", "p")
	p.Set("innerHTML", message)
	return document.Get("body").Call("appendChild", p)
}

func converter(this js.Value, inputs []js.Value) interface{} {
	// Create convert options
	convertOptions := convert.DefaultOptions
	convertOptions.FixedWidth = 100
	convertOptions.FixedHeight = 40

	// Create the image converter
	converter := convert.NewImageConverter()
	array := inputs[0]
	inBuf := make([]uint8, array.Get("byteLength").Int())
	js.CopyBytesToGo(inBuf, array)

	return converter.ImageFile2ASCIIString(inBuf, &convertOptions)
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
	js.Global().Set("convert", js.FuncOf(converter))
	<-c
}
