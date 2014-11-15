package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	sourceFile, err := ioutil.ReadFile(os.Args[1])
	var s string
	if err == nil {
		s = string(sourceFile)
		fmt.Println(s)
	}
	splitFile := strings.Split(s, "\n")
	splitFile = removeNewlines(splitFile)
	lineSlice := createLines(splitFile)
	debugFile, _ := os.Create("testfiles/c++/debug.cpp")

	for i := 0; i < len(lineSlice); i++ {
		debugFile.WriteString(lineSlice[i].s + "\n")
		AddPrint(debugFile, lineSlice[i])
	}
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

func AddPrintc(file *os.File, s string) {
	file.WriteString("printf(\"" + s + "\n\");\n");
}

func createLines(stringSlice []string) []lineType {
	var lineSlice []lineType
	for i := 0; i < len(stringSlice); i++ {
		lineSlice = append(lineSlice, lineType{stringSlice[i], 0})
	}
	return lineSlice
}

func removeNewlines(stringSlice []string) []string {
	var newSlice []string
	for i := 0; i < len(stringSlice); i++ {
		if stringSlice[i] != "" && stringSlice[i][0] != 10 {
			newSlice = append(newSlice, stringSlice[i])
		}
	}
	return newSlice
}
