package Func

import (
	"FuncGo/lib"
	"fmt"
	"io/ioutil"
	"regexp"
	"text/template"
	"os"
	"strings"
	"io"
)

func Make() {

	//fset := token.NewFileSet()

	//f, err := parser.ParseFile(fset, lib.GetRoot()+"/Func/AST.go", nil, parser.ParseComments)
	//if err != nil {
	//	panic(err)
	//}

	b, e := ioutil.ReadFile(lib.GetRoot() + "/Func/AST.go")
	if e != nil {
		fmt.Println("read file error")
		return
	}

	arr := strings.Split(string(b), "\n")

	cunS := ""

	cunS += arr[3] + "\n"
	cunS += arr[4] + "\n"
	cunS += arr[5] + "\n"
	cunS += arr[6] + "\n"
	cunS += arr[7] + "\n"
	cunS += arr[8] + "\n"
	cunS += arr[9] + "\n"
	cunS += arr[10] + "\n"
	cunS += arr[11] + "\n"
	cunS += arr[12] + "\n"
	cunS += arr[13] + "\n"
	cunS += arr[14] + "\n"

	r := regexp.MustCompile("int")

	// regexp包也可以用来将字符串的一部分替换为其他的值
	res := r.ReplaceAllString(cunS, "{{.}}")


	s := "int32"
	t := template.Must(template.New("test").Parse(res))

	fileName := lib.GetRoot() + "/code/max.go"

	var file *os.File

	if checkFileIsExist(fileName) { //如果文件存在
		f, err := os.OpenFile(fileName, os.O_APPEND, 0666) //打开文件
		file = f
		if err != nil {
			panic(err)
		}
	} else {
		f, err := os.Create(fileName) //创建文件
		file =f
		if err != nil {
			panic(err)
		}
	}

	io.WriteString(file, "package code\n")


	t.Execute(file, s)

}

func checkFileIsExist(filename string) (bool) {
	var exist = true;
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false;
	}
	return exist;
}
