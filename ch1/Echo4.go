// Echo4 - Выводит имя выполяемой команды, аргументы командной строки
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	os.Args[0] = "echo"
	fmt.Println(strings.Join(os.Args, " "))
}
