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
	debugFile, err := os.Create("testfiles/c++/debug.cpp")
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(splitFile); i++ {
		_, err := debugFile.WriteString(splitFile[i] + "\n")
		if err != nil {
			panic(err)
		}
		addPrint(debugFile, splitFile[i])
	}
}

func addPrint(file *os.File, s string) {
	file.WriteString("std::cout << \"" + s + "\" std::endl;\n")
}
