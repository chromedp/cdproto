package headlessexperimental

import (
	"fmt"
	"strings"
)

// Code generated by cdproto-gen. DO NOT EDIT.

// ScreenshotParams encoding options for a screenshot.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental#type-ScreenshotParams
type ScreenshotParams struct {
	Format           ScreenshotParamsFormat `json:"format,omitempty,omitzero"`  // Image compression format (defaults to png).
	Quality          int64                  `json:"quality,omitempty,omitzero"` // Compression quality from range [0..100] (jpeg and webp only).
	OptimizeForSpeed bool                   `json:"optimizeForSpeed"`           // Optimize image encoding for speed, not for resulting size (defaults to false)
}

// ScreenshotParamsFormat image compression format (defaults to png).
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental#type-ScreenshotParams
type ScreenshotParamsFormat string

// String returns the ScreenshotParamsFormat as string value.
func (t ScreenshotParamsFormat) String() string {
	return string(t)
}

// ScreenshotParamsFormat values.
const (
	ScreenshotParamsFormatJpeg ScreenshotParamsFormat = "jpeg"
	ScreenshotParamsFormatPng  ScreenshotParamsFormat = "png"
	ScreenshotParamsFormatWebp ScreenshotParamsFormat = "webp"
)

// UnmarshalJSON satisfies [json.Unmarshaler].
func (t *ScreenshotParamsFormat) UnmarshalJSON(buf []byte) error {
	s := string(buf)
	s = strings.TrimSuffix(strings.TrimPrefix(s, `"`), `"`)

	switch ScreenshotParamsFormat(s) {
	case ScreenshotParamsFormatJpeg:
		*t = ScreenshotParamsFormatJpeg
	case ScreenshotParamsFormatPng:
		*t = ScreenshotParamsFormatPng
	case ScreenshotParamsFormatWebp:
		*t = ScreenshotParamsFormatWebp
	default:
		return fmt.Errorf("unknown ScreenshotParamsFormat value: %v", s)
	}
	return nil
}
