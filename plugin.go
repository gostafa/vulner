// Gostafa 2026.
// SPDX-License-Identifier: Apache-2.0.

// Package vulner exposes the golangci-lint vulnerability analyzer module plugin.
package vulner

import (
	"errors"

	"github.com/golangci/plugin-module-register/register"
	"github.com/gostafa/vulner/internal/vulnerability/adapters/filesystem"
	golangciadapter "github.com/gostafa/vulner/internal/vulnerability/adapters/golangci"
	govulnadapter "github.com/gostafa/vulner/internal/vulnerability/adapters/govuln"
	"github.com/gostafa/vulner/internal/vulnerability/application"
	"golang.org/x/tools/go/analysis"
)

type (
	// Plugin adapts vulner to golangci-lint's module plugin interface.
	Plugin struct{}
)

const (
	// LinterName is the registered golangci-lint analyzer name.
	LinterName = "vulner"

	unsupportedSettingsError = "creating vulner module plugin: settings are not supported"
)

var errUnsupportedSettings = errors.New(unsupportedSettingsError)

//nolint:gochecknoinits // GolangCI-Lint module plugins must register themselves during package initialization.
func init() {
	register.Plugin(LinterName, New)
}

// New constructs the module plugin. Vulner does not accept settings.
//
//nolint:ireturn // GolangCI-Lint's module plugin constructor contract returns register.LinterPlugin.
func New(rawSettings any) (register.LinterPlugin, error) {
	if !emptySettings(rawSettings) {
		return nil, errUnsupportedSettings
	}

	return &Plugin{}, nil
}

// BuildAnalyzers returns the configured analyzer for golangci-lint.
func (*Plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{Analyzer()}, nil
}

// GetLoadMode returns the golangci-lint package loading mode.
func (*Plugin) GetLoadMode() string {
	return register.LoadModeSyntax
}

// Analyzer constructs the default vulnerability analyzer.
func Analyzer() *analysis.Analyzer {
	locator := filesystem.NewModuleLocator()
	scanner := govulnadapter.NewScanner(application.DefaultPackages)
	service := application.NewService(locator, scanner)

	return golangciadapter.Build(golangciadapter.NewAnalyzer(LinterName, service))
}

//nolint:revive // An empty map is the only supported plugin settings value.
func emptySettings(rawSettings any) bool {
	if rawSettings == nil {
		return true
	}

	settings, found := rawSettings.(map[string]any)

	return found && len(settings) == 0
}
