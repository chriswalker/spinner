// Package spinner is a simple implementation of a text-based 'spinner' UI
// widget, for use when programs need to indicate they are busy. It takes
// inspiration from Brian Down's spinner implementation (the erase() method
// in particular, and embedding an io.Writer for testing), and the simple
// spinner implementation in 'The Go Programming Lanuguage'
package spinner

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Spinner holds all the required options for the spinner instance;
// users supply a prefix if required (otherwise the spinner will just
// display as-is)
type Spinner struct {
	Prefix   string
	Writer   io.Writer
	doneChan chan struct{}
}

// NewSpinner returns an initialised Spinner structure; callers will
// need to supply a prefix to the spinner if required.
// The default behaviour outputs to stdout; this can be overridden
func NewSpinner() Spinner {
	return Spinner{
		Writer:   os.Stdout,
		doneChan: make(chan struct{}, 1),
	}
}

// Start initiates the spinner
func (s *Spinner) Start() {
	go func() {
		for {
			for _, r := range `-\|/` {
				select {
				case <-s.doneChan:
					return
				default:
					fmt.Fprintf(s.Writer, "\r%s%c", s.Prefix, r)
					time.Sleep(100 * time.Millisecond)
				}
			}
		}
	}()
}

// Stop stops the spinner, and erases all outputted characters
func (s *Spinner) Stop() {
	s.erase()
	s.doneChan <- struct{}{}
}

// erase overwrites the prefix plus any spinner chars with space
// characters; when done, the cursor position will return to column
// 1 on the current line
func (s *Spinner) erase() {
	l := len(s.Prefix) + 1
	for i := 0; i < l; i++ {
		for _, c := range []string{"\b", " ", "\b"} {
			fmt.Printf(c)
		}
	}
}
