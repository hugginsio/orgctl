// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package docid

import (
	"time"
)

// DatetimeGenerator generates a 12-character string representing the current date and time.
type DatetimeGenerator struct{}

func NewDatetimeGenerator() *DatetimeGenerator {
	return &DatetimeGenerator{}
}

func (g *DatetimeGenerator) Generate() (string, error) {
	return time.Now().Format("200601021504"), nil
}
