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
	phoneBook := make(map[string]string, n)

	i := 0
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		phoneBook[words[0]] = words[1]
		i++
		if i == n {
			break
		}
	}

	for scanner.Scan() {
		name := scanner.Text()
		p, found := phoneBook[name]
		if !found {
			fmt.Println("Not found")
			continue
		}

		fmt.Printf("%s=%s\n", name, p)
	}
}
