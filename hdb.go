package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	sourceFile, err := ioutil.ReadFile(os.Args[1])
	var s string
	var cFile bool = (filepath.Ext(os.Args[1]) == "c")

	if err == nil {
		s = string(sourceFile)
		//fmt.Println(s)
	}
	s = Removecomments(s);
	splitFile := strings.Split(s, "\n")
	splitFile = RemoveNewlines(splitFile)
	lineSlice := CreateLines(splitFile)
	MarkInvalid(lineSlice)
	debugFile, _ := os.Create("testfiles/c++/debug.cpp")

	for i := 0; i < len(lineSlice); i++ {
		AddPrint(debugFile, lineSlice[i], cFile)
	}
	debugFile.Close()
	CompileAndRun()
}

type lineType struct {
	s    string
	code int
}

func AddPrint(file *os.File, line lineType, cFile bool) {
	if cFile {
		if line.code == 0 {
			file.WriteString("printf(\"" + line.s + "\n\");\n")
			file.WriteString(line.s + "\n")
		}
	} else {
		if line.code == 0 {
			file.WriteString("std::cout << \"" + line.s + "\" << std::endl;\n")
			file.WriteString(line.s + "\n")
		}
	}
}

func MarkInvalid(lineSlice []lineType) {

	r, _ := regexp.Compile(`(.*)\((.*)\)(.*)\{`)
	opencount := 0
	closecount := 0

	for i := 0; i < len(lineSlice); i++ {
		fmt.Println(opencount, closecount, lineSlice[i].code)
		if opencount <= closecount {
			lineSlice[i].code = 1
		}
		if r.MatchString(lineSlice[i].s) {
			opencount++
		} else if strings.TrimSpace(lineSlice[i].s)[0] == 125 {
			closecount++
			lineSlice[i].code = 1
		}
	}
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
	gccCmd := exec.Command("g++", "-Wall", "testfiles/c++/debug.cpp", "-o", "testfiles/c++/out")
	gccCmd.Run()
	runCmd := exec.Command("./testfiles/c++/out")
	output, err := runCmd.Output()
	fmt.Println(string(output))
	fmt.Println(err)
}

func Removecomments(s string) string{
     r, _ := regexp.Compile(`/\*(.*?)\*/|//(.*?)\n`)
     ret := r.ReplaceAllString(s, "")
     fmt.Println("TEST\n" + ret + "\nENDTEST\n")
     return ret
}
