// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

// Package provider implements document metadata providers.
package provider

import (
	"fmt"
	"strings"
)

type DocumentMetadataProvider interface {
	Extension() string // Returns the supported file extension.
}

func DetermineProvider(str string) (*DocumentMetadataProvider, error) {
	var provider DocumentMetadataProvider
	var err error

	switch strings.ToLower(str) {
	case "org":
		provider = &OrgModeProvider{}
	default:
		err = fmt.Errorf("unrecognized document provider '%s'", str)
	}

	return &provider, err
}
