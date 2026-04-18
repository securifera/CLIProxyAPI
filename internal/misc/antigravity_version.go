// Package misc provides miscellaneous utility functions for the CLI Proxy API server.
package misc

import (
"context"
"fmt"

log "github.com/sirupsen/logrus"
)

const antigravityFallbackVersion = "1.21.9"

// StartAntigravityVersionUpdater is intentionally disabled — no remote version fetching.
// The hardcoded fallback version is used for all User-Agent strings.
func StartAntigravityVersionUpdater(_ context.Context) {
log.Info("antigravity version updater disabled (remote fetching removed), using static version " + antigravityFallbackVersion)
}

// AntigravityLatestVersion returns the static fallback antigravity version.
func AntigravityLatestVersion() string {
return antigravityFallbackVersion
}

// AntigravityUserAgent returns the User-Agent string for antigravity requests.
func AntigravityUserAgent() string {
return fmt.Sprintf("antigravity/%s darwin/arm64", antigravityFallbackVersion)
}
