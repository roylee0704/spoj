// package main solves: TEST - Life, the Universe, and Everything
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var token int
	for scanner.Scan() {
		token, _ = strconv.Atoi(scanner.Text())
		if token == 41 {
			break
		}
		fmt.Println(token)
	}
}
