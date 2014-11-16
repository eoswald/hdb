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
	var lastLine bool
	if len(os.Args) > 2 && os.Args[2] == "--lastline=true" {
		lastLine = true
	}

	if err == nil {
		s = string(sourceFile)
		//fmt.Println(s)
	}
	s = Removecomments(s)
	splitFile := strings.Split(s, "\n")
	splitFile = RemoveNewlines(splitFile)
	lineSlice := CreateLines(splitFile)
	MarkInvalid(lineSlice)
	debugFile, _ := os.Create("testfiles/c++/debug.cpp")

	for i := 0; i < len(lineSlice); i++ {
		AddPrint(debugFile, lineSlice[i], cFile)
	}
	debugFile.Close()
	CompileAndRun(lastLine)
}

type lineType struct {
	s    string
	code int
	info string
}

func AddPrint(file *os.File, line lineType, cFile bool) {
	if cFile {
		if line.code == 0 {
			file.WriteString("printf(\"" + line.s + "\n\");\n")
			file.WriteString(line.s + "\n")
		}
	} else {
		switch line.code {
		case 0:
			file.WriteString("std::cout << \"" + line.s + "\" << std::endl;\n")
			file.WriteString(line.s + "\n")
		case 343: //function
			file.WriteString(line.s + "\n")
			//file.WriteString("Entering " + line.info + "\n")
		case 666: //if
			file.WriteString("std::cout << \"" + line.s + "\" << std::endl;\n")
			file.WriteString(line.s + "\n")
			//file.WriteString(line.info + "evaluates to true\n")
		case 752: //else if
			file.WriteString(line.s + "\n")
			//file.WriteString(line.info + "evaluates to true\n")
		case 580: //else
			file.WriteString(line.s + "\n")
			//file.WriteString("Else\n")
		case 4: //for
			file.WriteString("std::cout << \"" + line.s + "\" << std::endl;\n")
			file.WriteString(line.s + "\n")
			//file.WriteString("Looping\n")
		case 603: //while
			file.WriteString("std::cout << \"" + line.s + "\" << std::endl;\n")
			file.WriteString(line.s + "\n")
			//file.WriteString(line.info + "evaluates to true, looping")
		case 1:
			file.WriteString(line.s + "\n")
		}
	}
}
func MarkInvalid(lineSlice []lineType) {

	elseifR, _ := regexp.Compile(`(\s*)\}(\s*)else(\s*)if(\s*)\((.*)\)(\s*)\{(\s*)`)
	ifR, _ := regexp.Compile(`(\s*)if(\s*)\((.*)\)(\s*)\{(\s*)`)
	forR, _ := regexp.Compile(`(\s*)for(\s*)\((.*)\)(\s*)\{(\s*)`)
	whileR, _ := regexp.Compile(`(\s*)while(\s*)\((.*)\)(\s*)\{(\s*)`)
	funcR, _ := regexp.Compile(`(.*)\((.*)\)(\s*)\{(\s*)`)

	classR, _ := regexp.Compile(`(\s*)class(.*)\{`)
	structR, _ := regexp.Compile(`(\s*)struct(.*)\{`)
	elseR, _ := regexp.Compile(`(\s*)\}(\s*)else(\s*)\{`)
	scopeR, _ := regexp.Compile(`(\s*)\{(\s*)`)
	closeR, _ := regexp.Compile(`(\s*)\}(\s*)`)
	stack := []int{}

	for i := 0; i < len(lineSlice); i++ {
		if elseifR.MatchString(lineSlice[i].s) {
			//elseif
			lineSlice[i].code = 752
			PopStack(stack)
			stack = append(stack, 752)
		} else if ifR.MatchString(lineSlice[i].s) {
			//if
			lineSlice[i].code = 666
			stack = append(stack, 666)
		} else if forR.MatchString(lineSlice[i].s) {
			//for
			lineSlice[i].code = 4
			stack = append(stack, 4)
		} else if whileR.MatchString(lineSlice[i].s) {
			//while
			lineSlice[i].code = 603
			stack = append(stack, 603)
		} else if funcR.MatchString(lineSlice[i].s) {
			//func
			lineSlice[i].code = 343
			stack = append(stack, 343)
		} else if classR.MatchString(lineSlice[i].s) {
			//class
			lineSlice[i].code = 1
			stack = append(stack, 1)
		} else if structR.MatchString(lineSlice[i].s) {
			//struct
			lineSlice[i].code = 1
			stack = append(stack, 1)
		} else if elseR.MatchString(lineSlice[i].s) {
			//else
			lineSlice[i].code = 580
			PopStack(stack)
			stack = append(stack, 580)
		} else if scopeR.MatchString(lineSlice[i].s) {
			//scope
			lineSlice[i].code = 0
			stack = append(stack, 0)
		} else if closeR.MatchString(lineSlice[i].s) {
			//close
			lineSlice[i].code = 1
			PopStack(stack)
		} else {
			if InFunction(stack) {
				lineSlice[i].code = 0
			} else {
				lineSlice[i].code = 1
			}
		}
	}
}

func InFunction(stack []int) (exists bool) {
	for i := 0; i < len(stack); i++ {
		if stack[i] == 343 {
			return true
		}
	}
	return false
}

func PopStack(stack []int) {
	stack = stack[:len(stack)-1]
}

func CreateLines(stringSlice []string) []lineType {
	var lineSlice []lineType
	for i := 0; i < len(stringSlice); i++ {
		lineSlice = append(lineSlice, lineType{stringSlice[i], 0, ""})
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

func CompileAndRun(lastLine bool) {
	gccCmd := exec.Command("g++", "-Wall", "testfiles/c++/debug.cpp", "-o", "testfiles/c++/out")
	gccCmd.Run()
	runCmd := exec.Command("./testfiles/c++/out")
	output, err := runCmd.Output()
	if !lastLine {
		fmt.Println(string(output))
	} else {
		outSlice := strings.Split(strings.TrimSpace(string(output)), "\n")
		fmt.Println(outSlice[len(outSlice)-1])
	}
	fmt.Println(err)
}

func Removecomments(s string) string {
	r, _ := regexp.Compile(`/\*(.*?)\*/|//(.*?)\n`)
	ret := r.ReplaceAllString(s, "")
	fmt.Println("TEST\n" + ret + "\nENDTEST\n")
	return ret
}
