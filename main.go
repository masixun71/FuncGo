package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	root, _ := os.Getwd()

	fmt.Println(root)
	index := strings.LastIndex(root, "FuncGo")


	s := string([]rune(root)[:index + 6])
	fmt.Println(s) //得到 "a我c"
}
