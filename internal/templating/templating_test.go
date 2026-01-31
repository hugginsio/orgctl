// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package templating_test

import (
	"testing"

	"github.com/hugginsio/orgctl/internal/document"
	"github.com/hugginsio/orgctl/internal/templating"
)

// TestExecute tests the Execute function with various templates.
func TestExecute(t *testing.T) {
	ctx := document.Context{
		ID:      "a2b9",
		Title:   "My Awesome Note",
		Content: "Some initial content",
	}

	tests := []struct {
		name        string
		template    string
		ctx         document.Context
		expected    string
		shouldError bool
	}{
		// Basic variable substitution
		{
			name:     "simple ID substitution",
			template: "{{.ID}}",
			ctx:      ctx,
			expected: "a2b9",
		},
		{
			name:     "simple title substitution",
			template: "{{.Title}}",
			ctx:      ctx,
			expected: "My Awesome Note",
		},
		{
			name:     "simple content substitution",
			template: "{{.Content}}",
			ctx:      ctx,
			expected: "Some initial content",
		},
		{
			name:     "slugify function on title",
			template: "{{slugify .Title}}",
			ctx:      ctx,
			expected: "my-awesome-note",
		},
		{
			name:     "combined with ID",
			template: "{{.ID}}-{{slugify .Title}}",
			ctx:      ctx,
			expected: "a2b9-my-awesome-note",
		},
		{
			name:     "org frontmatter with static values",
			template: "#+title: {{.Title}}\n#+id: {{.ID}}\n\n{{.Content}}",
			ctx:      ctx,
			expected: "#+title: My Awesome Note\n#+id: a2b9\n\nSome initial content",
		},
		{
			name:     "org with tags",
			template: "#+title: {{.Title}}\n#+filetags: :note:\n\n{{.Content}}",
			ctx:      ctx,
			expected: "#+title: My Awesome Note\n#+filetags: :note:\n\nSome initial content",
		},
		{
			name:     "empty content",
			template: "#+title: {{.Title}}\n\n{{.Content}}",
			ctx: document.Context{
				ID:      "x1y2",
				Title:   "Empty",
				Content: "",
			},
			expected: "#+title: Empty\n\n",
		},
		{
			name:        "invalid template syntax",
			template:    "{{.Title}",
			ctx:         ctx,
			shouldError: true,
		},
		{
			name:        "undefined variable",
			template:    "{{.UndefinedField}}",
			ctx:         ctx,
			shouldError: true, // Go templates error on undefined struct fields
		},
		{
			name:        "invalid function",
			template:    "{{invalidFunc .Title}}",
			ctx:         ctx,
			shouldError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := templating.Execute(tt.template, tt.ctx)

			if tt.shouldError && err == nil {
				t.Errorf("Execute() expected error but got nil")
			}
			if !tt.shouldError && err != nil {
				t.Errorf("Execute() unexpected error: %v", err)
			}

			if !tt.shouldError && result != tt.expected {
				t.Errorf("Execute() = %q, expected %q", result, tt.expected)
			}
		})
	}
}

func BenchmarkExecute(b *testing.B) {
	tmpl := "{{.ID}}-{{slugify .Title}}.org"
	ctx := document.Context{
		ID:    "a2b9",
		Title: "My Awesome Note",
	}

	b.ReportAllocs()
	for b.Loop() {
		templating.Execute(tmpl, ctx)
	}
}
