// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package docid_test

import (
	"testing"

	"github.com/hugginsio/orgctl/internal/docid"
)

func BenchmarkAlphanumericGenerator_Generate(b *testing.B) {
	gen := docid.NewAlphanumericGenerator()

	b.ReportAllocs()
	for b.Loop() {
		_, err := gen.Generate()
		if err != nil {
			b.Fatalf("Generate() returned error: %v", err)
		}
	}
}
