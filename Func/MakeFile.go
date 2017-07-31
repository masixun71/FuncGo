package Func

import (
	"FuncGo/lib"
	"fmt"
	"io/ioutil"
	"regexp"
	"text/template"
	"os"
	"io"
	"go/token"
	"go/parser"
	"go/ast"
	"reflect"
)

type MakeFile interface {
	MakeCode()
}

type MakeFiler struct {
	ReadPath lib.Path
	FuncName string
	Template *template.Template
}

func NewMakeFiler(readPath, funcName string) MakeFile {
	return &MakeFiler{ReadPath: lib.NewPath(readPath), FuncName: funcName}
}

func (m *MakeFiler) MakeCode() {
	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, m.ReadPath.MakePath(), nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	var start token.Pos
	var end token.Pos

	for _, decl := range f.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if ok && fn.Name.Name == m.FuncName {

			start = fn.Type.Func
			end = fn.Body.Rbrace
		}
	}

	b, e := ioutil.ReadFile(m.ReadPath.MakePath())
	if e != nil {
		fmt.Println("read file error")
		return
	}
	cunS := string(b[start-1: end])

	r := regexp.MustCompile("T")

	// regexp包也可以用来将字符串的一部分替换为其他的值
	res := r.ReplaceAllString(cunS, "{{.}}")

	t := template.Must(template.New("test").Parse(res))
	m.Template = t

	path := lib.NewPath("/code/" + m.FuncName + ".go")
	fileName := path.MakePath()

	var file *os.File

	if checkFileIsExist(fileName) { //如果文件存在
		_, err := os.OpenFile(fileName, os.O_APPEND, 0666) //打开文件
		if err != nil {
			panic(err)
		}
		os.Remove(fileName)
	}
	file, err = os.Create(fileName) //创建文件
	if err != nil {
		panic(err)
	}

	io.WriteString(file, "package code\n")

	for _, str := range TYPE_STRING {
		t.Execute(file, str)
		io.WriteString(file, "\n")
	}
}

func Make() {

	num := int(3)
	numP := &num
	numX := &numP
	fmt.Println(reflect.TypeOf(numX).String())
	fmt.Println(**numX)

	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, lib.GetRoot()+"/Func/AST.go", nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	var start token.Pos
	var end token.Pos

	for _, decl := range f.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if ok && fn.Name.Name == "CompareT" {

			start = fn.Type.Func
			end = fn.Body.Rbrace
		}
	}

	path := lib.NewPath("/Func/AST.go")

	b, e := ioutil.ReadFile(path.MakePath())
	if e != nil {
		fmt.Println("read file error")
		return
	}
	cunS := string(b[start-1: end])

	r := regexp.MustCompile("T")

	// regexp包也可以用来将字符串的一部分替换为其他的值
	res := r.ReplaceAllString(cunS, "{{.}}")

	t := template.Must(template.New("test").Parse(res))

	fileName := lib.GetRoot() + "/code/max.go"

	var file *os.File

	if checkFileIsExist(fileName) { //如果文件存在
		_, err := os.OpenFile(fileName, os.O_APPEND, 0666) //打开文件
		if err != nil {
			panic(err)
		}
		os.Remove(fileName)
	}
	file, err = os.Create(fileName) //创建文件
	if err != nil {
		panic(err)
	}

	io.WriteString(file, "package code\n")

	for _, str := range TYPE_STRING {
		t.Execute(file, str)
		io.WriteString(file, "\n")
	}

}

func checkFileIsExist(filename string) (bool) {
	var exist = true;
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false;
	}
	return exist;
}
