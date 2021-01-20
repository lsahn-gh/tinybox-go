package applets

import "os"

// CatMain returns 0 on success
func CatMain(args []string) int {
	const BUFSZ = 1024 * 4
	rdbuf := make([]byte, BUFSZ)

	for _, fname := range args {
		fp, err := os.Open(fname)
		if err != nil {
			panic(err)
		}

		for {
			n, err := fp.Read(rdbuf)
			if err != nil {
				panic(err)
			}
			os.Stdout.Write(rdbuf[:n])

			if n < BUFSZ {
				break
			}
		}

		fp.Close()
	}

	return 0
}
