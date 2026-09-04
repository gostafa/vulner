// Gostafa 2026.
// SPDX-License-Identifier: Apache-2.0.

package plugin

const (
	// Name is the registered golangci-lint analyzer name.
	Name = "vulner"

	registerDone             = 99
	zeroValue                = 0
	errBuildAnalyzers        = "BuildAnalyzers: %w"
	unsupportedSettingsError = "creating vulner module plugin: settings are not supported"
)
