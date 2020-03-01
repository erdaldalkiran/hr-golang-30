package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	lines := make([]string, n, n)

	i := 0
	for scanner.Scan() {
		lines[i] = scanner.Text()
		i++
		if i == n {
			break
		}
	}

	for _, line := range lines {
		e, o := split(line)
		fmt.Printf("%s %s\n", e, o)
	}

}

func split(line string) (even, odd string) {
	var e, o strings.Builder

	for i, rune := range line {
		if i%2 == 0 {
			e.WriteRune(rune)
			continue
		}

		o.WriteRune(rune)
	}

	return e.String(), o.String()
}
