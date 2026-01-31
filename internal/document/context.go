// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package document

// Context contains the variables available to templates when rendering documents.
// This is used by the templating engine to substitute values into template strings
// for filenames, filepaths, and content.
type Context struct {
	ID       string // The generated document ID
	Title    string // The document title
	Content  string // The initial content
	Filepath string // The relative path to the document
}
