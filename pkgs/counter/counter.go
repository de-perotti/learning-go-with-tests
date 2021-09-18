package counter

import (
	"fmt"
	"io"
	"time"
)

type Sleeper interface {
	Sleep(duration time.Duration)
}

type CounterSleeper struct {
}

func (c *CounterSleeper) Sleep(duration time.Duration) {
	time.Sleep(duration)
}

func Counter(w io.Writer, sleeper Sleeper) (err error) {
	i := 4

	for {
		i -= 1

		if i < 0 {
			break
		}

		if i == 0 {
			_, err = fmt.Fprint(w, "Go!")
		}

		if i > 0 {
			_, err = fmt.Fprintln(w, i)
			sleeper.Sleep(1 * time.Second)
		}

		if err != nil {
			return err
		}
	}

	return nil
}
