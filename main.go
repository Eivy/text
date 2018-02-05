package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
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
	re := regexp.MustCompile("\033\\[(([0-9]+;?)+([mABCDEFGJKSTHf])|\\?25[hl]|\\?2004[hl])")
	s := bufio.NewScanner(f)
	var t []string
	for s.Scan() {
		if re.MatchString(s.Text()) {
			m := re.Split(s.Text(), -1)
			p := re.FindAllStringSubmatch(s.Text(), -1)
			var tmp []string
			for i, sub := range p {
				switch sub[len(sub)-1] {
				case "J":
					n, err := strconv.Atoi(sub[len(sub)-2])
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
					if n > 0 {
						t = nil
						tmp = nil
					} else {
						tmp = append(tmp, m[i])
					}
				case "K":
					n, err := strconv.Atoi(sub[len(sub)-2])
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
					if n > 0 {
						tmp = nil
					} else {
						tmp = append(tmp, m[i])
					}
				default:
					tmp = append(tmp, m[i])
				}
			}
			tmp = append(tmp, m[len(m)-1])
			t = append(t, strings.Join(tmp, ""))
		}
	}
	fmt.Print(strings.Join(t, "\n"))
}
