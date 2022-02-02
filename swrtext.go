package main

import (
	"fmt"
	"strings"
	"time"
)

func SWrText(done *bool, message string, writingSpeedMs, numBlinks, blinkSpeedMs int) {
	writingSpeed := time.Millisecond * time.Duration(writingSpeedMs)
	blinkSpeed := time.Millisecond * time.Duration(blinkSpeedMs)

	for {
		var text string

		for _, char := range message {
			if *done {
				fmt.Printf("\r%s\n", strings.Repeat(" ", len(message)))
				return
			}

			text += string(char)

			fmt.Printf("\r%s", text)

			time.Sleep(writingSpeed)
		}

		time.Sleep(time.Second)

		for i := 0; i < numBlinks; i++ {
			fmt.Printf("\r%s", strings.Repeat(" ", len(message)))
			time.Sleep(blinkSpeed)

			fmt.Printf("\r%s", text)
			time.Sleep(blinkSpeed)
		}

		time.Sleep(time.Second)

		fmt.Printf("\r%s", strings.Repeat(" ", len(message)))
	}
}
