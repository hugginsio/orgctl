// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package config

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/adrg/xdg"
	"github.com/goccy/go-yaml"
)

//go:embed default.yaml
var defaultConfigFS embed.FS

// Open searches for a configuration file in the following priority order:
// 1. $ORGCTL_CONFIG_HOME/config.yaml
// 2. $XDG_CONFIG_HOME/orgctl/config.yaml
//
// Once a configuration file is found, it calls OpenFromFile with that path.
// Returns an error if no configuration file is found at any of the locations.
func Open() (*Configuration, error) {
	paths := []string{}

	if orgctlConfigHome := os.Getenv("ORGCTL_CONFIG_HOME"); orgctlConfigHome != "" {
		if strings.HasPrefix(orgctlConfigHome, "~") {
			if home, err := os.UserHomeDir(); err == nil {
				paths = append(paths, filepath.Join(home, orgctlConfigHome[1:], "config.yaml"))
			}
		} else {
			paths = append(paths, filepath.Join(orgctlConfigHome, "config.yaml"))
		}
	}

	paths = append(paths, GetDefaultConfigPath())

	// Search for config file in priority order
	for _, path := range paths {
		if _, err := os.Stat(path); err == nil {
			return OpenFromFile(path)
		}
	}

	return nil, fmt.Errorf("no configuration file found in: %v", paths)
}

// OpenFromFile loads the configuration from the specified file path.
// It loads the default configuration and overlays the user configuration on top of it.
// Returns an error if the file cannot be read or parsed.
func OpenFromFile(path string) (*Configuration, error) {
	config, err := loadDefault()
	if err != nil {
		return nil, fmt.Errorf("failed to load default config: %w", err)
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to load config from %s: %w", path, err)
	}

	if err := yaml.Unmarshal(content, config); err != nil {
		return nil, fmt.Errorf("failed to parse config from %s: %w", path, err)
	}

	return postProcess(config)
}

func loadDefault() (*Configuration, error) {
	content, err := GetDefaultConfigContent()
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
	if strings.HasPrefix(cfg.Collection.Path, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, fmt.Errorf("failed to get home directory: %w", err)
		}

		cfg.Collection.Path = filepath.Join(home, cfg.Collection.Path[1:])
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

// GetDefaultConfigPath returns the path where the default config should be written.
// This always returns $XDG_CONFIG_HOME/orgctl/config.yaml
func GetDefaultConfigPath() string {
	return filepath.Join(xdg.ConfigHome, "orgctl", "config.yaml")
}

// GetDefaultConfigContent returns the embedded default.yaml content.
func GetDefaultConfigContent() ([]byte, error) {
	return defaultConfigFS.ReadFile("default.yaml")
}
