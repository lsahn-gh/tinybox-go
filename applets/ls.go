package applets

import (
	"flag"
	"io/ioutil"
	"os"
)

var flagObj *flag.FlagSet
var optShowHidden *bool

// LsMain returns 0 on success
func LsMain(args []string) int {
	// TODO should be implemented as a common util
	flagObj = flag.NewFlagSet("ls", flag.ContinueOnError)
	optShowHidden = flagObj.Bool("a", false, "Show hidden files")
	flagObj.Parse(args)

	// adjust args's position
	args = args[flagObj.NFlag():]
	if len(args) == 0 {
		args = append(args, ".")
	}

	var list []os.FileInfo

	for _, fname := range args {
		fi, err := os.Stat(fname)
		if err != nil {
			continue
		}

		if isHiddenType(fi) && !*optShowHidden {
			continue
		}
		list = append(list, fi)
	}

	for _, entry := range list {
		if entry.IsDir() {
			scanDirectory(entry)
		} else {
			scanFile(entry)
		}
	}

	return 0
}

func isHiddenType(fi os.FileInfo) bool {
	return (fi.Name()[0] == '.' && len(fi.Name()) > 1)
}

func scanDirectory(fi os.FileInfo) {
	farr, err := ioutil.ReadDir(fi.Name())
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return
	}
	for _, f := range farr {
		if isHiddenType(f) && !*optShowHidden {
			continue
		}
		println(f.Name())
	}
}

func scanFile(fi os.FileInfo) {
	println(fi.Name())
}
