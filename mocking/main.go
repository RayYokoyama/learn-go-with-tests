package mocking

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(writer, strconv.Itoa(i))
	}
	sleeper.Sleep()
	fmt.Fprintf(writer, "Go!")
}

func main() {
	sleeper := ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, &sleeper)
}
