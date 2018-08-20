// Provides a simple front end to the goeval simple evaluator.
package main

import (
	"bufio"
	"fmt"
	"github.com/paulgriffiths/goeval/simple/eval"
	"os"
)

func main() {
	fmt.Printf("Enter simple mathematical expressions, 'q' to quit\n")

	input := bufio.NewScanner(os.Stdin)
	fmt.Printf("> ")
	for input.Scan() {
		if input.Text() == "q" {
			break
		}
		result, err := eval.Evaluate(input.Text())
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			fmt.Printf("  = %v\n", result)
		}

		fmt.Printf("> ")
	}
}
