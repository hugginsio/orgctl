// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package cmd

import (
	"fmt"
	"strings"

	"github.com/hugginsio/orgctl/internal/util"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new [subgroup]",
	Short: "Create a new document",
	Args:  cobra.RangeArgs(0, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		group := cfg.Group
		if len(args) > 0 {
			newGroup, err := util.DetermineGroup(cfg, strings.ToLower(args[0]))
			if err != nil {
				return err
			}

			group = *newGroup
		}

		fmt.Println(group)

		return nil
	},
}

func init() {
	newCmd.Flags().StringP(FlagTitle, "t", "", "Title of the new document")

	rootCmd.AddCommand(newCmd)
}
