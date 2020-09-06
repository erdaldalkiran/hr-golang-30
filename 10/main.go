package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type result struct {
	head       int
	tail       int
	contains   int
	fullOfOnes bool
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	br := strconv.FormatInt(int64(n), 2)

	res := <-process(br)
	m := max(max(res.head, res.tail), res.contains)
	fmt.Println(m)
}

func process(br string) chan result {
	out := make(chan result)

	go func() {
		defer close(out)

		var r result
		l := len(br)
		if l < 3 {
			switch br {
			case "11":
				r = result{2, 2, 2, true}
			case "10":
				r = result{1, 0, 1, false}
			case "01":
				r = result{0, 1, 1, false}
			case "00":
				r = result{0, 0, 0, false}
			case "1":
				r = result{1, 1, 1, true}
			case "0":
				r = result{0, 0, 0, false}
			}

			out <- r
			return
		}

		half := len(br) / 2
		left := br[0:half]
		right := br[half:]

		chs := make([]chan result, 0)
		chs = append(chs, process(left))
		chs = append(chs, process(right))
		var leftResult result
		var rightResult result

		for i, ch := range chs {
			if i == 0 {
				leftResult = <-ch
			} else {
				rightResult = <-ch
			}
		}

		r = result{}
		//head calculation
		r.head = leftResult.head
		if leftResult.fullOfOnes {
			r.head += rightResult.head
		}

		//tail calculation
		r.tail = rightResult.tail
		if rightResult.fullOfOnes {
			r.tail += leftResult.tail
		}

		r.contains = max(max(max(r.head, r.tail), max(leftResult.contains, rightResult.contains)), leftResult.tail+rightResult.head)
		r.fullOfOnes = leftResult.fullOfOnes && rightResult.fullOfOnes

		out <- r
	}()

	return out
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
