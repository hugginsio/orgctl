// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package cmd

import (
	"context"
	"fmt"
	"os"

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
		if cmd.Name() == "init" {
			return nil
		}

		loadedCfg, err := config.Open()
		if err != nil {
			return fmt.Errorf("Failed to load config: %w. Try running orgctl init", err)
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
