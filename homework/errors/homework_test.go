package main

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type MultiError struct {
	errs []error
}

func (e *MultiError) Error() string {
	if len(e.errs) == 0 {
		return ""
	}

	var b strings.Builder
	if len(e.errs) == 1 {
		b.WriteString("1 error occured:\n")
	} else {
		fmt.Fprintf(&b, "%d errors occured:\n", len(e.errs))
	}

	for _, err := range e.errs {
		fmt.Fprintf(&b, "\t* %v", err)
	}
	b.WriteString("\n")

	return b.String()
}

func Append(err error, errs ...error) *MultiError {
	var multiErr *MultiError
	if err != nil {
		var ok bool
		multiErr, ok = err.(*MultiError)
		if !ok {
			multiErr = &MultiError{errs: []error{err}}
		}
	} else {
		multiErr = &MultiError{}
	}

	for _, e := range errs {
		if e != nil {
			multiErr.errs = append(multiErr.errs, e)
		}
	}

	return multiErr
}

func TestMultiError(t *testing.T) {
	var err error
	err = Append(err, errors.New("error 1"))
	err = Append(err, errors.New("error 2"))

	expectedMessage := "2 errors occured:\n\t* error 1\t* error 2\n"
	assert.EqualError(t, err, expectedMessage)
}
