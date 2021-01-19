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
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("usage : tinybox-go <applet>")
		}
	}()
	if len(os.Args) < 2 {
		panic("Oops")
	}

	appletTbl[os.Args[1]](os.Args[2:])
}
