package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// loop through all files
	for _, fileName := range os.Args[1:] {
		// 1) read the file
		fileHandle, err := os.Open(fileName)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		var lc, wc, cc int

		scan := bufio.NewScanner(fileHandle)

		for scan.Scan() {
			line := scan.Text()
			wc += len(strings.Fields(line))
			cc += len(line)
			lc++
		}

		fmt.Printf(" %7d %7d %7d  %s \n", lc, wc, cc, fileName)
	}

}
