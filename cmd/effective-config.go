// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var effectiveConfigCmd = &cobra.Command{
	Use:    "effective-config",
	Hidden: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		bytes, err := json.MarshalIndent(cfg, "", "  ")
		if err != nil {
			return err
		}

		fmt.Println(string(bytes))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(effectiveConfigCmd)
}
