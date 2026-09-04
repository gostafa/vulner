// Gostafa 2026.
// SPDX-License-Identifier: Apache-2.0.

package plugin

import (
	"errors"

	"github.com/golangci/plugin-module-register/register"
)

var (
	errUnsupportedSettings                       = errors.New(unsupportedSettingsError)
	_                      register.LinterPlugin = Plugin(nil)
	_                                            = registerVulner()
)
