package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/tomill/lougo"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		out, err := lougo.Lou(scanner.Text())
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			continue
		}

		fmt.Println(out)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}
