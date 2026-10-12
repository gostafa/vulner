// Gostafa 2026.
// SPDX-License-Identifier: Apache-2.0.

package plugin_test

import (
	"strings"
	"testing"

	"github.com/golangci/plugin-module-register/register"
	"github.com/gostafa/vulner/plugin"
)

func TestPluginRegistration(t *testing.T) {
	t.Parallel()

	plug, err := registeredPlugin(t)(nil)
	if err != nil {
		t.Fatal(err)
	}

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

func registeredPlugin(t *testing.T) register.NewPlugin {
	t.Helper()

	constructor, err := register.GetPlugin(plugin.Name)
	if err != nil {
		t.Fatalf("GetPlugin: %v", err)
	}

	return constructor
}
