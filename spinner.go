package indicators

import (
	"fmt"
	"strings"
	"time"
)

func Spinner(done *bool, delayMs int) {
	delay := time.Millisecond * time.Duration(delayMs)

	spinnerText := `-\|/`

	for {
		for _, r := range spinnerText {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}

		if *done {
			fmt.Printf("\r%s\n", strings.Repeat(" ", len(spinnerText)))
			return
		}
	}
}
