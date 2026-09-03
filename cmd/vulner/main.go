// Gostafa 2026.
// SPDX-License-Identifier: Apache-2.0.

// Command vulner reports reachable vulnerabilities in the current Go module.
package main

import (
	"os"

	"github.com/gostafa/vulner/internal/vulnerability/standalone"
)

func main() {
	os.Exit(standalone.Run(os.Args[1:], os.Stdout, os.Stderr))
}
