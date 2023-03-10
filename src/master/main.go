package main

import (
	"fmt"
	"github.com/chushi0/qqbot-go/src/master/module"
)

func main() {
	fmt.Println("Hello world")
	err := (&module.ModuleManager{}).LoadModule("echo")
	if err != nil {
		panic(err)
	}
}
