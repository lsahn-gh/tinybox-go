package main

import (
	"fmt"
	"os"

	"github.com/memnoth/tinybox-go/applets"
)

type funcSym func(args []string) int

// TODO: how to kill the warn?
var appletTbl map[string]funcSym = map[string]funcSym{
	"cat": applets.CatMain,
}

func main() {
	defer panicException()
	if len(os.Args) < 2 {
		panic(nil)
	}

	appletTbl[os.Args[1]](os.Args[2:])
}

func panicException() {
	err := recover()
	if err == nil {
		fmt.Print(
			"usage: tinybox-go <applets>\n\n",
			"- applets\n",
			"    cat\n")
	} else {
		fmt.Println(err)
	}
}
