package main

import (
	"bufio"
	"fmt"
	"os"
	"strings" // Kita akan menggunakan ini untuk fitur debug

	"mylang/internal/lexer"
	"mylang/pkg/token"
)

const PROMPT = ">> "

func main() {
	fmt.Println("Welcome to MyLang programming language!")
	fmt.Println("Feel free to type in commands")
	fmt.Println("Type 'exit' or 'quit' to exit")
	fmt.Println("Type 'debug:<code>' to see tokens")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		if line == "exit" || line == "quit" {
			fmt.Println("Goodbye!")
			return
		}

		// Gunakan strings.HasPrefix untuk fitur debug (ini memperbaiki error)
		if strings.HasPrefix(line, "debug:") {
			debugInput := strings.TrimPrefix(line, "debug:")
			l := lexer.New(debugInput)
			for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
				fmt.Printf("Type: %s, Literal: %s, Line: %d, Column: %d\n",
					   tok.Type, tok.Literal, tok.Line, tok.Column)
			}
			continue
		}

		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("Type: %s, Literal: %s, Line: %d, Column: %d\n",
				   tok.Type, tok.Literal, tok.Line, tok.Column)
		}
	}
}
