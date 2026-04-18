// Package managementasset provides management control panel asset helpers.
// Remote auto-updating has been disabled for security — only locally provided
// assets are served.
package managementasset

import (
"context"
"os"
"path/filepath"
"strings"

"github.com/router-for-me/CLIProxyAPI/v6/internal/config"
"github.com/router-for-me/CLIProxyAPI/v6/internal/util"
log "github.com/sirupsen/logrus"
)

const managementAssetName = "management.html"

// ManagementFileName exposes the control panel asset filename.
const ManagementFileName = managementAssetName

// SetCurrentConfig is a no-op — retained for API compatibility.
func SetCurrentConfig(_ *config.Config) {}

// StartAutoUpdater is intentionally disabled — no remote asset fetching.
func StartAutoUpdater(_ context.Context, _ string) {
log.Info("management asset auto-updater disabled (remote fetching removed)")
}

// StaticDir resolves the directory where the management panel asset lives.
func StaticDir(configFilePath string) string {
if override := strings.TrimSpace(os.Getenv("MANAGEMENT_STATIC_PATH")); override != "" {
cleaned := filepath.Clean(override)
if strings.EqualFold(filepath.Base(cleaned), managementAssetName) {
return filepath.Dir(cleaned)
}
return cleaned
}

if writable := util.WritablePath(); writable != "" {
return filepath.Join(writable, "static")
}

configFilePath = strings.TrimSpace(configFilePath)
if configFilePath == "" {
return ""
}

base := filepath.Dir(configFilePath)
fileInfo, err := os.Stat(configFilePath)
if err == nil && fileInfo.IsDir() {
base = configFilePath
}

return filepath.Join(base, "static")
}

// FilePath resolves the absolute path to the management control panel asset.
func FilePath(configFilePath string) string {
if override := strings.TrimSpace(os.Getenv("MANAGEMENT_STATIC_PATH")); override != "" {
cleaned := filepath.Clean(override)
if strings.EqualFold(filepath.Base(cleaned), managementAssetName) {
return cleaned
}
return filepath.Join(cleaned, ManagementFileName)
}

dir := StaticDir(configFilePath)
if dir == "" {
return ""
}
return filepath.Join(dir, ManagementFileName)
}

// EnsureLatestManagementHTML checks if a local management.html exists on disk.
// Remote fetching is disabled — returns true only if the file already exists locally.
func EnsureLatestManagementHTML(_ context.Context, staticDir string, _ string, _ string) bool {
staticDir = strings.TrimSpace(staticDir)
if staticDir == "" {
return false
}
_, err := os.Stat(filepath.Join(staticDir, managementAssetName))
return err == nil
}
