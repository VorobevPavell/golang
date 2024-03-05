// Echo5 - выводит индекс и значение каждого аргумента по одному аргументу в строке.
package main

import (
	"fmt"
	"os"
)

func main() {
	os.Args[0] = "echo"
	s, sep := "", ""
	fmt.Printf("command: %s\n", os.Args[0])
	for index, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
		fmt.Printf("index: %d, value: %s\n", index, arg)
	}
}
