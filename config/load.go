// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package config

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/goccy/go-yaml"
)

//go:embed default.yaml
var defaultConfigFS embed.FS

func Load(path string) (*Configuration, error) {
	config, err := loadDefault()
	if err != nil {
		return nil, fmt.Errorf("failed to load default config: %w", err)
	}

	if path != "" {
		if _, err := os.Stat(path); err == nil {
			content, err := os.ReadFile(path)
			if err != nil {
				return nil, fmt.Errorf("failed to load config from %s: %w", path, err)
			}

			// Unmarshal user config on top of defaults, if it exists
			if err := yaml.Unmarshal(content, config); err != nil {
				return nil, fmt.Errorf("failed to parse config from %s: %w", path, err)
			}
		}
	}

	return postProcess(config)
}

func loadDefault() (*Configuration, error) {
	content, err := defaultConfigFS.ReadFile("default.yaml")
	if err != nil {
		return nil, err
	}

	var config Configuration
	if err := yaml.Unmarshal(content, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func postProcess(cfg *Configuration) (*Configuration, error) {
	// Expand `~` to user home directory
	if strings.HasPrefix(cfg.Group.Path, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, fmt.Errorf("failed to get home directory: %w", err)
		}

		cfg.Group.Path = filepath.Join(home, cfg.Group.Path[1:])
	}

	// Determine default editor
	if cfg.Tools.Editor == "" {
		editor := os.Getenv("EDITOR")
		if editor != "" {
			cfg.Tools.Editor = editor
		} else {
			return nil, fmt.Errorf("editor not set")
		}
	}

	return cfg, nil
}
