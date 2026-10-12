// Gostafa 2026.
// SPDX-License-Identifier: Apache-2.0.

package plugin

import (
	"errors"
	"strings"
	"testing"

	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

const (
	firstAnalyzerIndex = 0
	singleAnalyzer     = 1
)

func TestNew(t *testing.T) {
	t.Parallel()

	plug, err := newPlugin(nil)
	assertNoPluginError(t, err)

	analyzers := buildAnalyzersForTest(t, plug)
	assertAnalyzer(t, analyzers)
	assertBuildAnalyzersAgain(t, plug)
}

func TestNewSetsLoadMode(t *testing.T) {
	t.Parallel()

	plug, err := newPlugin(nil)
	assertNoPluginError(t, err)
	assertLoadMode(t, plug)
}

func TestNewRejectsSettings(t *testing.T) {
	t.Parallel()

	plug, err := newPlugin(map[string]any{"packages": "./cmd/..."})
	if err == nil {
		t.Fatalf("expected an error, got %#v", plug)
	}
}

func TestNewRejectsNonMapSettings(t *testing.T) {
	t.Parallel()

	plug, err := newPlugin("not settings")
	if err == nil {
		t.Fatalf("expected an error, got %#v", plug)
	}
}

func TestNewAcceptsEmptySettings(t *testing.T) {
	t.Parallel()

	plug, err := newPlugin(map[string]any{})
	assertNoPluginError(t, err)

	if plug == nil {
		t.Fatal("nil plugin")
	}
}

func TestAnalyzerBuildsDefaultAnalyzer(t *testing.T) {
	t.Parallel()

	analyzer := newAnalyzer()
	if analyzer == nil || analyzer.Name != Name || analyzer.Run == nil {
		t.Fatalf("unexpected analyzer: %#v", analyzer)
	}
}

func TestBuildAnalyzersWrapsError(t *testing.T) {
	t.Parallel()

	want := errors.New("boom")
	plug := linterPlugin(func() ([]*analysis.Analyzer, error) {
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

	if analyzers[firstAnalyzerIndex].Name != Name {
		t.Fatalf("analyzer name = %q, want %q", analyzers[firstAnalyzerIndex].Name, Name)
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

func assertNoPluginError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatal(err)
	}
}

func TestEmptySettings(t *testing.T) {
	t.Parallel()

	if !emptySettings(nil) {
		t.Fatal("emptySettings(nil) = false")
	}

	if !emptySettings(map[string]any{}) {
		t.Fatal("emptySettings(empty map) = false")
	}

	if emptySettings(map[string]any{"packages": "./..."}) {
		t.Fatal("emptySettings(non-empty map) = true")
	}

	if emptySettings("not settings") {
		t.Fatal("emptySettings(string) = true")
	}
}

func TestAnalyzerBuildsNamedAnalyzer(t *testing.T) {
	t.Parallel()

	analyzer := newAnalyzer()
	if analyzer == nil || analyzer.Name != Name || analyzer.Run == nil {
		t.Fatalf("unexpected analyzer: %#v", analyzer)
	}
}
