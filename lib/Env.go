package lib

import (
	"os"
	"strings"
)

var root string

func GetRoot() string  {
	if root != "" {
		return root
	}

	roots, _ := os.Getwd()
	index := strings.LastIndex(roots, "FuncGo")
	root = string([]rune(roots)[:index + 6])

	return root
}
