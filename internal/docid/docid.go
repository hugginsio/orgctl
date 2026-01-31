// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

// Package docid provides implementations for generating document IDs.
package docid

import (
	"fmt"
	"strings"
)

// DocIdGenerator generates document IDs.
type DocIdGenerator interface {
	Generate() (string, error)
}

func DetermineGenerator(str string) (*DocIdGenerator, error) {
	var generator DocIdGenerator
	var err error

	switch strings.ToLower(str) {
	case "alphanumeric", "alphanum":
		generator = &AlphanumericGenerator{}
	default:
		err = fmt.Errorf("unrecognized ID generator '%s'", str)
	}

	return &generator, err
}
