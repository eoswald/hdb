package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"path/filepath"
	"strings"
)

func main() {
	sourceFile, err := ioutil.ReadFile(os.Args[1])
	var s string
	var cfile bool = (filepath.Ext(os.Args[1]) == "c")

	if err == nil {
		s = string(sourceFile)
		fmt.Println(s)
	}
	splitFile := strings.Split(s, "\n")
	splitFile = RemoveNewlines(splitFile)
	lineSlice := CreateLines(splitFile)
	debugFile, _ := os.Create("testfiles/c++/debug.cpp")

	for i := 0; i < len(lineSlice); i++ {
		debugFile.WriteString(lineSlice[i].s + "\n")
		if cfile {
			AddPrintc(debugFile, lineSlice[i].s)
		} else {
			AddPrint(debugFile, lineSlice[i])
		}
	}
	debugFile.Close()
	CompileAndRun()
}

type lineType struct {
	s    string
	code int
}

func AddPrint(file *os.File, line lineType) {
	if line.code == 0 {
		file.WriteString("std::cout << \"" + line.s + "\" std::endl;\n")
	}
}

func MarkInvalid(lineSlice []lineType) {
	r, _ := regexp.Compile(`(.*)\((.*)\)(.*)\{`)

	for i := 0; i < len(lineSlice); i++{
		if r.MatchString(lineSlice[i].s){
			lineSlice[i].code = 0
		}
	}
}

func AddPrintc(file *os.File, s string) {
	file.WriteString("printf(\"" + s + "\n\");\n")
}

func CreateLines(stringSlice []string) []lineType {
	var lineSlice []lineType
	for i := 0; i < len(stringSlice); i++ {
		lineSlice = append(lineSlice, lineType{stringSlice[i], 0})
	}
	return lineSlice
}

func RemoveNewlines(stringSlice []string) []string {
	var newSlice []string
	for i := 0; i < len(stringSlice); i++ {
		if stringSlice[i] != "" && stringSlice[i][0] != 10 {
			newSlice = append(newSlice, stringSlice[i])
		}
	}
	return newSlice
}

func CompileAndRun() {
	clangCmd := exec.Command("clang++", "-Wall", "testfiles/c++/debug.cpp", "-o", "testfiles/c++/out")
	gccCmd := exec.Command("g++", "-Wall", "testfiles/c++/debug.cpp", "-o", "testfiles/c++/out")
	err := clangCmd.Run()
	fmt.Println(err)
	if err != nil {
		gccCmd.Run()
	}
	runCmd := exec.Command("./testfiles/c++/out")
	err = runCmd.Run()
	fmt.Println(err)
}
