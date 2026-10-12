// Gostafa 2026.
// SPDX-License-Identifier: Apache-2.0.

package vulner

import (
	"io"

	"github.com/gostafa/vulner/internal/vulnerability/standalone"
)

// Run scans the current Go module and writes diagnostics to stdout and stderr.
// It accepts no arguments and returns 0 for success, 1 for failure, or 2 for usage errors.
func Run(arguments []string, stdout, stderr io.Writer) int {
	return standalone.Run(arguments, stdout, stderr)
}
