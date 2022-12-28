package main

import (
	"io"
	"strings"
)

func processData(reader io.Reader, writer io.Writer) {

	count, err := io.Copy(writer, reader)

	if err == nil {
		Printfln("Read %v bytes", count)

	} else {
		Printfln("Error: %v", err.Error())

	}

}

func main() {

	r := strings.NewReader("Kayak")
	var builder strings.Builder
	processData(r, &builder)
	Printfln("String builder contents: %s", builder.String())

}
