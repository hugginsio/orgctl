// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package editor

import (
	"os"
	"os/exec"
)

func EditorCapture(editor string, path string) error {
	cmd := exec.Command(editor, path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
