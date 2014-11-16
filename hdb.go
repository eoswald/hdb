package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	var s string
	var cFile bool = (filepath.Ext(os.Args[1]) == "c")
	lastLine := false
	trackAll := false
	if len(os.Args) > 2 && os.Args[2] == "--lastline=true" {
		lastLine = true
	} else if len(os.Args) > 2 && os.Args[2] == "--trackall=true" {
		trackAll = true
	}

	//MakeFile()
	s = Removecomments(Format(os.Args[1]))
	splitFile := strings.Split(s, "\n")
	splitFile = RemoveNewlines(splitFile)
	lineSlice := CreateLines(splitFile)
	MarkInvalid(lineSlice)
	debugFile, _ := os.Create("testfiles/c++/debug.cpp")

	for i := 0; i < len(lineSlice); i++ {
		AddPrint(debugFile, lineSlice[i], cFile, trackAll)
	}
	debugFile.Close()
	CompileAndRun(lastLine)
}

type lineType struct {
	s    string
	code int
	info string
}

func AddPrint(file *os.File, line lineType, cFile bool, trackAll bool) {
	if cFile {
		if line.code == 0 {
			file.WriteString("printf(\"" + line.s + "\n\");\n")
			file.WriteString(line.s + "\n")
		}
	} else {
		whiteSpaceSize := len(line.s) - len(strings.TrimLeft(line.s, "\t"))
		if whiteSpaceSize == 0 {
			whiteSpaceSize = 4
		}
		whiteSpaceIf := strings.Repeat("\t", whiteSpaceSize+1)
		whiteSpaceLoop := strings.Repeat("\t", whiteSpaceSize)
		switch line.code {
		case 0:
			file.WriteString("std::cout << \"\033[32m" + strings.Replace(line.s, "\"", "\\\"", -1) + "\033[0m\" << std::endl;\n")
			file.WriteString(line.s + "\n")
			exp := strings.Split(line.s, " ")
			if trackAll && len(exp) > 1 && exp[1] == "=" {
				file.WriteString("std::cout << \"" + exp[0] + " = \"" + "<<" + exp[0] + " << std::endl;\n")
			}
		case 343: //function
			file.WriteString(line.s + "\n")
			file.WriteString("std::cout << \"\033[34m" + "Entering " + line.info + "\033[0m\" << std::endl;\n")
		case 666: //if
			file.WriteString("std::cout << \"\033[32m" + strings.Replace(line.s, "\"", "\\\"", -1) + "\033[0m\" << std::endl;\n")
			file.WriteString(line.s + "\n")
			file.WriteString("std::cout << \"\033[34m" + whiteSpaceIf + "if statement: " + line.info + " evaluates to true\033[0m\" << std::endl;\n")
		case 752: //else if
			file.WriteString(line.s + "\n")
			file.WriteString("std::cout << \"\033[34m" + whiteSpaceIf + "else if statement: " + line.info + " evaluates to true\033[0m\" << std::endl;\n")
		case 580: //else
			file.WriteString(line.s + "\n")
			file.WriteString("std::cout << \"\033[34m" +  whiteSpaceIf + "else statement:\033[0m\" << std::endl;\n")
		case 4: //for
			file.WriteString("std::cout << \"\033[32m" + strings.Replace(line.s, "\"", "\\\"", -1) + "\033[0m\" << std::endl;\n")
			file.WriteString(line.s + "\n")

			forArgs := strings.Split(line.info, ";")
			if strings.Contains(forArgs[0], "int ") {
				index := strings.Split(forArgs[0], " ")[1]
				file.WriteString("std::cout << \"\033[34m" + whiteSpaceLoop + "looping: " + "\033[0m\" << " + index + " << std::endl;\n")
			} else {
				file.WriteString("std::cout << \"\033[34m" + whiteSpaceLoop + "looping: " + line.info + "\033[0m\" << std::endl;\n")
			}
		case 603: //while
			file.WriteString("std::cout << \"\033[32m" + strings.Replace(line.s, "\"", "\\\"", -1) + "\033[0m\" << std::endl;\n")
			file.WriteString(line.s + "\n")
			file.WriteString("std::cout << \"\033[34m" + whiteSpaceLoop + line.info + " evaluates to true\033[0m\" << std::endl;\n")
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
			stack = PopStack(stack)
			stack = append(stack, 752)
			lineSlice[i].info = GetParenContents(lineSlice[i].s)
		} else if ifR.MatchString(lineSlice[i].s) {
			//if
			lineSlice[i].code = 666
			stack = append(stack, 666)
			lineSlice[i].info = GetParenContents(lineSlice[i].s)
		} else if forR.MatchString(lineSlice[i].s) {
			//for
			lineSlice[i].code = 4
			stack = append(stack, 4)
			lineSlice[i].info = GetParenContents(lineSlice[i].s)
		} else if whileR.MatchString(lineSlice[i].s) {
			//while
			lineSlice[i].code = 603
			stack = append(stack, 603)
			lineSlice[i].info = GetParenContents(lineSlice[i].s)
		} else if funcR.MatchString(lineSlice[i].s) {
			//func
			lineSlice[i].code = 343
			stack = append(stack, 343)
			drill := strings.SplitN(lineSlice[i].s, "(", 2)[0]
			drillList := strings.Fields(drill)
			lineSlice[i].info = drillList[len(drillList)-1]
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
			stack = PopStack(stack)
			stack = append(stack, 580)
		} else if scopeR.MatchString(lineSlice[i].s) {
			//scope
			lineSlice[i].code = 0
			stack = append(stack, 0)
		} else if closeR.MatchString(lineSlice[i].s) {
			//close
			lineSlice[i].code = 1
			stack = PopStack(stack)
			fmt.Println("asdf")
		} else {
			if InFunction(stack) {
				fmt.Println(stack)
				fmt.Println(lineSlice[i].s)
				lineSlice[i].code = 0
			} else {
				lineSlice[i].code = 1
			}
		}
	}
}

func GetParenContents(line string) string {
	parensR, _ := regexp.Compile(`\((.*)\)`)
	parensR.Longest()
	return parensR.FindString(line)
}

func InFunction(stack []int) (exists bool) {
	for i := 0; i < len(stack); i++ {
		if stack[i] == 343 {
			return true
		}
	}
	return false
}

func PopStack(stack []int) ([]int){
	return stack[:len(stack)-1]
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
		if stringSlice[i] != "" && strings.TrimLeft(stringSlice[i], " ") != "" && strings.TrimLeft(stringSlice[i], " ")[0] != 10 {
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
	ret := r.ReplaceAllString(s, "\n")
	fmt.Println("TEST\n" + ret + "\nENDTEST\n")
	return ret
}

func Format(f string) string {
	var s string
	filename := os.Args[1]
	out, err := exec.Command("uncrustify/src/uncrustify","-c","uncrustify/etc/hdb.cfg","-f",filename).Output()
	if err !=nil {
		panic(err)
	}
	s = string(out)
	return s
}
