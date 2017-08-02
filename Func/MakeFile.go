package Func

import (
	"FuncGo/lib"
	"io/ioutil"
	"regexp"
	"text/template"
	"os"
	"io"
	"go/token"
	"go/parser"
	"go/ast"
	"bytes"
	"reflect"
	"strings"
	"errors"
)

var mapFunc map[string]int = make(map[string]int, 2)
var funcCache map[string]int = make(map[string]int, 2)

type MakeFile interface {
	MakeFuncSourceWithString(str string) (bool, error)
	MakeFuncSourceWithFile(reader *os.File) (bool, error)
	MakeFuncSourceWithFunc(readPath lib.Path, funcName string) (bool, error)
	MakeMethod(valueS interface{}, readPath lib.Path, funcName string) (bool, error)
}

type BuildType struct {
	FuncString string
	TypeString string
}

type MakeFiler struct {
	ReplaceObject string
	BuildType     *BuildType
	OutputDir     lib.Path
}

func NewMakeFilerSimple(replace int, outputDir string) (MakeFile, error) {

	var buildType BuildType

	switch replace {
	case TypeT:
		buildType = BuildType{FuncString: "TF", TypeString: "T"}
	case TypeInterface:
		buildType = BuildType{FuncString: "TFInterface", TypeString: "TInterface"}
	default:
		return nil, errors.New("Unknown Type")
	}

	path := lib.NewPath(outputDir)
	if !path.IsDir() {
		return nil, errors.New("outputDir can not find or not dir")
	}

	return &MakeFiler{BuildType: &buildType, OutputDir: path}, nil
}

func NewMakeFiler(buildType *BuildType, outputDir string) (MakeFile, error) {

	if strings.ToUpper(buildType.FuncString) == strings.ToUpper(buildType.TypeString) {
		return nil, errors.New("buildType.FuncString must not equal buildType.TypeString")
	}

	path := lib.NewPath(outputDir)
	if !path.IsDir() {
		return nil, errors.New("outputDir can not find or not dir")
	}

	return &MakeFiler{BuildType: buildType, OutputDir: path}, nil
}

func (m *MakeFiler) MakeMethod(valueS interface{}, readPath lib.Path, funcName string) (bool, error) {

	arrStr := strings.Split(reflect.TypeOf(valueS).String(), ".")
	m.ReplaceObject = arrStr[len(arrStr)-1]
	return m.MakeFuncSourceWithFunc(readPath, funcName)
}

func (m *MakeFiler) MakeFuncSourceWithString(str string) (bool, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", str, parser.ParseComments)
	if err != nil {
		return false, err
	}
	return m.makeCodeForAst(f, []byte(str))
}
func (m *MakeFiler) MakeFuncSourceWithFile(reader *os.File) (bool, error) {

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", reader, parser.ParseComments)
	if err != nil {
		return false, err
	}

	b, e := ioutil.ReadFile(reader.Name())
	if e != nil {
		return false, err
	}

	return m.makeCodeForAst(f, b)
}

func (m *MakeFiler) makeCodeForAst(f *ast.File, buf []byte) (bool, error) {
	for _, decl := range f.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if ok {
			_, err := m.makeCode(buf[fn.Type.Func-1: fn.Body.Rbrace], fn.Name.Name)
			if err != nil {
				return false, err
			}
		}
	}

	return true, nil
}

func (m *MakeFiler) MakeFuncSourceWithFunc(readPath lib.Path, funcName string) (bool, error) {
	b, err := ioutil.ReadFile(readPath.GetPathByRoot())
	if err != nil {
		return false, err
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, readPath.GetPathByRoot(), nil, parser.ParseComments)
	if err != nil {
		return false, err
	}

	var start token.Pos
	var end token.Pos

	for _, decl := range f.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if ok && fn.Name.Name == funcName {

			start = fn.Type.Func
			end = fn.Body.Rbrace
		}
	}

	return m.makeCode(b[start-1: end], funcName)
}

func (m *MakeFiler) makeCode(str []byte, funcName string) (bool, error) {
	fileName := m.getFileName(funcName)
	return m.makeFileByString(str, fileName, funcName)
}

func (m *MakeFiler) makeFileByString(cunS []byte, fileName, funcName string) (bool, error) {

	var file *os.File
	defer file.Close()
	if funcCache[fileName] != 1 {
		if !checkFileIsExist(fileName) { //如果文件存在
			tmpFile, err := os.Create(fileName) //创建文件
			file = tmpFile
			if err != nil {
				return false, err
			}
			io.WriteString(file, "package "+strings.Trim(m.OutputDir.GetPath(), "/")+"\n\n")
			if len(m.ReplaceObject) != 0 {
				io.WriteString(file, "")
			}
		} else {
			tmpFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
			file = tmpFile
			if err != nil {
				return false, err
			}
		}

		var buffer bytes.Buffer
		if len(m.ReplaceObject) != 0 {
			buffer.Write(cunS[0:5])
			buffer.WriteString("(f " + m.ReplaceObject + ") ")
			buffer.Write(cunS[5:len(cunS)])
		} else {
			buffer.Write(cunS)
		}

		tem := m.getTemplate(buffer.String())

		arrType, err := m.checkFuncInit(fileName, funcName)
		if err != nil {
			return false, err
		}
		for _, buildType := range arrType {
			err := tem.Execute(file, buildType)
			if err != nil {
				return false, err
			}
			mapFunc[funcName+buildType.FuncString] = 1
			io.WriteString(file, "\n\n")
		}
	}

	return true, nil
}

func checkFileIsExist(filename string) (bool) {
	var exist = true;
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false;
	}
	return exist;
}

func (m *MakeFiler) getTemplate(str string) *template.Template {
	rF := regexp.MustCompile(m.BuildType.FuncString)

	// regexp包也可以用来将字符串的一部分替换为其他的值
	resString := rF.ReplaceAllString(str, "{{.FuncString}}")

	rT := regexp.MustCompile(m.BuildType.TypeString)

	// regexp包也可以用来将字符串的一部分替换为其他的值
	res := rT.ReplaceAllString(resString, "{{.TypeString}}")

	return template.Must(template.New("test").Parse(res))
}

func (m *MakeFiler) getFileName(funcName string) string {
	path := lib.NewPath(strings.TrimRight(m.OutputDir.GetPath(), "/") + "/" + funcName + ".go")
	return path.GetPathByRoot()
}

func (m *MakeFiler) checkFuncInit(filename, funcName string) ([]BuildType, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	arrType := make([]BuildType, 0)

	rF := regexp.MustCompile(m.BuildType.FuncString)
	mapObject := f.Scope.Objects
	for _, str := range TYPE_STRING {
		realFuncName := rF.ReplaceAllString(funcName, TypeToFuncName(str))
		_, ok := mapObject[realFuncName]
		if ok {
			mapFunc[realFuncName] = 1
		} else {
			arrType = append(arrType, BuildType{TypeString: str, FuncString: TypeToFuncName(str)})
		}
	}

	funcCache[filename] = 1

	return arrType, nil
}
