package cast

// Code generated by cdproto-gen. DO NOT EDIT.

// Sink [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Cast#type-Sink
type Sink struct {
	Name    string `json:"name"`
	ID      string `json:"id"`
	Session string `json:"session,omitempty,omitzero"` // Text describing the current session. Present only if there is an active session on the sink.
}
