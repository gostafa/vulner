// Gostafa 2026.
// SPDX-License-Identifier: Apache-2.0.

package plugin

import (
	"fmt"

	"github.com/golangci/plugin-module-register/register"
	"github.com/gostafa/vulner/internal/vulnerability/adapters/filesystem"
	golangciadapter "github.com/gostafa/vulner/internal/vulnerability/adapters/golangci"
	govulnadapter "github.com/gostafa/vulner/internal/vulnerability/adapters/govuln"
	"github.com/gostafa/vulner/internal/vulnerability/application"
	"golang.org/x/tools/go/analysis"
)

func registerVulner() int {
	register.Plugin(Name, func(raw any) (register.LinterPlugin, error) {
		pluginInstance, err := New(raw)
		if err != nil {
			return nil, fmt.Errorf("registerVulner: %w", err)
		}

		return pluginInstance, nil
	})

	return registerDone
}

// New constructs the module plugin. Vulner does not accept settings.
func New(raw any) (Plugin, error) {
	if !emptySettings(raw) {
		return nil, fmt.Errorf("New: %w", errUnsupportedSettings)
	}

	return Plugin(analyzerBuilder()), nil
}

func analyzerBuilder() func() ([]*analysis.Analyzer, error) {
	return func() ([]*analysis.Analyzer, error) {
		return []*analysis.Analyzer{Analyzer()}, nil
	}
}

func emptySettings(rawSettings any) bool {
	if rawSettings == nil {
		return true
	}

	return emptySettingsMap(rawSettings)
}

func emptySettingsMap(rawSettings any) bool {
	settings, found := rawSettings.(map[string]any)

	return found && len(settings) == zeroValue
}

// Analyzer constructs the default vulnerability analyzer.
func Analyzer() *analysis.Analyzer {
	locator := filesystem.NewModuleLocator()
	scanner := govulnadapter.NewScanner(application.DefaultPackages)
	service := application.NewService(locator, scanner)

	return golangciadapter.Build(golangciadapter.NewAnalyzer(Name, service))
}

// BuildAnalyzers returns the configured analyzer for golangci-lint.
func (build Plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	analyzers, err := build()
	if err != nil {
		return nil, fmt.Errorf(errBuildAnalyzers, err)
	}

	return analyzers, nil
}

func (loadMode) value() string {
	return register.LoadModeSyntax
}

// GetLoadMode returns the golangci-lint package loading mode.
func (Plugin) GetLoadMode() string {
	return loadMode{}.value()
}
