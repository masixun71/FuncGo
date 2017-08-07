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
	"os/exec"
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

type TypeUsePointer struct {
	IsUse  bool
	UseStr string
}

type ObjectReplace struct {
	ReplaceObject string
	IsUse         bool
}

type MakeFiler struct {
	GenerateObject   *ObjectReplace
	BuildType        *BuildType
	OutputDir        lib.Path
	TypePointer      *TypeUsePointer
	TMaper           TypeDependent
	SpecialOperation uint
}

func NewTypeUsePointer(typeStr string) *TypeUsePointer {

	if len(typeStr) == 0 {
		return &TypeUsePointer{IsUse: false, UseStr: typeStr}
	} else {
		return &TypeUsePointer{IsUse: true, UseStr: typeStr}
	}
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

func NewMakeFilerByBasicType(replace int, outputDir, TypePointer string) (MakeFile, error) {

	buildType, path := checkMakeFilerForBuildType(replace, outputDir)

	return &MakeFiler{BuildType: buildType, OutputDir: path, GenerateObject: &ObjectReplace{IsUse: false}, TypePointer: NewTypeUsePointer(TypePointer), TMaper: TmapByBasicType()}, nil
}

func NewMakeFiler(buildType *BuildType, outputDir string, usePointer bool, TypePointer string, tmap TypeDependent, specialOperation uint) (MakeFile, error) {

	if strings.ToUpper(buildType.FuncString) == strings.ToUpper(buildType.TypeString) {
		return nil, errors.New("buildType.FuncString must not equal buildType.TypeString")
	}

	path := lib.NewPath(outputDir)
	if !path.IsDir() {
		return nil, errors.New("outputDir can not find or not dir")
	}

	return &MakeFiler{BuildType: buildType, OutputDir: path, GenerateObject: &ObjectReplace{IsUse: usePointer}, TypePointer: NewTypeUsePointer(TypePointer), TMaper: tmap, SpecialOperation: specialOperation}, nil
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
	if m.GenerateObject.IsUse {
		m.GenerateObject.ReplaceObject = "*" + strings.TrimLeft(saveStr, "*")
	} else {
		m.GenerateObject.ReplaceObject = strings.TrimLeft(saveStr, "*")
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
			resetBuf, _ := m.doForMultiplePointers(fn, buf)

			_, err := m.makeCode(resetBuf, fn.Name.Name)
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

	for _, decl := range f.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if ok && fn.Name.Name == funcName {

			resetBuf, _ := m.doForMultiplePointers(fn, b)
			return m.makeCode(resetBuf, funcName)
		}
	}

	return false, errors.New("can not find func")
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
		} else {
			tmpFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
			file = tmpFile
			if err != nil {
				return false, err
			}
		}

		var buffer bytes.Buffer
		if m.GenerateObject.IsUse {
			buffer.Write(cunS[0:5])
			buffer.WriteString("(f " + m.GenerateObject.ReplaceObject + ") ")
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

		exec.Command("gofmt", fileName)

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

		if m.TypePointer.IsUse {
			realFuncName = rF.ReplaceAllString(funcName, PointerTypeToFuncName(str, m.TypePointer.UseStr))
		} else {
			realFuncName = rF.ReplaceAllString(funcName, TypeToFuncName(str))
		}
		_, ok := mapObject[realFuncName]
		if ok {
			mapFunc[realFuncName] = 1
		} else {
			if m.TypePointer.IsUse {
				arrType = append(arrType, BuildType{TypeString: str, FuncString: PointerTypeToFuncName(str, m.TypePointer.UseStr)})
			} else {
				arrType = append(arrType, BuildType{TypeString: str, FuncString: TypeToFuncName(str)})
			}
		}
	}

	funcCache[filename] = 1

	return arrType, nil
}

func (m *MakeFiler) doForMultiplePointers(fn *ast.FuncDecl, buf []byte) ([]byte, error) {

	if !m.TypePointer.IsUse {
		return buf[fn.Type.Func-1: fn.Body.Rbrace], nil
	}

	str := string(buf[fn.Type.Params.Opening:fn.Type.Params.Closing])
	savePointerStr := make([]string, 1)
	fileBuffer := new(bytes.Buffer)
	var buffer bytes.Buffer

	var typName string
	for _, field := range fn.Type.Params.List {
		for _, name := range field.Names {
			//varLen := len(name.Name)
			index := strings.Index(str, name.Name)
			str = str[0:index] + "p" + str[index:len(str)]

			switch paramsType := field.Type.(type) {
			case *ast.Ident:
				typName = paramsType.Name
				savePointerStr = append(savePointerStr, "\n\t"+name.Name+" := "+m.TypePointer.UseStr+"p"+name.Name+"\n")
			case *ast.Ellipsis:
				if ident, ok := paramsType.Elt.(*ast.Ident); ok {
					typName = ident.Name
					savePointerStr = append(savePointerStr, "\n\t"+name.Name+" := make([]"+ident.Name+", 0)\n\tfor _,value := range p"+name.Name+" { "+name.Name+" = append("+name.Name+", "+m.TypePointer.UseStr+"value)}\n")
				}
			default:
				panic("params type must be Ident or Ellipsis")
			}
		}

		rF := regexp.MustCompile(typName)
		// regexp包也可以用来将字符串的一部分替换为其他的值
		resString := rF.ReplaceAllString(str, "{{.}}")
		tem := template.Must(template.New("field").Parse(resString))
		tem.Execute(fileBuffer, m.TypePointer.UseStr+typName)

	}

	rootStr := string(buf[fn.Type.Func-1: fn.Body.Rbrace])
	buffer.WriteString(rootStr[0:fn.Type.Params.Opening-(fn.Type.Func-1)])
	buffer.WriteString(fileBuffer.String())
	index := strings.Index(rootStr, "{")
	buffer.WriteString(rootStr[fn.Type.Params.Closing-(fn.Type.Func-1):index+1])
	for _, pointerStr := range savePointerStr {
		buffer.WriteString(pointerStr)
	}
	buffer.WriteString(rootStr[index+1:len(rootStr)])

	return buffer.Bytes(), nil
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
