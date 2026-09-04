// Gostafa 2026.
// SPDX-License-Identifier: Apache-2.0.

package plugin

import (
	"testing"
)

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

	analyzer := Analyzer()
	if analyzer == nil || analyzer.Name != Name || analyzer.Run == nil {
		t.Fatalf("unexpected analyzer: %#v", analyzer)
	}
}
