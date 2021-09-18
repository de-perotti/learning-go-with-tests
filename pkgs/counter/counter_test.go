package counter_test

import (
	"bytes"
	"fmt"
	"github.com/de-perotti/learning-go-with-tests/pkgs/counter"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type UltraSonicSpy struct {
	Calls  []string
	Buffer *bytes.Buffer
}

const write = "write"
const sleep = "sleep"

func (spy *UltraSonicSpy) Sleep(duration time.Duration) {
	spy.Calls = append(spy.Calls, sleep)
}

func (spy *UltraSonicSpy) Write(b []byte) (n int, err error) {
	spy.Calls = append(spy.Calls, write)
	return fmt.Fprint(spy.Buffer, string(b))
}

func TestCounter(t *testing.T) {
	ultraSpy := &UltraSonicSpy{
		Calls:  []string{},
		Buffer: &bytes.Buffer{},
	}

	err := counter.Counter(ultraSpy, ultraSpy)
	assert.NoError(t, err)

	expected := `3
2
1
Go!`
	assert.Equal(t, expected, ultraSpy.Buffer.String())
	assert.Equal(t, []string{
		write,
		sleep,
		write,
		sleep,
		write,
		sleep,
		write,
	}, ultraSpy.Calls)
}
