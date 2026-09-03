// Gostafa 2026.
// SPDX-License-Identifier: Apache-2.0.

// Package vulner_test verifies the public plugin API.
package vulner_test

import (
	"testing"

	"github.com/golangci/plugin-module-register/register"
	"github.com/gostafa/vulner"
)

const (
	firstAnalyzerIndex = 0
	singleAnalyzer     = 1
)

// TestNew verifies the plugin builds one configured analyzer.
func TestNew(t *testing.T) {
	t.Parallel()

	plugin, err := vulner.New(nil)
	assertNoPluginError(t, err)

	analyzers, err := plugin.BuildAnalyzers()
	assertNoPluginError(t, err)

	if len(analyzers) != singleAnalyzer || analyzers[firstAnalyzerIndex] == nil {
		t.Fatalf("unexpected analyzers: %#v", analyzers)
	}

	if analyzers[firstAnalyzerIndex].Name != vulner.LinterName {
		t.Fatalf("unexpected analyzer name: %q", analyzers[firstAnalyzerIndex].Name)
	}
}

// TestNewSetsLoadMode verifies the plugin requests syntax-only loading.
func TestNewSetsLoadMode(t *testing.T) {
	t.Parallel()

	plugin, err := vulner.New(nil)
	assertNoPluginError(t, err)

	if plugin.GetLoadMode() != register.LoadModeSyntax {
		t.Fatalf("unexpected load mode: %q", plugin.GetLoadMode())
	}
}

// TestNewRejectsSettings verifies settings are not supported.
func TestNewRejectsSettings(t *testing.T) {
	t.Parallel()

	plugin, err := vulner.New(map[string]any{"packages": "./cmd/..."})
	if err == nil {
		t.Fatalf("expected an error, got %#v", plugin)
	}
}

// TestNewAcceptsEmptySettings verifies golangci-lint's empty configuration is accepted.
func TestNewAcceptsEmptySettings(t *testing.T) {
	t.Parallel()

	plugin, err := vulner.New(map[string]any{})
	assertNoPluginError(t, err)

	if plugin == nil {
		t.Fatal("nil plugin")
	}
}

// TestPluginRegistration verifies the module plugin is registered on import.
func TestPluginRegistration(t *testing.T) {
	t.Parallel()

	constructor, err := register.GetPlugin(vulner.LinterName)
	assertNoPluginError(t, err)

	plugin, err := constructor(nil)
	assertNoPluginError(t, err)

	if plugin == nil {
		t.Fatal("nil plugin")
	}
}

// TestAnalyzerBuildsDefaultAnalyzer verifies the analyzer requires no settings.
func TestAnalyzerBuildsDefaultAnalyzer(t *testing.T) {
	t.Parallel()

	analyzer := vulner.Analyzer()

	if analyzer == nil || analyzer.Name != vulner.LinterName || analyzer.Run == nil {
		t.Fatalf("unexpected analyzer: %#v", analyzer)
	}
}

func assertNoPluginError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatal(err)
	}
}
