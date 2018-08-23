// Provides a simple front end to the goeval simple evaluator.
package main

import (
	"bufio"
	"fmt"
	"github.com/paulgriffiths/goeval/simple/vareval"
	"os"
)

func main() {
	fmt.Printf("Enter simple mathematical expressions, or use keyword ")
	fmt.Printf("'let' to set variables.\nEnter 'q' to quit\n")

	input := bufio.NewScanner(os.Stdin)
	fmt.Printf("> ")
	env := vareval.NewStdEnv()
	for input.Scan() {
		if input.Text() == "q" {
			break
		}
		result, err := vareval.ParseStatement(input.Text())
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			if err := result.Execute(env); err != nil {
				fmt.Printf("error: %v\n", err)
			}
		}

		fmt.Printf("> ")
	}
}
