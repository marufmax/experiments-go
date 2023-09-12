package dependancy_injection

import (
	"bytes"
	"fmt"
)

func Greet(writter *bytes.Buffer, name string) {
	fmt.Fprintf(writter, "Hello, %s", name)
}
