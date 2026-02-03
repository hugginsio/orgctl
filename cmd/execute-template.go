// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/hugginsio/orgctl/internal/templating"
	"github.com/spf13/cobra"
)

var executeTemplateCmd = &cobra.Command{
	Use:   "execute-template [template]",
	Short: "Execute a text template",
	Long:  "Execute a text template. Useful for testing templates or programmatically invoking orgctl. If no template is specified, the template is read from stdin.",
	Example: `
orgctl execute-template '{{ "hello!" | upper | repeat 5 }}'
echo '{{ nospace "hello w o r l d" }}' | orgctl execute-template`,
	Args: cobra.RangeArgs(0, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var tmpl string
		if len(args) == 0 {
			incoming, err := io.ReadAll(os.Stdin)
			if err != nil {
				return err
			}

			tmpl = string(incoming)
		} else {
			tmpl = args[0]
		}

		str, err := templating.Execute(tmpl, nil)
		if err != nil {
			return err
		}

		fmt.Println(str)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(executeTemplateCmd)
}
