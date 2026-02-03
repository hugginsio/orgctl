// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hugginsio/orgctl/internal/docid"
	"github.com/hugginsio/orgctl/internal/document"
	"github.com/hugginsio/orgctl/internal/editor"
	"github.com/hugginsio/orgctl/internal/templating"
	"github.com/hugginsio/orgctl/internal/util"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new [group]",
	Short: "Create a new document",
	Args:  cobra.RangeArgs(0, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		group := cfg.Collection
		newGroup, err := util.DetermineGroup(cfg, args)
		if err != nil {
			return err
		}

		group = *newGroup
		parentDir := group.Path

		if !filepath.IsAbs(parentDir) {
			parentDir = filepath.Join(cfg.Path, group.Path)
		}

		id, err := cmd.Flags().GetString(FlagId)
		if err != nil {
			return err
		}

		if id == "" {
			if generatedId, err := docid.NewAlphanumericGenerator().Generate(); err != nil {
				return err
			} else {
				id = generatedId
			}
		}

		title, err := cmd.Flags().GetString(FlagTitle)
		if err != nil {
			return err
		}

		if title == "" {
			title = group.DefaultTitle
		}

		docCtx := &document.Context{
			ID:       id,
			Title:    title,
			Content:  group.ContentTemplate,
			Filepath: group.FilenameTemplate,
		}

		if templating.IsTemplate(docCtx.Title) {
			if docCtx.Title, err = templating.Execute(docCtx.Title, docCtx); err != nil {
				return err
			}
		}

		// TODO: pull content template from path
		// TODO: support input flag

		if templating.IsTemplate(docCtx.Content) {
			if docCtx.Content, err = templating.Execute(docCtx.Content, docCtx); err != nil {
				return err
			}
		}

		if templating.IsTemplate(docCtx.Filepath) {
			if newFilepath, err := templating.Execute(docCtx.Filepath, docCtx); err != nil {
				return err
			} else {
				docCtx.Filepath = filepath.Join(parentDir, newFilepath)
			}
		}

		if filepath.Ext(docCtx.Filepath) == "" {
			docCtx.Filepath += ".org"
		}

		// TODO: need to check if file already exists

		if docCtx.Content == "" {
			if err := editor.EditorCapture(cfg.Tools.Editor, docCtx.Filepath); err != nil {
				return err
			}
		} else {
			tmp := filepath.Join(filepath.Dir(docCtx.Filepath), docCtx.ID+".org")

			if err := os.WriteFile(tmp, []byte(docCtx.Content), 0600); err != nil {
				return fmt.Errorf("Failed to write temporary file: %s", err)
			}

			if err := editor.EditorCapture(cfg.Tools.Editor, tmp); err != nil {
				return err
			}

			// TODO: improvement: check if file unchanged from content template, delete if so
			if err := os.Rename(tmp, docCtx.Filepath); err != nil && !errors.Is(err, os.ErrNotExist) {
				return err
			}
		}

		return nil
	},
}

func init() {
	newCmd.Flags().String(FlagId, "", "Override document ID generation")
	newCmd.Flags().StringP(FlagTitle, "t", "", "Title of the new document")

	rootCmd.AddCommand(newCmd)
}
