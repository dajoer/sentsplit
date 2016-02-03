package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
	"strings"
	"regexp"
)

func ReadSentences(r io.Reader) ([]string) {
	scanner := bufio.NewScanner(r)
	splitre := regexp.MustCompile("([.] |[?] |! |: )")
	var out []string
	var tmp string
	for scanner.Scan() {
		tmp = scanner.Text()
		if tmp != "" {
			out = append(out, splitre.Split(tmp, -1)...)
		}
	}
	return out
}

func CleanSentences(s []string) ([]string) {
	re := regexp.MustCompile("[^-a-z0-9_ äöüß]")
	for i,_ := range s {
		s[i] = strings.ToLower(s[i])
		s[i] = re.ReplaceAllString(s[i], "")
	}
	return s
}

func ReadCleanSentences(r io.Reader) ([]string) {
	return CleanSentences(ReadSentences(r))
}

func main() {
	sentences := ReadCleanSentences(os.Stdin)
	for _,a := range sentences {
		fmt.Println(a)
	}
}
