package main

import "fmt"

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)

}

func main() {

	number := 250

	Printfln("Binary: %b", number)
	Printfln("Decimal: %d", number)
	Printfln("Octal: %o, %O", number, number)
	Printfln("Hexadecimal: %x, %X", number, number)

}
