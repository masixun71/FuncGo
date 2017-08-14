package Func

import (
	"go/token"
	"go/parser"
	"go/ast"
	"FuncGo/lib"
	"os"
	"io/ioutil"
	"errors"
	"fmt"
	"strings"
	"text/template"
)

const (
	IDENT    = 1
	ELLIPSIS = 2
)
//todo Switch function is too limited, temporarily write a beta version




func (m *MakeFiler) AddSwitchWithString(switchName, str string) (bool, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", str, parser.ParseComments)
	if err != nil {
		return false, err
	}

	for _, decl := range f.Decls {
		var saveTypeMap = make(map[string]bool)
		var returnFuncStr string
		switchType := 0

		fn, ok := decl.(*ast.FuncDecl)
		if ok {
			returnFuncStr = str[fn.Name.NamePos-1:fn.Type.Params.Opening]

			for _, params := range fn.Type.Params.List {

				switch paramsType := params.Type.(type) {
				case *ast.Ident:
					if switchType != 0 {
						switchType = IDENT
					}

					for _, name := range params.Names {
						returnFuncStr += name.Name + ","
					}
					saveTypeMap[paramsType.Name] = true
				case *ast.Ellipsis:
					if switchType != 0 {
						switchType = ELLIPSIS
					}
					for _, name := range params.Names {
						returnFuncStr += name.Name + "...,"
					}

					ident, ok := paramsType.Elt.(*ast.Ident)
					if ok {
						saveTypeMap[ident.Name] = true
					}

				default:
					panic("params type must be Ident or Ellipsis")
				}
			}

			returnFuncStr = strings.TrimRight(returnFuncStr, ",")
			returnFuncStr += ")"

			if len(saveTypeMap) > 1 {
				panic("type is more")
			}


			file, err := os.OpenFile(m.OutputDir.GetPathByRoot() + "/" + fn.Name.Name + ".go", os.O_RDWR, 0777)
			defer file.Close()
			if err != nil {
				return false, err
			}
			readFset := token.NewFileSet()
			readF, err := parser.ParseFile(readFset, "", file, parser.ParseComments)
			if err != nil {
				return false, err
			}

			caseStr := `
				case ` + m.BuildType.TypeString + `:
					return ` + returnFuncStr + `
				`

			caseTem := m.getTemplate(caseStr)

			arrType := make([]BuildType, 0)
			for _, str := range *m.TMaper.GetMap() {
				if m.TypePointer.IsUse {
					arrType = append(arrType, BuildType{TypeString: str, FuncString: PointerTypeToFuncName(str, m.TypePointer.UseStr)})
				} else {
					arrType = append(arrType, BuildType{TypeString: str, FuncString: TypeToFuncName(str)})
				}
			}

			if _, ok := readF.Scope.Objects[switchName]; ok {
				for _, decl := range f.Decls {
					fn, ok := decl.(*ast.FuncDecl)
					if ok {
						if fn.Name.Name == switchName {
							for _, list := range fn.Body.List {
								if switchStmt, ok := list.(*ast.TypeSwitchStmt); ok {
									fmt.Println(switchStmt)
									file.Seek(int64(switchStmt.Body.Rbrace), 0)
								}
							}
						}
					}
				}
			} else {
				file.Seek(0, 2)

				file.WriteString("func " + switchName + "(values ...interface{})  (T, error) {\n")
				file.WriteString("switch getvalue := values.(type) {\n")
				addCase(caseTem, &arrType, file)
				file.WriteString("}\npanic(\"can't find type\")\n}\n")
			}

		}
	}

	//ast.Print(fset, f)

	return true, nil
}

func addCase(tem *template.Template, arrType *[]BuildType, file *os.File) {

	for _, typ := range *arrType {
		tem.Execute(file, typ)
	}
}


func (m *MakeFiler) AddSwitchWithFunc(switchName, funcName string, readPath lib.Path) (bool, error) {
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

func (m *MakeFiler) AddSwitchWithFile(switchName string, reader *os.File) (bool, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", reader, parser.ParseComments)
	if err != nil {
		return false, err
	}
	ast.Print(fset, f)

	b, e := ioutil.ReadFile(reader.Name())
	fmt.Println(b)
	if e != nil {
		return false, err
	}

	return true, nil
}
