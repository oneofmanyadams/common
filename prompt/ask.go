package prompt

import (
	"bufio"
	"io"
	"strings"
)

// ask is the base funtion that performs the output of question q to the user
// and returns the answer a provided by the user.
// If an error occurs, a is returned as empty string.
func ask(q string, input io.Reader, output io.Writer) (a string, err error) {
	rdr := bufio.NewReader(input)

	output.Write([]byte(q+"\n"))
	output.Write([]byte("#: ")) //+"\n" <---- Do I need this?

	raw_answer, read_error := rdr.ReadString('\n')

	if read_error != nil {
		return "", read_error
	}

	cleanup_input := strings.NewReplacer("\n", "")
	a = strings.TrimSpace(cleanup_input.Replace(raw_answer))

	return
}