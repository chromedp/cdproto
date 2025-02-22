// Package systeminfo provides the Chrome DevTools Protocol
// commands, types, and events for the SystemInfo domain.
//
// The SystemInfo domain defines methods and events for querying low-level
// system information.
//
// Generated by the cdproto-gen command.
package systeminfo

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"

	"github.com/chromedp/cdproto/cdp"
)

// GetInfoParams returns information about the system.
type GetInfoParams struct{}

// GetInfo returns information about the system.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/SystemInfo#method-getInfo
func GetInfo() *GetInfoParams {
	return &GetInfoParams{}
}

// GetInfoReturns return values.
type GetInfoReturns struct {
	GPU          *GPUInfo `json:"gpu,omitempty,omitzero"`          // Information about the GPUs on the system.
	ModelName    string   `json:"modelName,omitempty,omitzero"`    // A platform-dependent description of the model of the machine. On Mac OS, this is, for example, 'MacBookPro'. Will be the empty string if not supported.
	ModelVersion string   `json:"modelVersion,omitempty,omitzero"` // A platform-dependent description of the version of the machine. On Mac OS, this is, for example, '10.1'. Will be the empty string if not supported.
	CommandLine  string   `json:"commandLine,omitempty,omitzero"`  // The command line string used to launch the browser. Will be the empty string if not supported.
}

// Do executes SystemInfo.getInfo against the provided context.
//
// returns:
//
//	gpu - Information about the GPUs on the system.
//	modelName - A platform-dependent description of the model of the machine. On Mac OS, this is, for example, 'MacBookPro'. Will be the empty string if not supported.
//	modelVersion - A platform-dependent description of the version of the machine. On Mac OS, this is, for example, '10.1'. Will be the empty string if not supported.
//	commandLine - The command line string used to launch the browser. Will be the empty string if not supported.
func (p *GetInfoParams) Do(ctx context.Context) (gpu *GPUInfo, modelName string, modelVersion string, commandLine string, err error) {
	// execute
	var res GetInfoReturns
	err = cdp.Execute(ctx, CommandGetInfo, nil, &res)
	if err != nil {
		return nil, "", "", "", err
	}

	return res.GPU, res.ModelName, res.ModelVersion, res.CommandLine, nil
}

// GetFeatureStateParams returns information about the feature state.
type GetFeatureStateParams struct {
	FeatureState string `json:"featureState"`
}

// GetFeatureState returns information about the feature state.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/SystemInfo#method-getFeatureState
//
// parameters:
//
//	featureState
func GetFeatureState(featureState string) *GetFeatureStateParams {
	return &GetFeatureStateParams{
		FeatureState: featureState,
	}
}

// GetFeatureStateReturns return values.
type GetFeatureStateReturns struct {
	FeatureEnabled bool `json:"featureEnabled,omitempty,omitzero"`
}

// Do executes SystemInfo.getFeatureState against the provided context.
//
// returns:
//
//	featureEnabled
func (p *GetFeatureStateParams) Do(ctx context.Context) (featureEnabled bool, err error) {
	// execute
	var res GetFeatureStateReturns
	err = cdp.Execute(ctx, CommandGetFeatureState, p, &res)
	if err != nil {
		return false, err
	}

	return res.FeatureEnabled, nil
}

// GetProcessInfoParams returns information about all running processes.
type GetProcessInfoParams struct{}

// GetProcessInfo returns information about all running processes.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/SystemInfo#method-getProcessInfo
func GetProcessInfo() *GetProcessInfoParams {
	return &GetProcessInfoParams{}
}

// GetProcessInfoReturns return values.
type GetProcessInfoReturns struct {
	ProcessInfo []*ProcessInfo `json:"processInfo,omitempty,omitzero"` // An array of process info blocks.
}

// Do executes SystemInfo.getProcessInfo against the provided context.
//
// returns:
//
//	processInfo - An array of process info blocks.
func (p *GetProcessInfoParams) Do(ctx context.Context) (processInfo []*ProcessInfo, err error) {
	// execute
	var res GetProcessInfoReturns
	err = cdp.Execute(ctx, CommandGetProcessInfo, nil, &res)
	if err != nil {
		return nil, err
	}

	return res.ProcessInfo, nil
}

// Command names.
const (
	CommandGetInfo         = "SystemInfo.getInfo"
	CommandGetFeatureState = "SystemInfo.getFeatureState"
	CommandGetProcessInfo  = "SystemInfo.getProcessInfo"
)
