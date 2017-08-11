package main

import (
	"FuncGo/Func"
	"fmt"
)

func main() {

	makeFiler, _ := Func.NewMakeFilerByBasicType(Func.TypeT, "/code", "*")

	_,err :=makeFiler.MakeFuncSourceWithString(`
		package Func
		func MaxTF(values ...T) (T, error) {

		len := len(values)

		max := values[0]
		for i := 1; i < len; i++ {
			tmpvalue := values[i]

			if max < tmpvalue {
				max = tmpvalue
			}
		}

		return max, nil
	}`)
	if err != nil {
		fmt.Println(err)
	}

	makeFiler.AddSwitchWithString("switchMax", `
		package Func
		func MaxTF(values ...T) (T, error) {

		len := len(values)

		max := values[0]
		for i := 1; i < len; i++ {
			tmpvalue := values[i]

			if max < tmpvalue {
				max = tmpvalue
			}
		}

		return max, nil
	}`)
	//makeFiler, err := Func.NewMakeFilerByBasicType(Func.TypeT, "/code")

	//path := lib.NewPath("/Func/ast1.go")
	//
	//
	//pointer := "***"
	//
	//fset := token.NewFileSet()
	//f, _ := parser.ParseFile(fset, path.GetPathByRoot(), nil, parser.ParseComments)
	//b, _ := ioutil.ReadFile(path.GetPathByRoot())
	//
	//for _, decl := range f.Decls {
	//	fn, ok := decl.(*ast.FuncDecl)
	//	if ok && fn.Name.Name == "MaxTFP" {
	//
	//		str := string(b[fn.Type.Params.Opening:fn.Type.Params.Closing])
	//		savePointerStr := make([]string, 1)
	//		fileBuffer := new(bytes.Buffer)
	//		var buffer bytes.Buffer
	//
	//		var typName string
	//		for _, field := range fn.Type.Params.List {
	//			for _, name := range field.Names {
	//				//varLen := len(name.Name)
	//				index := strings.Index(str, name.Name)
	//				str = str[0:index] + "p" + str[index:len(str)]
	//
	//				switch paramsType := field.Type.(type) {
	//				case *ast.Ident:
	//					typName = paramsType.Name
	//					savePointerStr = append(savePointerStr, "\n\t"+name.Name+" := "+pointer+"p"+name.Name+"\n")
	//				case *ast.Ellipsis:
	//					if ident, ok := paramsType.Elt.(*ast.Ident); ok {
	//						typName = ident.Name
	//						savePointerStr = append(savePointerStr, "\n\t"+name.Name+" := make([]"+ident.Name+", 0)\n\tfor _,value := range p"+name.Name+" { "+name.Name+" = append("+name.Name+", "+pointer+"value)}\n")
	//					}
	//				default:
	//					panic("params type must be Ident or Ellipsis")
	//				}
	//			}
	//
	//			rF := regexp.MustCompile(typName)
	//			// regexp包也可以用来将字符串的一部分替换为其他的值
	//			resString := rF.ReplaceAllString(str, "{{.}}")
	//			tem := template.Must(template.New("field").Parse(resString))
	//			tem.Execute(fileBuffer, pointer + typName)
	//
	//		}
	//
	//		rootStr := string(b[fn.Type.Func-1: fn.Body.Rbrace])
	//		buffer.WriteString(rootStr[0:fn.Type.Params.Opening - (fn.Type.Func-1)])
	//		buffer.WriteString(fileBuffer.String())
	//		index := strings.Index(rootStr, "{")
	//		buffer.WriteString(rootStr[fn.Type.Params.Closing - (fn.Type.Func-1):index+1])
	//		for _, pointerStr := range savePointerStr {
	//			buffer.WriteString(pointerStr)
	//		}
	//		buffer.WriteString(rootStr[index+1:len(rootStr)])
	//		fmt.Println(buffer.String())
	//	}
	//}

}
