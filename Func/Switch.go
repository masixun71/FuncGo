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
)

func (m *MakeFiler) AddSwitchWithString(switchName , str string) (bool, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", str, parser.ParseComments)
	if err != nil {
		return false, err
	}

	var saveTypeMap = make(map[string]bool)


	for _, decl := range f.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if ok {
			fmt.Println(str[fn.Name.NamePos-1:fn.Type.Params.Opening])

			for _, params := range fn.Type.Params.List {
				for _, name := range params.Names {
					fmt.Println(name.Name)
				}

				switch paramsType := params.Type.(type) {
				case *ast.Ident:
					saveTypeMap[paramsType.Name] = true
				case *ast.Ellipsis:
					ident, ok  := paramsType.Elt.(*ast.Ident)
					if ok {
						saveTypeMap[ident.Name] = true
					}

				default:
					panic("params type must be Ident or Ellipsis")
				}
			}

			if len(saveTypeMap) > 1 {
				panic("type is more")
			}

		}
	}

	//ast.Print(fset, f)

	return true, nil
}

func (m *MakeFiler) AddSwitchWithFunc(switchName , funcName string, readPath lib.Path)  (bool, error) {
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


func (m *MakeFiler) AddSwitchWithFile(switchName string, reader *os.File)  (bool, error) {
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
