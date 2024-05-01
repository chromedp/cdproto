// Package pwa provides the Chrome DevTools Protocol
// commands, types, and events for the PWA domain.
//
// This domain allows interacting with the browser to control PWAs.
//
// Generated by the cdproto-gen command.
package pwa

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"

	"github.com/chromedp/cdproto/cdp"
)

// GetOsAppStateParams returns the following OS state for the given manifest
// id.
type GetOsAppStateParams struct {
	ManifestID string `json:"manifestId"` // The id from the webapp's manifest file, commonly it's the url of the site installing the webapp. See https://web.dev/learn/pwa/web-app-manifest.
}

// GetOsAppState returns the following OS state for the given manifest id.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/PWA#method-getOsAppState
//
// parameters:
//
//	manifestID - The id from the webapp's manifest file, commonly it's the url of the site installing the webapp. See https://web.dev/learn/pwa/web-app-manifest.
func GetOsAppState(manifestID string) *GetOsAppStateParams {
	return &GetOsAppStateParams{
		ManifestID: manifestID,
	}
}

// GetOsAppStateReturns return values.
type GetOsAppStateReturns struct {
	BadgeCount   int64          `json:"badgeCount,omitempty"`
	FileHandlers []*FileHandler `json:"fileHandlers,omitempty"`
}

// Do executes PWA.getOsAppState against the provided context.
//
// returns:
//
//	badgeCount
//	fileHandlers
func (p *GetOsAppStateParams) Do(ctx context.Context) (badgeCount int64, fileHandlers []*FileHandler, err error) {
	// execute
	var res GetOsAppStateReturns
	err = cdp.Execute(ctx, CommandGetOsAppState, p, &res)
	if err != nil {
		return 0, nil, err
	}

	return res.BadgeCount, res.FileHandlers, nil
}

// InstallParams installs the given manifest identity, optionally using the
// given install_url or IWA bundle location. TODO(crbug.com/337872319) Support
// IWA to meet the following specific requirement. IWA-specific install
// description: If the manifest_id is isolated-app://, install_url_or_bundle_url
// is required, and can be either an http(s) URL or file:// URL pointing to a
// signed web bundle (.swbn). The .swbn file's signing key must correspond to
// manifest_id. If Chrome is not in IWA dev mode, the installation will fail,
// regardless of the state of the allowlist.
type InstallParams struct {
	ManifestID            string `json:"manifestId"`
	InstallURLOrBundleURL string `json:"installUrlOrBundleUrl,omitempty"` // The location of the app or bundle overriding the one derived from the manifestId.
}

// Install installs the given manifest identity, optionally using the given
// install_url or IWA bundle location. TODO(crbug.com/337872319) Support IWA to
// meet the following specific requirement. IWA-specific install description: If
// the manifest_id is isolated-app://, install_url_or_bundle_url is required,
// and can be either an http(s) URL or file:// URL pointing to a signed web
// bundle (.swbn). The .swbn file's signing key must correspond to manifest_id.
// If Chrome is not in IWA dev mode, the installation will fail, regardless of
// the state of the allowlist.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/PWA#method-install
//
// parameters:
//
//	manifestID
func Install(manifestID string) *InstallParams {
	return &InstallParams{
		ManifestID: manifestID,
	}
}

// WithInstallURLOrBundleURL the location of the app or bundle overriding the
// one derived from the manifestId.
func (p InstallParams) WithInstallURLOrBundleURL(installURLOrBundleURL string) *InstallParams {
	p.InstallURLOrBundleURL = installURLOrBundleURL
	return &p
}

// Do executes PWA.install against the provided context.
func (p *InstallParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandInstall, p, nil)
}

// UninstallParams uninstals the given manifest_id and closes any opened app
// windows.
type UninstallParams struct {
	ManifestID string `json:"manifestId"`
}

// Uninstall uninstals the given manifest_id and closes any opened app
// windows.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/PWA#method-uninstall
//
// parameters:
//
//	manifestID
func Uninstall(manifestID string) *UninstallParams {
	return &UninstallParams{
		ManifestID: manifestID,
	}
}

// Do executes PWA.uninstall against the provided context.
func (p *UninstallParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandUninstall, p, nil)
}

// Command names.
const (
	CommandGetOsAppState = "PWA.getOsAppState"
	CommandInstall       = "PWA.install"
	CommandUninstall     = "PWA.uninstall"
)
