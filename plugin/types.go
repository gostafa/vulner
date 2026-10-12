// Gostafa 2026.
// SPDX-License-Identifier: Apache-2.0.

// Package plugin adapts the vulner analyzer to golangci-lint.
package plugin

import (
	"golang.org/x/tools/go/analysis"
)

type (
	// linterPlugin is the golangci-lint module plugin for vulner.
	linterPlugin func() ([]*analysis.Analyzer, error)

	loadMode struct{}
)
