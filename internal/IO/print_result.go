package IO

import (
	"fmt"
	"io"
)

func PrintGeneratedUrl(output io.Writer, url string, err error) {
	if err != nil {
		fmt.Fprintf(output, "Error while getting your url: %s", err)
		return
	}

	fmt.Fprintf(output, "Generated url: %s", url)
}
