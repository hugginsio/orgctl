// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package templating

import (
	"strings"

	"github.com/gosimple/slug"
)

// Slugify converts a string to a lowercase URL-friendly slug.
// This functionality is provided by github.com/gosimple/slug.
func Slugify(s string) string {
	return strings.ToLower(slug.Make(s))
}
