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
	lineSlice := createLines(splitFile)
	debugFile, err := os.Create("testfiles/c++/debug.cpp")
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(lineSlice); i++ {
		_, err := debugFile.WriteString(splitFile[i] + "\n")
		if err != nil {
			panic(err)
		}
		AddPrint(debugFile, lineSlice[i].s)
	}
}

type lineType struct {
	s string
	code int
}

func AddPrint(file *os.File, s string) {
	file.WriteString("std::cout << \"" + s + "\" std::endl;\n")
}

func createLines(stringSlice []string) ([]lineType) {
	var lineSlice []lineType
	for i := 0; i < len(stringSlice); i++ {
		lineSlice = append(lineSlice, lineType{stringSlice[i], 0})
	}
	return lineSlice
}
