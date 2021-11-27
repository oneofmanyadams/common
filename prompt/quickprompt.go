package prompt

import (
	"os"
)

// QuickPrompt is designed as a 1 line solution to get user input from the
// command line. 
func QuickPrompt(question string) (answer string, err error) {
	return ask(question, os.Stdin, os.Stdout)
}