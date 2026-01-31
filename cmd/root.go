// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	goversion "github.com/caarlos0/go-version"
	"github.com/charmbracelet/fang"
	"github.com/hugginsio/orgctl/config"
	"github.com/spf13/cobra"
)

var cfg *config.Configuration

var rootCmd = &cobra.Command{
	Use:   "orgctl",
	Short: "High velocity note taking with org-mode",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		path := filepath.Join(xdg.ConfigHome, "orgctl", "config.yaml")

		loadedCfg, err := config.Load(path)
		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}

		cfg = loadedCfg
		return nil
	},
}

func Execute() {
	if err := fang.Execute(
		context.Background(),
		rootCmd,
		fang.WithCommit(goversion.GetVersionInfo().GitCommit),
		fang.WithVersion(goversion.GetVersionInfo().GitVersion),
		fang.WithoutCompletions(),
		fang.WithoutManpage(),
	); err != nil {
		os.Exit(1)
	}
}
