// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package docid

import (
	"crypto/rand"
	"fmt"
)

const (
	// crockfordAlphabet contains the 32 characters from Crockford's Base32 spec
	// Excludes: I, L, O, U to avoid confusion
	crockfordAlphabet = "0123456789abcdefghjkmnpqrstvwxyz"
	idLength          = 4
)

// AlphanumericGenerator generates short document IDs using characters from Crockford's Base32.
type AlphanumericGenerator struct{}

func NewAlphanumericGenerator() *AlphanumericGenerator {
	return &AlphanumericGenerator{}
}

// Generate creates a random 4-character lowercase string using Crockford's Base32 alphabet.
// Returns an error if the random number generator fails.
func (g *AlphanumericGenerator) Generate() (string, error) {
	// Generate 4 random bytes
	randomBytes := make([]byte, idLength)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %w", err)
	}

	// Map each random byte to a character in the alphabet
	result := make([]byte, idLength)
	for i := range idLength {
		result[i] = crockfordAlphabet[randomBytes[i]%32]
	}

	return string(result), nil
}
