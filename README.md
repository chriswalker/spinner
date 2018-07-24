# Spinner
[![CircleCI](https://circleci.com/gh/chriswalker/spinner/tree/master.svg?style=svg)](https://circleci.com/gh/chriswalker/spinner/tree/master)

A simple spinner for CLI UIs, building upon the spinner outlined in 'The Go Programming Language' book, and some ideas from Brian Downs' [Spinner](http://github.com/briandowns/spinner), mainly around testing and erasing spinner output.

It's a very minimalist implementation with a single spinner 'style', and an API consisting of `Start()` and `Stop()` methods:


```golang
spinner := spinner.NewSpinner()
// Set optional prefix text; if none specified, only the
// spinner is displayed
spinner.Prefix = "Getting things "
spinner.Start()
// Do some intensive stuff
spinner.Stop()
```

The spinner structure contains an `io.Writer`; in the default implementation this is just `os.Stdout`, but this can be replaced with something else if required (as is the case for the tests).
