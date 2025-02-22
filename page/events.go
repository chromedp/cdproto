package page

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/runtime"
)

// EventDomContentEventFired [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#event-domContentEventFired
type EventDomContentEventFired struct {
	Timestamp *cdp.MonotonicTime `json:"timestamp"`
}

// EventFileChooserOpened emitted only when page.interceptFileChooser is
// enabled.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#event-fileChooserOpened
type EventFileChooserOpened struct {
	FrameID       cdp.FrameID           `json:"frameId"`                          // Id of the frame containing input node.
	Mode          FileChooserOpenedMode `json:"mode"`                             // Input mode.
	BackendNodeID cdp.BackendNodeID     `json:"backendNodeId,omitempty,omitzero"` // Input node id. Only present for file choosers opened via an <input type="file"> element.
}

// EventFrameAttached fired when frame has been attached to its parent.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#event-frameAttached
type EventFrameAttached struct {
	FrameID       cdp.FrameID         `json:"frameId"`                  // Id of the frame that has been attached.
	ParentFrameID cdp.FrameID         `json:"parentFrameId"`            // Parent frame identifier.
	Stack         *runtime.StackTrace `json:"stack,omitempty,omitzero"` // JavaScript stack trace of when frame was attached, only set if frame initiated from script.
}

// EventFrameDetached fired when frame has been detached from its parent.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#event-frameDetached
type EventFrameDetached struct {
	FrameID cdp.FrameID         `json:"frameId"` // Id of the frame that has been detached.
	Reason  FrameDetachedReason `json:"reason"`
}

// EventFrameSubtreeWillBeDetached fired before frame subtree is detached.
// Emitted before any frame of the subtree is actually detached.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#event-frameSubtreeWillBeDetached
type EventFrameSubtreeWillBeDetached struct {
	FrameID cdp.FrameID `json:"frameId"` // Id of the frame that is the root of the subtree that will be detached.
}

// EventFrameNavigated fired once navigation of the frame has completed.
// Frame is now associated with the new loader.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#event-frameNavigated
type EventFrameNavigated struct {
	Frame *cdp.Frame     `json:"frame"` // Frame object.
	Type  NavigationType `json:"type"`
}

// EventDocumentOpened fired when opening document to write to.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#event-documentOpened
type EventDocumentOpened struct {
	Frame *cdp.Frame `json:"frame"` // Frame object.
}

// EventFrameResized [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#event-frameResized
type EventFrameResized struct{}

// EventFrameStartedNavigating fired when a navigation starts. This event is
// fired for both renderer-initiated and browser-initiated navigations. For
// renderer-initiated navigations, the event is fired after
// frameRequestedNavigation. Navigation may still be cancelled after the event
// is issued. Multiple events can be fired for a single navigation, for example,
// when a same-document navigation becomes a cross-document navigation (such as
// in the case of a frameset).
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#event-frameStartedNavigating
type EventFrameStartedNavigating struct {
	FrameID        cdp.FrameID                          `json:"frameId"`  // ID of the frame that is being navigated.
	URL            string                               `json:"url"`      // The URL the navigation started with. The final URL can be different.
	LoaderID       cdp.LoaderID                         `json:"loaderId"` // Loader identifier. Even though it is present in case of same-document navigation, the previously committed loaderId would not change unless the navigation changes from a same-document to a cross-document navigation.
	NavigationType FrameStartedNavigatingNavigationType `json:"navigationType"`
}

// EventFrameRequestedNavigation fired when a renderer-initiated navigation
// is requested. Navigation may still be cancelled after the event is issued.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#event-frameRequestedNavigation
type EventFrameRequestedNavigation struct {
	FrameID     cdp.FrameID                 `json:"frameId"`     // Id of the frame that is being navigated.
	Reason      ClientNavigationReason      `json:"reason"`      // The reason for the navigation.
	URL         string                      `json:"url"`         // The destination URL for the requested navigation.
	Disposition ClientNavigationDisposition `json:"disposition"` // The disposition for the navigation.
}

// EventFrameStartedLoading fired when frame has started loading.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#event-frameStartedLoading
type EventFrameStartedLoading struct {
	FrameID cdp.FrameID `json:"frameId"` // Id of the frame that has started loading.
}

// EventFrameStoppedLoading fired when frame has stopped loading.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#event-frameStoppedLoading
type EventFrameStoppedLoading struct {
	FrameID cdp.FrameID `json:"frameId"` // Id of the frame that has stopped loading.
}

// EventInterstitialHidden fired when interstitial page was hidden.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#event-interstitialHidden
type EventInterstitialHidden struct{}

// EventInterstitialShown fired when interstitial page was shown.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#event-interstitialShown
type EventInterstitialShown struct{}

// EventJavascriptDialogClosed fired when a JavaScript initiated dialog
// (alert, confirm, prompt, or onbeforeunload) has been closed.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#event-javascriptDialogClosed
type EventJavascriptDialogClosed struct {
	Result    bool   `json:"result"`    // Whether dialog was confirmed.
	UserInput string `json:"userInput"` // User input in case of prompt.
}

// EventJavascriptDialogOpening fired when a JavaScript initiated dialog
// (alert, confirm, prompt, or onbeforeunload) is about to open.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#event-javascriptDialogOpening
type EventJavascriptDialogOpening struct {
	URL               string     `json:"url"`                              // Frame url.
	Message           string     `json:"message"`                          // Message that will be displayed by the dialog.
	Type              DialogType `json:"type"`                             // Dialog type.
	HasBrowserHandler bool       `json:"hasBrowserHandler"`                // True iff browser is capable showing or acting on the given dialog. When browser has no dialog handler for given target, calling alert while Page domain is engaged will stall the page execution. Execution can be resumed via calling Page.handleJavaScriptDialog.
	DefaultPrompt     string     `json:"defaultPrompt,omitempty,omitzero"` // Default dialog prompt.
}

// EventLifecycleEvent fired for lifecycle events (navigation, load, paint,
// etc) in the current target (including local frames).
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#event-lifecycleEvent
type EventLifecycleEvent struct {
	FrameID   cdp.FrameID        `json:"frameId"`  // Id of the frame.
	LoaderID  cdp.LoaderID       `json:"loaderId"` // Loader identifier. Empty string if the request is fetched from worker.
	Name      string             `json:"name"`
	Timestamp *cdp.MonotonicTime `json:"timestamp"`
}

// EventBackForwardCacheNotUsed fired for failed bfcache history navigations
// if BackForwardCache feature is enabled. Do not assume any ordering with the
// Page.frameNavigated event. This event is fired only for main-frame history
// navigation where the document changes (non-same-document navigations), when
// bfcache navigation fails.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#event-backForwardCacheNotUsed
type EventBackForwardCacheNotUsed struct {
	LoaderID                    cdp.LoaderID                                `json:"loaderId"`                                       // The loader id for the associated navigation.
	FrameID                     cdp.FrameID                                 `json:"frameId"`                                        // The frame id of the associated frame.
	NotRestoredExplanations     []*BackForwardCacheNotRestoredExplanation   `json:"notRestoredExplanations"`                        // Array of reasons why the page could not be cached. This must not be empty.
	NotRestoredExplanationsTree *BackForwardCacheNotRestoredExplanationTree `json:"notRestoredExplanationsTree,omitempty,omitzero"` // Tree structure of reasons why the page could not be cached for each frame.
}

// EventLoadEventFired [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#event-loadEventFired
type EventLoadEventFired struct {
	Timestamp *cdp.MonotonicTime `json:"timestamp"`
}

// EventNavigatedWithinDocument fired when same-document navigation happens,
// e.g. due to history API usage or anchor navigation.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#event-navigatedWithinDocument
type EventNavigatedWithinDocument struct {
	FrameID        cdp.FrameID                           `json:"frameId"`        // Id of the frame.
	URL            string                                `json:"url"`            // Frame's new url.
	NavigationType NavigatedWithinDocumentNavigationType `json:"navigationType"` // Navigation type
}

// EventScreencastFrame compressed image data requested by the
// startScreencast.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#event-screencastFrame
type EventScreencastFrame struct {
	Data      string                   `json:"data"`      // Base64-encoded compressed image.
	Metadata  *ScreencastFrameMetadata `json:"metadata"`  // Screencast frame metadata.
	SessionID int64                    `json:"sessionId"` // Frame number.
}

// EventScreencastVisibilityChanged fired when the page with currently
// enabled screencast was shown or hidden .
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#event-screencastVisibilityChanged
type EventScreencastVisibilityChanged struct {
	Visible bool `json:"visible"` // True if the page is visible.
}

// EventWindowOpen fired when a new window is going to be opened, via
// window.open(), link click, form submission, etc.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#event-windowOpen
type EventWindowOpen struct {
	URL            string   `json:"url"`            // The URL for the new window.
	WindowName     string   `json:"windowName"`     // Window name.
	WindowFeatures []string `json:"windowFeatures"` // An array of enabled window features.
	UserGesture    bool     `json:"userGesture"`    // Whether or not it was triggered by user gesture.
}

// EventCompilationCacheProduced issued for every compilation cache
// generated. Is only available if Page.setGenerateCompilationCache is enabled.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Page#event-compilationCacheProduced
type EventCompilationCacheProduced struct {
	URL  string `json:"url"`
	Data string `json:"data"` // Base64-encoded data
}
