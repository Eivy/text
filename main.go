package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) == 1 {
		print(os.Stdin)
	} else {
		for _, p := range os.Args[1:] {
			f, err := os.Open(p)
			if err != nil {
				panic(err)
			}
			print(f)
		}
	}
}

func print(f *os.File) {
	re := regexp.MustCompile("\033\\[([0-9;]+[mABCDEFGJKSTHf]|\\?(25|2004)[hl])")
	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(re.ReplaceAllString(s.Text(), ""))
	}
}
