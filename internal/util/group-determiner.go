// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package util

import (
	"fmt"
	"strings"

	"github.com/hugginsio/orgctl/config"
)

func DetermineGroup(cfg *config.Configuration, name string) (*config.Group, error) {
	if name == "" {
		return &cfg.Group, nil
	}

	// TODO: match on shortest unique prefix as well

	candidate := strings.ToLower(name)
	if subgroup, ok := cfg.Subgroup[candidate]; ok {
		return &subgroup, nil
	}

	return nil, fmt.Errorf("Unknown subgroup '%s'", name)
}
