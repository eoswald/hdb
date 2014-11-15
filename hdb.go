package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, err := ioutil.ReadFile(os.Args[1])
	if err == nil{
		s := string(f)
		fmt.Println(s)
	}
}
