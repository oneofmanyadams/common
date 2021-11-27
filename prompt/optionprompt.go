package prompt

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// OptionPrompt is used for creating a defined set of options to present
// to a user. This should not be used directly, but instead created by calling 
// the New() function.
type OptionPrompt struct {
	Question string
	Options []string
	Answer string
	Ask bool

	Input io.Reader
	Output io.Writer

	UserInput string
}

////////////////////////////////////////////////////////////////////////////////
// Constructor
////////////////////////////////////////////////////////////////////////////////

// New should be the only way for creating an OptionPrompt instance.
func New(question string) (op OptionPrompt) {
	op.Question = question
	op.Ask = true
	op.Input = os.Stdin
	op.Output = os.Stdout
	return
}

////////////////////////////////////////////////////////////////////////////////
// Public Functions
////////////////////////////////////////////////////////////////////////////////

// Option does several things:
// 1. Validates the answer provided by the user matches this specific option.
// This match is currently only done based on the option key. This does
// does not actually check the option_name string.
// 2. Adds option_name as a valid option if it doesn't already exist.
func (s *OptionPrompt) Option(option_name string) bool {
	key, option_exists := s.optionExists(option_name);
	if option_exists == false{
		key = s.addOption(option_name)
	}
	
	// Convert key to string to compare to user input.
	if s.UserInput == strconv.Itoa(key) {
		s.Answer = option_name
		s.Ask = false
		return true
	}
	return false
}

// PromptUser takes the originally provided question, adds all options onto the
// end of the question, send that string to s.Output, then reads what the user
// provided through s.Input.
// This function will only prompt the user if s.Ask is set to true. if s.Ask
// is false then the assumption is that the user already provided a valid answer
// so we don't need to keep asking.
func (s *OptionPrompt) PromptUser() {
	// Only proceed if we still need input from the user.
	if s.Ask == false {
		return
	}

	// Build question string based on originally provided question with all
	// options appended to the end.
	// This uses each option's array key as the key presented to the
	// user as well.
	question := s.Question
	for option_k, option_v := range s.Options {
		question = fmt.Sprintf("%s\n %d %s", question, option_k, option_v)
	}

	// Prompt user for input.
	user_input, err := ask(question, s.Input, s.Output)
	if err != nil {
		log.Fatal(err)
	}
	s.UserInput = user_input
}

////////////////////////////////////////////////////////////////////////////////
// Private Functions
////////////////////////////////////////////////////////////////////////////////

// optionExists determines if the provided option_name already
// exists in s.Options.
// If the option does exist it's corresponding key is returned,
// as well as the boolean value true.
// If the option does not exist then the next available key is returned,
// as well as the boolean value false.
func (s *OptionPrompt) optionExists(option_name string) (key int, exists bool) {
	for k, v := range s.Options {
		if option_name == v {
			return k, true
		}
	}
	return len(s.Options), false
}

// addOption adds an option to s.Options and returns it's key. 
func (s *OptionPrompt) addOption(option_name string) (key int){
	key = len(s.Options)
	s.Options = append(s.Options, option_name)
	return key
}