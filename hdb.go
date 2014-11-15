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
	debugFile, err := os.Create("debug.cpp")
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(splitFile); i++ {
		x, err := debugFile.WriteString(splitFile[i] + "\n")
		fmt.Println(x)
		if err != nil {
			panic(err)
		}
	}
}
