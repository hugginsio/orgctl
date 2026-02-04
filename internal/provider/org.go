// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

// Package provider implements document metadata providers.
package provider

type OrgModeProvider struct{}

func (o *OrgModeProvider) Extension() string {
	return ".org"
}
