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

const (
	None             = 0
	MultiplePointers = 1
	FuncGlobalSwitch = 1 << 1
)

var mapFunc map[string]int = make(map[string]int, 2)
var funcCache map[string]int = make(map[string]int, 2)

type MakeFile interface {
	SetSpecialOperation(specialOperation uint)
	MakeFuncSourceWithString(str string) (bool, error)
	MakeFuncSourceWithFile(reader *os.File) (bool, error)
	MakeFuncSourceWithFunc(readPath lib.Path, funcName string) (bool, error)
	MakeMethodSourceWithFunc(valueS interface{}, readPath lib.Path, funcName string) (bool, error)
	MakeMethodSourceWithString(valueS interface{}, str string) (bool, error)
	MakeMethodSourceWithFile(valueS interface{}, reader *os.File) (bool, error)
}

type BuildType struct {
	FuncString string
	TypeString string
}

type MakeFiler struct {
	ReplaceObject    string
	BuildType        *BuildType
	OutputDir        lib.Path
	UsePointer       bool
	TMaper           TypeDependent
	SpecialOperation uint
}

func getSimpleType(replace int) (*BuildType, error) {
	var buildType BuildType

	switch replace {
	case TypeT:
		buildType = BuildType{FuncString: "TF", TypeString: "T"}
	case TypeInterface:
		buildType = BuildType{FuncString: "TFInterface", TypeString: "TInterface"}
	default:
		return nil, errors.New("Unknown Type")
	}

	return &buildType, nil
}

func checkMakeFilerForBuildType(replace int, outputDir string) (*BuildType, lib.Path) {
	buildType, err := getSimpleType(replace)
	if err != nil {
		panic(err)
	}

	path := lib.NewPath(outputDir)
	if !path.IsDir() {
		panic("outputDir can not find or not dir")
	}

	return buildType, path
}

func NewMakeFilerByBasicType(replace int, outputDir string) (MakeFile, error) {

	buildType, path := checkMakeFilerForBuildType(replace, outputDir)

	return &MakeFiler{BuildType: buildType, OutputDir: path, UsePointer: false, TMaper: TmapByBasicType()}, nil
}

func NewMakeFilerByBasicPointerType(replace int, outputDir string) (MakeFile, error) {
	buildType, path := checkMakeFilerForBuildType(replace, outputDir)
	return &MakeFiler{BuildType: buildType, OutputDir: path, UsePointer: true, TMaper: TmapByBasicType()}, nil
}

func NewMakeFilerSimple(replace int, outputDir string, usePointer bool, tmap TypeDependent) (MakeFile, error) {
	buildType, path := checkMakeFilerForBuildType(replace, outputDir)
	return &MakeFiler{BuildType: buildType, OutputDir: path, UsePointer: usePointer, TMaper: tmap}, nil
}

func NewMakeFiler(buildType *BuildType, outputDir string, usePointer bool, tmap TypeDependent) (MakeFile, error) {

	if strings.ToUpper(buildType.FuncString) == strings.ToUpper(buildType.TypeString) {
		return nil, errors.New("buildType.FuncString must not equal buildType.TypeString")
	}

	path := lib.NewPath(outputDir)
	if !path.IsDir() {
		return nil, errors.New("outputDir can not find or not dir")
	}

	return &MakeFiler{BuildType: buildType, OutputDir: path, UsePointer: usePointer, TMaper: tmap}, nil
}

func (m *MakeFiler) SetSpecialOperation(specialOperation uint) {
	m.SpecialOperation = specialOperation
}

func (m *MakeFiler) MakeMethodSourceWithFunc(valueS interface{}, readPath lib.Path, funcName string) (bool, error) {
	m.setReplaceObject(valueS)

	return m.MakeFuncSourceWithFunc(readPath, funcName)
}

func (m *MakeFiler) MakeMethodSourceWithString(valueS interface{}, str string) (bool, error) {
	m.setReplaceObject(valueS)

	return m.MakeFuncSourceWithString(str)
}

func (m *MakeFiler) MakeMethodSourceWithFile(valueS interface{}, reader *os.File) (bool, error) {
	m.setReplaceObject(valueS)

	return m.MakeFuncSourceWithFile(reader)
}

func (m *MakeFiler) setReplaceObject(valueS interface{}) {

	var saveStr string

	switch s := valueS.(type) {
	case string:
		saveStr = s
	default:
		saveStr = reflect.TypeOf(valueS).String()
	}

	reflect.TypeOf(valueS).String()
	if m.UsePointer {
		m.ReplaceObject = "*" + strings.TrimLeft(saveStr, "*")
	} else {
		m.ReplaceObject = strings.TrimLeft(saveStr, "*")
	}
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

		m.doForSpecialOpearation(file, arrType)
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
	for _, str := range *m.TMaper.GetMap() {
		var realFuncName string

		if m.UsePointer {
			realFuncName = rF.ReplaceAllString(funcName, PointerTypeToFuncName(str))
		} else {
			realFuncName = rF.ReplaceAllString(funcName, TypeToFuncName(str))
		}
		_, ok := mapObject[realFuncName]
		if ok {
			mapFunc[realFuncName] = 1
		} else {
			arrType = append(arrType, BuildType{TypeString: str, FuncString: PointerTypeToFuncName(str)})
		}
	}

	funcCache[filename] = 1

	return arrType, nil
}

func (m *MakeFiler) doForSpecialOpearation(file *os.File, buildTypes []BuildType) {
	//if m.SpecialOperation == None {
	//	return
	//}
	//fset := token.NewFileSet()
	//f, err := parser.ParseFile(fset, file.Name(), nil, parser.ParseComments)
	//if err != nil {
	//	panic(err)
	//}

}
