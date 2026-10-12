// Gostafa 2026.
// SPDX-License-Identifier: Apache-2.0.

package vulner_test

import (
	"bytes"
	"testing"

	"github.com/gostafa/vulner"
)

func TestRunRejectsArguments(t *testing.T) {
	t.Parallel()

	var stdout, stderr bytes.Buffer
	code := vulner.Run([]string{"./..."}, &stdout, &stderr)
	if code != 2 || stdout.Len() != 0 || stderr.String() != "usage: vulner\n" {
		t.Fatalf("Run = %d, stdout = %q, stderr = %q", code, stdout.String(), stderr.String())
	}
}
