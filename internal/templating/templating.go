// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

// Package templating provides the template engine used to render documents and filenames.
//
// ## Template Syntax
//
// Templates support all Sprig functions (see https://masterminds.github.io/sprig/),
// plus these custom functions:
//
//   - {{slugify .Title}} - Converts a string to a lowercase URL-friendly slug
//
// The template engine uses Go's text/template syntax and accepts any context type.
// This allows flexible rendering of filenames, filepaths, and content templates.
//
// ## Time Functions
//
// For date/time formatting, use Sprig's built-in functions:
//
//   - {{now | date "2006-01-02"}} - Current date
//   - {{now | date "Mon Jan 2 15:04:05 MST 2006"}} - Full timestamp
package templating

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

// Execute processes a template string with the given context.
// The context can be any type that supports field access (usually a struct).
func Execute(tmplStr string, ctx any) (string, error) {
	funcMap := sprig.FuncMap()

	// Add custom functions
	funcMap["slugify"] = Slugify

	// Create and parse the template
	t, err := template.New("template").Funcs(funcMap).Parse(tmplStr)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	// Execute the template
	var buf bytes.Buffer
	if err := t.Execute(&buf, ctx); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return buf.String(), nil
}

// Returns true if the provided string contains a template expression.
func IsTemplate(str string) bool {
	// NOTE: this is a rough search, consider moving to parsing text/template nodes
	return strings.Contains(str, "{{") && strings.Contains(str, "}}")
}
