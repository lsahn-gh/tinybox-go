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
		panic("usage")
	}

	appletTbl[os.Args[1]](os.Args[2:])
}

func panicException() {
	err := recover()
	switch err {
	case nil: // skip
	default:
		fmt.Println(err)
	case "usage":
		fmt.Print(
			"usage: tinybox-go <applets>\n\n",
			"- applets\n",
			"    cat\n")
	}
}
