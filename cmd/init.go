// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/charmbracelet/huh"
	"github.com/hugginsio/orgctl/config"
	"github.com/spf13/cobra"
)

var (
	dryRun     bool
	forceWrite bool
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize with default configuration",
	RunE: func(cmd *cobra.Command, args []string) error {
		content, err := config.GetDefaultConfigContent()
		if err != nil {
			return fmt.Errorf("failed to load default config: %w", err)
		}

		if dryRun {
			fmt.Println(string(content))
			return nil
		}

		configPath := config.GetDefaultConfigPath()

		if _, err := os.Stat(configPath); err == nil {
			if !forceWrite {
				return fmt.Errorf("config file already exists at %s. Use --force to overwrite", configPath)
			}
		}

		var confirm bool
		if !forceWrite {
			err := huh.NewConfirm().
				Title(fmt.Sprintf("Write default configuration to %s?", configPath)).
				Affirmative("Yes").
				Negative("No").
				Value(&confirm).
				Run()

			if err != nil {
				return fmt.Errorf("Prompt failed: %w", err)
			}

			if !confirm {
				fmt.Println("No changes were made.")
				return nil
			}
		}

		configDir := filepath.Dir(configPath)
		if err := os.MkdirAll(configDir, 0755); err != nil {
			return fmt.Errorf("Failed to create config directory: %w", err)
		}

		if err := os.WriteFile(configPath, content, 0644); err != nil {
			return fmt.Errorf("Failed to write config file: %w", err)
		}

		fmt.Printf("Configuration written to %s\n", configPath)
		return nil
	},
}

func init() {
	initCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Print the default configuration without writing")
	initCmd.Flags().BoolVar(&forceWrite, "force", false, "Bypass prompt and write configuration")

	rootCmd.AddCommand(initCmd)
}
