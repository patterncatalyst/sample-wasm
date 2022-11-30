package main
import (
	"fmt"
	"strconv"
	"syscall/js"
	)
func main() {
	fmt.Println("hello, webassembly!")
	document := js.Global().Get("document")
	p := document.Call("createElement", "p")
	p.Set("innerHTML", "Hello WASM from Go!")
	document.Get("body").Call("appendChild", p)
	js.Global().Set("add", js.FuncOf(add))
	select {} // Code must not finish
}
func add(this js.Value, arg []js.Value) interface{} {
	value1 := js.Global().Get("document").Call("getElementById", arg[0].String()).Get("value").String()
	value2 := js.Global().Get("document").Call("getElementById", arg[1].String()).Get("value").String()
	numTwo, _ := strconv.Atoi(value2)
	numOne, _ := strconv.Atoi(value1)
	document := js.Global().Get("document")
	div := document.Call("createElement", "div")
	div.Set("innerHTML", numOne+numTwo)
	document.Get("body").Call("appendChild", div)
	println(numOne + numTwo)
	return nil
}

