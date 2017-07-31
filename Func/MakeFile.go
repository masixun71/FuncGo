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
)

var initTsruct bool = false
var mapFunc map[string]int = make(map[string]int, 2)
var funcCache map[string]int = make(map[string]int, 2)

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

	if !initTsruct {
		makeTStruct()
	}

	path := lib.NewPath("/code/" + m.FuncName + ".go")
	fileName := path.MakePath()

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

	var file *os.File

	if funcCache[fileName] != 1 {
		if !checkFileIsExist(fileName) { //如果文件存在
			file, err = os.Create(fileName) //创建文件
			if err != nil {
				panic(err)
			}
			io.WriteString(file, "package code\n")
		}
		arrType := m.checkFuncInit(fileName)
		for _, str := range arrType {
			t.Execute(file, str)
			mapFunc[m.FuncName+str] = 1
			io.WriteString(file, "\n\n")
		}
	}

}

func checkFileIsExist(filename string) (bool) {
	var exist = true;
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false;
	}
	return exist;
}

func (m *MakeFiler) checkFuncInit(filename string) []string {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	arrType := make([]string, 0)

	mapObject := f.Scope.Objects
	for _, str := range TYPE_STRING {
		_, ok := mapObject[m.FuncName+str]
		if ok {
			mapFunc[m.FuncName+str] = 1
		} else {
			arrType = append(arrType, str)
		}
	}

	funcCache[filename] = 1

	return arrType
}

func makeTStruct() {
	path := lib.NewPath("/code/" + "TStruct" + ".go")
	fileName := path.MakePath()

	if !checkFileIsExist(fileName) { //如果文件存在
		file, err := os.Create(fileName) //创建文件
		if err != nil {
			panic(err)
		}
		io.WriteString(file, `package code
type ts struct {}`)
	}

	initTsruct = true
}
