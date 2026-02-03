// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package util

import (
	"fmt"
	"strings"

	"github.com/hugginsio/orgctl/config"
)

func DetermineGroup(cfg *config.Configuration, args []string) (*config.Collection, error) {
	if len(args) < 1 || args[0] == "" {
		return &cfg.Collection, nil
	}

	// TODO: match on shortest unique prefix as well

	if group, ok := cfg.Group[strings.ToLower(args[0])]; ok {
		return &group, nil
	}

	return nil, fmt.Errorf("Unknown group '%s'", args[0])
}
