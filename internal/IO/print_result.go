package IO

import (
	"fmt"
	"io"
)

const (
	errorMessagePrefix   = "Error while getting your url: "
	successMessagePrefix = "Generated url: "
)

func PrintGeneratedUrl(output io.Writer, url string, err error) {
	if err != nil {
		fmt.Fprintf(output, "%s%s\n", errorMessagePrefix, err)
		return
	}

	fmt.Fprintf(output, "%s%s\n", successMessagePrefix, url)
}
