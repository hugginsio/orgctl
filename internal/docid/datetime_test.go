// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package docid_test

import (
	"testing"

	"github.com/hugginsio/orgctl/internal/docid"
)

func TestGenerate(t *testing.T) {
	gen := docid.NewDatetimeGenerator()

	t.Run("validate datetime shape", func(t *testing.T) {
		str, err := gen.Generate()

		if err != nil {
			t.Errorf("Generate() unexpected error: %v", err)
		}

		if len(str) != 12 {
			t.Errorf("Expected length to be 12, got: %v", len(str))
		}

		// TODO: validate based on current datetime
	})
}

func BenchmarkDatetimeGenerator_Generate(b *testing.B) {
	gen := docid.NewDatetimeGenerator()

	b.ReportAllocs()
	for b.Loop() {
		_, err := gen.Generate()
		if err != nil {
			b.Fatalf("Generate() returned error: %v", err)
		}
	}
}
