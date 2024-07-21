package emulation

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"fmt"

	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// ScreenOrientation screen orientation.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#type-ScreenOrientation
type ScreenOrientation struct {
	Type  OrientationType `json:"type"`  // Orientation type.
	Angle int64           `json:"angle"` // Orientation angle.
}

// DisplayFeature [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#type-DisplayFeature
type DisplayFeature struct {
	Orientation DisplayFeatureOrientation `json:"orientation"` // Orientation of a display feature in relation to screen
	Offset      int64                     `json:"offset"`      // The offset from the screen origin in either the x (for vertical orientation) or y (for horizontal orientation) direction.
	MaskLength  int64                     `json:"maskLength"`  // A display feature may mask content such that it is not physically displayed - this length along with the offset describes this area. A display feature that only splits content will have a 0 mask_length.
}

// DevicePosture [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#type-DevicePosture
type DevicePosture struct {
	Type DevicePostureType `json:"type"` // Current posture of the device
}

// MediaFeature [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#type-MediaFeature
type MediaFeature struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// VirtualTimePolicy advance: If the scheduler runs out of immediate work,
// the virtual time base may fast forward to allow the next delayed task (if
// any) to run; pause: The virtual time base may not advance;
// pauseIfNetworkFetchesPending: The virtual time base may not advance if there
// are any pending resource fetches.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#type-VirtualTimePolicy
type VirtualTimePolicy string

// String returns the VirtualTimePolicy as string value.
func (t VirtualTimePolicy) String() string {
	return string(t)
}

// VirtualTimePolicy values.
const (
	VirtualTimePolicyAdvance                      VirtualTimePolicy = "advance"
	VirtualTimePolicyPause                        VirtualTimePolicy = "pause"
	VirtualTimePolicyPauseIfNetworkFetchesPending VirtualTimePolicy = "pauseIfNetworkFetchesPending"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t VirtualTimePolicy) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t VirtualTimePolicy) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *VirtualTimePolicy) UnmarshalEasyJSON(in *jlexer.Lexer) {
	v := in.String()
	switch VirtualTimePolicy(v) {
	case VirtualTimePolicyAdvance:
		*t = VirtualTimePolicyAdvance
	case VirtualTimePolicyPause:
		*t = VirtualTimePolicyPause
	case VirtualTimePolicyPauseIfNetworkFetchesPending:
		*t = VirtualTimePolicyPauseIfNetworkFetchesPending

	default:
		in.AddError(fmt.Errorf("unknown VirtualTimePolicy value: %v", v))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *VirtualTimePolicy) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// UserAgentBrandVersion used to specify User Agent Client Hints to emulate.
// See https://wicg.github.io/ua-client-hints.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#type-UserAgentBrandVersion
type UserAgentBrandVersion struct {
	Brand   string `json:"brand"`
	Version string `json:"version"`
}

// UserAgentMetadata used to specify User Agent Client Hints to emulate. See
// https://wicg.github.io/ua-client-hints Missing optional values will be filled
// in by the target with what it would normally use.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#type-UserAgentMetadata
type UserAgentMetadata struct {
	Brands          []*UserAgentBrandVersion `json:"brands,omitempty"`          // Brands appearing in Sec-CH-UA.
	FullVersionList []*UserAgentBrandVersion `json:"fullVersionList,omitempty"` // Brands appearing in Sec-CH-UA-Full-Version-List.
	Platform        string                   `json:"platform"`
	PlatformVersion string                   `json:"platformVersion"`
	Architecture    string                   `json:"architecture"`
	Model           string                   `json:"model"`
	Mobile          bool                     `json:"mobile"`
	Bitness         string                   `json:"bitness,omitempty"`
	Wow64           bool                     `json:"wow64,omitempty"`
}

// SensorType used to specify sensor types to emulate. See
// https://w3c.github.io/sensors/#automation for more information.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#type-SensorType
type SensorType string

// String returns the SensorType as string value.
func (t SensorType) String() string {
	return string(t)
}

// SensorType values.
const (
	SensorTypeAbsoluteOrientation SensorType = "absolute-orientation"
	SensorTypeAccelerometer       SensorType = "accelerometer"
	SensorTypeAmbientLight        SensorType = "ambient-light"
	SensorTypeGravity             SensorType = "gravity"
	SensorTypeGyroscope           SensorType = "gyroscope"
	SensorTypeLinearAcceleration  SensorType = "linear-acceleration"
	SensorTypeMagnetometer        SensorType = "magnetometer"
	SensorTypeProximity           SensorType = "proximity"
	SensorTypeRelativeOrientation SensorType = "relative-orientation"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t SensorType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t SensorType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *SensorType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	v := in.String()
	switch SensorType(v) {
	case SensorTypeAbsoluteOrientation:
		*t = SensorTypeAbsoluteOrientation
	case SensorTypeAccelerometer:
		*t = SensorTypeAccelerometer
	case SensorTypeAmbientLight:
		*t = SensorTypeAmbientLight
	case SensorTypeGravity:
		*t = SensorTypeGravity
	case SensorTypeGyroscope:
		*t = SensorTypeGyroscope
	case SensorTypeLinearAcceleration:
		*t = SensorTypeLinearAcceleration
	case SensorTypeMagnetometer:
		*t = SensorTypeMagnetometer
	case SensorTypeProximity:
		*t = SensorTypeProximity
	case SensorTypeRelativeOrientation:
		*t = SensorTypeRelativeOrientation

	default:
		in.AddError(fmt.Errorf("unknown SensorType value: %v", v))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *SensorType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// SensorMetadata [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#type-SensorMetadata
type SensorMetadata struct {
	Available        bool    `json:"available,omitempty"`
	MinimumFrequency float64 `json:"minimumFrequency,omitempty"`
	MaximumFrequency float64 `json:"maximumFrequency,omitempty"`
}

// SensorReadingSingle [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#type-SensorReadingSingle
type SensorReadingSingle struct {
	Value float64 `json:"value"`
}

// SensorReadingXYZ [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#type-SensorReadingXYZ
type SensorReadingXYZ struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

// SensorReadingQuaternion [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#type-SensorReadingQuaternion
type SensorReadingQuaternion struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
	W float64 `json:"w"`
}

// SensorReading [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#type-SensorReading
type SensorReading struct {
	Single     *SensorReadingSingle     `json:"single,omitempty"`
	Xyz        *SensorReadingXYZ        `json:"xyz,omitempty"`
	Quaternion *SensorReadingQuaternion `json:"quaternion,omitempty"`
}

// PressureSource [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#type-PressureSource
type PressureSource string

// String returns the PressureSource as string value.
func (t PressureSource) String() string {
	return string(t)
}

// PressureSource values.
const (
	PressureSourceCPU PressureSource = "cpu"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t PressureSource) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t PressureSource) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *PressureSource) UnmarshalEasyJSON(in *jlexer.Lexer) {
	v := in.String()
	switch PressureSource(v) {
	case PressureSourceCPU:
		*t = PressureSourceCPU

	default:
		in.AddError(fmt.Errorf("unknown PressureSource value: %v", v))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *PressureSource) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// PressureState [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#type-PressureState
type PressureState string

// String returns the PressureState as string value.
func (t PressureState) String() string {
	return string(t)
}

// PressureState values.
const (
	PressureStateNominal  PressureState = "nominal"
	PressureStateFair     PressureState = "fair"
	PressureStateSerious  PressureState = "serious"
	PressureStateCritical PressureState = "critical"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t PressureState) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t PressureState) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *PressureState) UnmarshalEasyJSON(in *jlexer.Lexer) {
	v := in.String()
	switch PressureState(v) {
	case PressureStateNominal:
		*t = PressureStateNominal
	case PressureStateFair:
		*t = PressureStateFair
	case PressureStateSerious:
		*t = PressureStateSerious
	case PressureStateCritical:
		*t = PressureStateCritical

	default:
		in.AddError(fmt.Errorf("unknown PressureState value: %v", v))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *PressureState) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// PressureMetadata [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#type-PressureMetadata
type PressureMetadata struct {
	Available bool `json:"available,omitempty"`
}

// DisabledImageType enum of image types that can be disabled.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#type-DisabledImageType
type DisabledImageType string

// String returns the DisabledImageType as string value.
func (t DisabledImageType) String() string {
	return string(t)
}

// DisabledImageType values.
const (
	DisabledImageTypeAvif DisabledImageType = "avif"
	DisabledImageTypeWebp DisabledImageType = "webp"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t DisabledImageType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t DisabledImageType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *DisabledImageType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	v := in.String()
	switch DisabledImageType(v) {
	case DisabledImageTypeAvif:
		*t = DisabledImageTypeAvif
	case DisabledImageTypeWebp:
		*t = DisabledImageTypeWebp

	default:
		in.AddError(fmt.Errorf("unknown DisabledImageType value: %v", v))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *DisabledImageType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// OrientationType orientation type.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#type-ScreenOrientation
type OrientationType string

// String returns the OrientationType as string value.
func (t OrientationType) String() string {
	return string(t)
}

// OrientationType values.
const (
	OrientationTypePortraitPrimary    OrientationType = "portraitPrimary"
	OrientationTypePortraitSecondary  OrientationType = "portraitSecondary"
	OrientationTypeLandscapePrimary   OrientationType = "landscapePrimary"
	OrientationTypeLandscapeSecondary OrientationType = "landscapeSecondary"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t OrientationType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t OrientationType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *OrientationType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	v := in.String()
	switch OrientationType(v) {
	case OrientationTypePortraitPrimary:
		*t = OrientationTypePortraitPrimary
	case OrientationTypePortraitSecondary:
		*t = OrientationTypePortraitSecondary
	case OrientationTypeLandscapePrimary:
		*t = OrientationTypeLandscapePrimary
	case OrientationTypeLandscapeSecondary:
		*t = OrientationTypeLandscapeSecondary

	default:
		in.AddError(fmt.Errorf("unknown OrientationType value: %v", v))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *OrientationType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// DisplayFeatureOrientation orientation of a display feature in relation to
// screen.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#type-DisplayFeature
type DisplayFeatureOrientation string

// String returns the DisplayFeatureOrientation as string value.
func (t DisplayFeatureOrientation) String() string {
	return string(t)
}

// DisplayFeatureOrientation values.
const (
	DisplayFeatureOrientationVertical   DisplayFeatureOrientation = "vertical"
	DisplayFeatureOrientationHorizontal DisplayFeatureOrientation = "horizontal"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t DisplayFeatureOrientation) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t DisplayFeatureOrientation) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *DisplayFeatureOrientation) UnmarshalEasyJSON(in *jlexer.Lexer) {
	v := in.String()
	switch DisplayFeatureOrientation(v) {
	case DisplayFeatureOrientationVertical:
		*t = DisplayFeatureOrientationVertical
	case DisplayFeatureOrientationHorizontal:
		*t = DisplayFeatureOrientationHorizontal

	default:
		in.AddError(fmt.Errorf("unknown DisplayFeatureOrientation value: %v", v))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *DisplayFeatureOrientation) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// DevicePostureType current posture of the device.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#type-DevicePosture
type DevicePostureType string

// String returns the DevicePostureType as string value.
func (t DevicePostureType) String() string {
	return string(t)
}

// DevicePostureType values.
const (
	DevicePostureTypeContinuous DevicePostureType = "continuous"
	DevicePostureTypeFolded     DevicePostureType = "folded"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t DevicePostureType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t DevicePostureType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *DevicePostureType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	v := in.String()
	switch DevicePostureType(v) {
	case DevicePostureTypeContinuous:
		*t = DevicePostureTypeContinuous
	case DevicePostureTypeFolded:
		*t = DevicePostureTypeFolded

	default:
		in.AddError(fmt.Errorf("unknown DevicePostureType value: %v", v))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *DevicePostureType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// SetEmitTouchEventsForMouseConfiguration touch/gesture events
// configuration. Default: current platform.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#method-setEmitTouchEventsForMouse
type SetEmitTouchEventsForMouseConfiguration string

// String returns the SetEmitTouchEventsForMouseConfiguration as string value.
func (t SetEmitTouchEventsForMouseConfiguration) String() string {
	return string(t)
}

// SetEmitTouchEventsForMouseConfiguration values.
const (
	SetEmitTouchEventsForMouseConfigurationMobile  SetEmitTouchEventsForMouseConfiguration = "mobile"
	SetEmitTouchEventsForMouseConfigurationDesktop SetEmitTouchEventsForMouseConfiguration = "desktop"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t SetEmitTouchEventsForMouseConfiguration) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t SetEmitTouchEventsForMouseConfiguration) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *SetEmitTouchEventsForMouseConfiguration) UnmarshalEasyJSON(in *jlexer.Lexer) {
	v := in.String()
	switch SetEmitTouchEventsForMouseConfiguration(v) {
	case SetEmitTouchEventsForMouseConfigurationMobile:
		*t = SetEmitTouchEventsForMouseConfigurationMobile
	case SetEmitTouchEventsForMouseConfigurationDesktop:
		*t = SetEmitTouchEventsForMouseConfigurationDesktop

	default:
		in.AddError(fmt.Errorf("unknown SetEmitTouchEventsForMouseConfiguration value: %v", v))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *SetEmitTouchEventsForMouseConfiguration) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// SetEmulatedVisionDeficiencyType vision deficiency to emulate. Order:
// best-effort emulations come first, followed by any physiologically accurate
// emulations for medically recognized color vision deficiencies.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Emulation#method-setEmulatedVisionDeficiency
type SetEmulatedVisionDeficiencyType string

// String returns the SetEmulatedVisionDeficiencyType as string value.
func (t SetEmulatedVisionDeficiencyType) String() string {
	return string(t)
}

// SetEmulatedVisionDeficiencyType values.
const (
	SetEmulatedVisionDeficiencyTypeNone            SetEmulatedVisionDeficiencyType = "none"
	SetEmulatedVisionDeficiencyTypeBlurredVision   SetEmulatedVisionDeficiencyType = "blurredVision"
	SetEmulatedVisionDeficiencyTypeReducedContrast SetEmulatedVisionDeficiencyType = "reducedContrast"
	SetEmulatedVisionDeficiencyTypeAchromatopsia   SetEmulatedVisionDeficiencyType = "achromatopsia"
	SetEmulatedVisionDeficiencyTypeDeuteranopia    SetEmulatedVisionDeficiencyType = "deuteranopia"
	SetEmulatedVisionDeficiencyTypeProtanopia      SetEmulatedVisionDeficiencyType = "protanopia"
	SetEmulatedVisionDeficiencyTypeTritanopia      SetEmulatedVisionDeficiencyType = "tritanopia"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t SetEmulatedVisionDeficiencyType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t SetEmulatedVisionDeficiencyType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *SetEmulatedVisionDeficiencyType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	v := in.String()
	switch SetEmulatedVisionDeficiencyType(v) {
	case SetEmulatedVisionDeficiencyTypeNone:
		*t = SetEmulatedVisionDeficiencyTypeNone
	case SetEmulatedVisionDeficiencyTypeBlurredVision:
		*t = SetEmulatedVisionDeficiencyTypeBlurredVision
	case SetEmulatedVisionDeficiencyTypeReducedContrast:
		*t = SetEmulatedVisionDeficiencyTypeReducedContrast
	case SetEmulatedVisionDeficiencyTypeAchromatopsia:
		*t = SetEmulatedVisionDeficiencyTypeAchromatopsia
	case SetEmulatedVisionDeficiencyTypeDeuteranopia:
		*t = SetEmulatedVisionDeficiencyTypeDeuteranopia
	case SetEmulatedVisionDeficiencyTypeProtanopia:
		*t = SetEmulatedVisionDeficiencyTypeProtanopia
	case SetEmulatedVisionDeficiencyTypeTritanopia:
		*t = SetEmulatedVisionDeficiencyTypeTritanopia

	default:
		in.AddError(fmt.Errorf("unknown SetEmulatedVisionDeficiencyType value: %v", v))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *SetEmulatedVisionDeficiencyType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}
