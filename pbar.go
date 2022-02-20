package indicators

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
	"time"
)

const (
	prBarFinishingSpeedMs = 20
)

const (
	prBarImitationMinSpeedMs       = 50
	prBarImitationMaxSpeedMs       = 1500
	prBarImitationFinishingSpeedMs = 10
)

func PrBar(done *bool, firstChar, nextChar string, speedMs, length int) {
	waiting := time.Millisecond * time.Duration(speedMs)

	for {
		for i := 1; i <= length; i++ {
			fmt.Printf("\r[%s%s]",
				strings.Repeat(firstChar, i), strings.Repeat(nextChar, length-i))

			time.Sleep(waiting)

			if *done {
				for i := i; i <= length; i++ {
					fmt.Printf("\r[%s%s][ok]",
						strings.Repeat(firstChar, i), strings.Repeat(nextChar, length-i))

					time.Sleep(time.Millisecond * prBarFinishingSpeedMs)
				}

				fmt.Printf("\n")
				return
			}
		}
	}
}

func PrBarImitation(done *bool, firstChar, nextChar string, length int) {
	waiting := time.Millisecond * randNum()

	var count = 0

	for {
		count++

		for i := 1; i <= length; i++ {
			fmt.Printf("\r[%s%s][%d]",
				strings.Repeat(firstChar, i), strings.Repeat(nextChar, length-i), count)

			if i%3 == 0 {
				waiting = time.Millisecond * randNum()
			}

			time.Sleep(waiting)

			if *done {
				for i := i; i <= length; i++ {
					fmt.Printf("\r[%s%s][ok]",
						strings.Repeat(firstChar, i), strings.Repeat(nextChar, length-i))

					time.Sleep(time.Millisecond * prBarImitationFinishingSpeedMs)
				}

				fmt.Printf("\n")
				return
			}
		}
	}
}

func randNum() time.Duration {
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(prBarImitationMaxSpeedMs)))
	return time.Duration(prBarImitationMinSpeedMs + int(n.Int64()))
}
