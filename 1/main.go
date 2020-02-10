package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var i uint64 = 4
var d float64 = 4.0
var s string = "HackerRank "

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	var ri uint64
	ri, err := strconv.ParseUint(scanner.Text(), 10, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	scanner.Scan()
	rd, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	scanner.Scan()
	rs := scanner.Text()

	fmt.Printf("%d\n%.1f\n%s\n", i+ri, d+rd, s+rs)

}
