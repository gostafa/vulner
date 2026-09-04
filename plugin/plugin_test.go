// Gostafa 2026.
// SPDX-License-Identifier: Apache-2.0.

package plugin_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/golangci/plugin-module-register/register"
	"github.com/gostafa/vulner/plugin"
	"golang.org/x/tools/go/analysis"
)

const (
	firstAnalyzerIndex = 0
	singleAnalyzer     = 1
)

func TestNew(t *testing.T) {
	t.Parallel()

	plug, err := plugin.New(nil)
	assertNoPluginError(t, err)

	analyzers := buildAnalyzersForTest(t, plug)
	assertAnalyzer(t, analyzers)
	assertBuildAnalyzersAgain(t, plug)
}

func TestNewSetsLoadMode(t *testing.T) {
	t.Parallel()

	plug, err := plugin.New(nil)
	assertNoPluginError(t, err)
	assertLoadMode(t, plug)
}

func TestNewRejectsSettings(t *testing.T) {
	t.Parallel()

	plug, err := plugin.New(map[string]any{"packages": "./cmd/..."})
	if err == nil {
		t.Fatalf("expected an error, got %#v", plug)
	}
}

func TestNewRejectsNonMapSettings(t *testing.T) {
	t.Parallel()

	plug, err := plugin.New("not settings")
	if err == nil {
		t.Fatalf("expected an error, got %#v", plug)
	}
}

func TestNewAcceptsEmptySettings(t *testing.T) {
	t.Parallel()

	plug, err := plugin.New(map[string]any{})
	assertNoPluginError(t, err)

	if plug == nil {
		t.Fatal("nil plugin")
	}
}

func TestPluginRegistration(t *testing.T) {
	t.Parallel()

	plug, err := registeredPlugin(t)(nil)
	assertNoPluginError(t, err)

	if plug == nil {
		t.Fatal("nil plugin")
	}
}

func TestRegisteredPluginWrapsSettingsError(t *testing.T) {
	t.Parallel()

	_, err := registeredPlugin(t)(map[string]any{"packages": "./cmd/..."})
	if err == nil || !strings.Contains(err.Error(), "registerVulner") {
		t.Fatalf("error = %v, want registerVulner wrapper", err)
	}
}

func TestAnalyzerBuildsDefaultAnalyzer(t *testing.T) {
	t.Parallel()

	analyzer := plugin.Analyzer()
	if analyzer == nil || analyzer.Name != plugin.Name || analyzer.Run == nil {
		t.Fatalf("unexpected analyzer: %#v", analyzer)
	}
}

func TestBuildAnalyzersWrapsError(t *testing.T) {
	t.Parallel()

	want := errors.New("boom")
	plug := plugin.Plugin(func() ([]*analysis.Analyzer, error) {
		return nil, want
	})

	_, err := plug.BuildAnalyzers()
	if err == nil || !strings.Contains(err.Error(), "BuildAnalyzers") {
		t.Fatalf("error = %v, want BuildAnalyzers wrapper", err)
	}
}

func assertLoadMode(t *testing.T, plug register.LinterPlugin) {
	t.Helper()

	if got := plug.GetLoadMode(); got != register.LoadModeSyntax {
		t.Fatalf("GetLoadMode() = %q, want %q", got, register.LoadModeSyntax)
	}
}

func buildAnalyzersForTest(t *testing.T, plug register.LinterPlugin) []*analysis.Analyzer {
	t.Helper()

	analyzers, err := plug.BuildAnalyzers()
	if err != nil {
		t.Fatalf("BuildAnalyzers: %v", err)
	}

	return analyzers
}

func assertAnalyzer(t *testing.T, analyzers []*analysis.Analyzer) {
	t.Helper()

	if len(analyzers) != singleAnalyzer {
		t.Fatalf("analyzers = %#v, want one analyzer", analyzers)
	}

	if analyzers[firstAnalyzerIndex].Name != plugin.Name {
		t.Fatalf("analyzer name = %q, want %q", analyzers[firstAnalyzerIndex].Name, plugin.Name)
	}

	if analyzers[firstAnalyzerIndex].Run == nil {
		t.Fatal("analyzer Run is nil")
	}
}

func assertBuildAnalyzersAgain(t *testing.T, plug register.LinterPlugin) {
	t.Helper()

	again, err := plug.BuildAnalyzers()
	if err != nil {
		t.Fatalf("second BuildAnalyzers: %v", err)
	}

	if again[firstAnalyzerIndex].Run == nil {
		t.Fatal("analyzer Run is nil")
	}
}

func registeredPlugin(t *testing.T) register.NewPlugin {
	t.Helper()

	constructor, err := register.GetPlugin(plugin.Name)
	if err != nil {
		t.Fatalf("GetPlugin: %v", err)
	}

	return constructor
}

func assertNoPluginError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatal(err)
	}
}
