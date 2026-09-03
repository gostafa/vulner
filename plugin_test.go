// Gostafa 2026.
// SPDX-License-Identifier: Apache-2.0.

package vulner

import (
	"testing"
)

// TestAnalyzerBuildsNamedAnalyzer verifies same-package construction details.
func TestAnalyzerBuildsNamedAnalyzer(t *testing.T) {
	t.Parallel()

	analyzer := Analyzer()

	if analyzer == nil || analyzer.Name != LinterName || analyzer.Run == nil {
		t.Fatalf("unexpected analyzer: %#v", analyzer)
	}
}
