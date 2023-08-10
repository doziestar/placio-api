//go:build docker
// +build docker

package errors_test

import (
	"fmt"
	"placio-pkg/errors"
)

func ExampleNew() {
	err := errors.New("example")

	fmt.Printf("%s\n", err.Error())
	var e *errors.AppError
	if errors.As(err, &e) {
		fmt.Printf("%s", e.StackTrace())
	}

	// Output:
	// example
}

func ExampleWrap() {
	subErr := errors.New("example")
	err := errors.Wrap(subErr)

	fmt.Printf("%s\n", err.Error())
	var e *errors.AppError
	if errors.As(err, &e) {
		fmt.Printf("%s", e.StackTrace())
	}

	// Output:
	// example
}

func ExampleWrap_second() {
	originalErr := fmt.Errorf("original")
	wrappedErr := errors.Wrap(originalErr)

	err := fmt.Errorf("test2: %w", wrappedErr)

	fmt.Printf("%s\n", err.Error())
	var e *errors.AppError
	if errors.As(err, &e) {
		fmt.Printf("%s", e.StackTrace())
	}

	// Output:
	// test2: original
}

func ExampleWrap_third() {
	first := errors.Wrap(fmt.Errorf("first"))
	wrapped := errors.Wrap(first)

	second := errors.Wrap(fmt.Errorf("second: %w", wrapped))
	third := errors.Wrap(fmt.Errorf("third: %w", second))
	err := errors.Wrap(third)

	fmt.Printf("%s\n", err.Error())
	var e *errors.AppError
	if errors.As(err, &e) {
		fmt.Printf("%s", e.StackTrace())
	}

	// Output:
	// third: second: first
}
