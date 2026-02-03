// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package config

// Configuration for a collection of org-mode documents.
type Configuration struct {
	Collection `yaml:",inline"`      // Group defines a single collection of org-mode documents.
	Group      map[string]Collection `yaml:"group"` // Groups are named collections of org-mode documents within the greater collection.
	Tools      Tools                 `yaml:"tools"` // Tools configures the companion editor.
}

// Collection defines a single group of org-mode documents and their configuration.
type Collection struct {
	Path             string `yaml:"path"`              // The path to the containing directory for the Collection.
	DefaultTitle     string `yaml:"default-title"`     // The default title used if none is provided.
	FilenameTemplate string `yaml:"filename-template"` // The template used to generate the document filename.
	ContentTemplate  string `yaml:"content-template"`  // The path to the template file used to generate the initial content. TODO: only if path use path
}

type Tools struct {
	Editor string `yaml:"editor"` // The editor to launch when opening documents.
}
