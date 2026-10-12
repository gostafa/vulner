// Gostafa 2026.
// SPDX-License-Identifier: Apache-2.0.

// Command vulner reports reachable vulnerabilities in the current Go module.
package main

type (
	mainRuntime = struct {
		run  func([]string) int
		exit func(int)
	}

	mainRuntimeError func() mainRuntime
)
