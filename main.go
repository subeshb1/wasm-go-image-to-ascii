package main

import (
	"encoding/json"
	_ "image/jpeg"
	_ "image/png"
	"syscall/js"

	"github.com/subeshb1/wasm-go-image-to-ascii/convert"
)

func converter(this js.Value, inputs []js.Value) interface{} {
	array := inputs[0]
	inBuf := make([]uint8, array.Get("byteLength").Int())
	js.CopyBytesToGo(inBuf, array)
	convertOptions := convert.Options{}
	err := json.Unmarshal([]byte(inputs[1]), &convertOptions)
	if err != nil {
		convertOptions = convert.DefaultOptions
	}

	converter := convert.NewImageConverter()
	return converter.ImageFile2ASCIIString(inBuf, &convertOptions)
}

func main() {
	c := make(chan bool)
	js.Global().Set("convert", js.FuncOf(converter))
	<-c
}
