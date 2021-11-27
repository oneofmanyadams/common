package prompt

import (
	"fmt"
)

////////////////////////////////////////////////////////////////////////////////
//// Examples
////////////////////////////////////////////////////////////////////////////////

func ExampleOptionPrompt() {
    // Quick little hack for abusing for loops to have user prompts flow
	// with the program code (no need to split user prompting into separate
	// program logic). If the user does not provide an answer that matches an
	// option then the user is prompted again.
	// Users don't type in the option name, they instead enter the number
	// associated with the option (so for example: the first option is
	// selected by the user entering 0). 
	for p := New("Do what?"); p.Ask; p.PromptUser() {
        
		if p.Option("First thing") {
            fmt.Println("One")
        }
        
		if p.Option("Second thing") {
            fmt.Println("Two")
        }

    }
	//Output: Do what?
	// 0 First thing
	// 1 Second thing
	//#:
	//-------------
	//Inputting 0 would output:
	// One
}
