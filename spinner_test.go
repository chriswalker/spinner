package spinner_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/chriswalker/spinner"
)

// Test implementation of io.Writer; gets passed into
// the spinner so we can test written values
type TestWriter struct {
	output []string
}

// Write just adds the supplied data to the output slice
// fro later inspection
func (tw *TestWriter) Write(data []byte) (int, error) {
	tw.output = append(tw.output, string(data))
	return len(data), nil
}

func TestSpinner(t *testing.T) {
	expected := []string{"\r-", "\r\\", "\r|", "\r/"}

	s := spinner.NewSpinner()
	testWriter := &TestWriter{}
	s.Writer = testWriter

	s.Start()
	time.Sleep(350 * time.Millisecond)
	s.Stop()

	for i, r := range expected {
		if r != testWriter.output[i] {
			t.Errorf("unexpected output; want '%s', got '%s'\n", r, testWriter.output[i])
		}
	}
}

func TestSpinnerWithPrefix(t *testing.T) {
	prefix := "test prefix "
	expected := []string{
		fmt.Sprintf("\r%s-", prefix),
		fmt.Sprintf("\r%s\\", prefix),
		fmt.Sprintf("\r%s|", prefix),
		fmt.Sprintf("\r%s/", prefix),
	}

	s := spinner.NewSpinner()
	s.Prefix = prefix
	testWriter := &TestWriter{}
	s.Writer = testWriter

	s.Start()
	time.Sleep(350 * time.Millisecond)
	s.Stop()

	for i, r := range expected {
		if r != testWriter.output[i] {
			t.Errorf("unexpected output; want '%s', got '%s'\n", r, testWriter.output[i])
		}
	}

}
