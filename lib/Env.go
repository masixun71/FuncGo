package lib

import (
	"os"
	"strings"
)

var root string

func GetRoot() string {
	if root != "" {
		return root
	}

	roots, _ := os.Getwd()

	strPath := strings.Split(roots, "/")
	projectPath := ""
	lenPath := len(strPath)

	for i := 0; i < lenPath-1; i++ {
		if strPath[i] == "go" && strPath[i+1] == "src" {
			if i+2 <= lenPath-1 {
				projectPath = strings.Join(strPath[0:i+3], "/")
			} else {
				panic("your project not in go/src")
			}
		}
	}

	if projectPath == "" {
		panic("your project not in go/src")
	}
	root = projectPath

	return projectPath
}

type Path interface {
	GetPathByRoot() string
}

type Pather struct {
	path string
}

func NewPath(path string) Path {
	return &Pather{path: path}
}

func (p *Pather) GetPathByRoot() string {
	return GetRoot() + p.path
}
