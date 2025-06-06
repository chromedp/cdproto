package dom

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"fmt"
	"strings"

	"github.com/chromedp/cdproto/cdp"
	"github.com/go-json-experiment/json/jsontext"
)

// PhysicalAxes containerSelector physical axes.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#type-PhysicalAxes
type PhysicalAxes string

// String returns the PhysicalAxes as string value.
func (t PhysicalAxes) String() string {
	return string(t)
}

// PhysicalAxes values.
const (
	PhysicalAxesHorizontal PhysicalAxes = "Horizontal"
	PhysicalAxesVertical   PhysicalAxes = "Vertical"
	PhysicalAxesBoth       PhysicalAxes = "Both"
)

// UnmarshalJSON satisfies [json.Unmarshaler].
func (t *PhysicalAxes) UnmarshalJSON(buf []byte) error {
	s := string(buf)
	s = strings.TrimSuffix(strings.TrimPrefix(s, `"`), `"`)

	switch PhysicalAxes(s) {
	case PhysicalAxesHorizontal:
		*t = PhysicalAxesHorizontal
	case PhysicalAxesVertical:
		*t = PhysicalAxesVertical
	case PhysicalAxesBoth:
		*t = PhysicalAxesBoth
	default:
		return fmt.Errorf("unknown PhysicalAxes value: %v", s)
	}
	return nil
}

// LogicalAxes containerSelector logical axes.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#type-LogicalAxes
type LogicalAxes string

// String returns the LogicalAxes as string value.
func (t LogicalAxes) String() string {
	return string(t)
}

// LogicalAxes values.
const (
	LogicalAxesInline LogicalAxes = "Inline"
	LogicalAxesBlock  LogicalAxes = "Block"
	LogicalAxesBoth   LogicalAxes = "Both"
)

// UnmarshalJSON satisfies [json.Unmarshaler].
func (t *LogicalAxes) UnmarshalJSON(buf []byte) error {
	s := string(buf)
	s = strings.TrimSuffix(strings.TrimPrefix(s, `"`), `"`)

	switch LogicalAxes(s) {
	case LogicalAxesInline:
		*t = LogicalAxesInline
	case LogicalAxesBlock:
		*t = LogicalAxesBlock
	case LogicalAxesBoth:
		*t = LogicalAxesBoth
	default:
		return fmt.Errorf("unknown LogicalAxes value: %v", s)
	}
	return nil
}

// ScrollOrientation physical scroll orientation.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#type-ScrollOrientation
type ScrollOrientation string

// String returns the ScrollOrientation as string value.
func (t ScrollOrientation) String() string {
	return string(t)
}

// ScrollOrientation values.
const (
	ScrollOrientationHorizontal ScrollOrientation = "horizontal"
	ScrollOrientationVertical   ScrollOrientation = "vertical"
)

// UnmarshalJSON satisfies [json.Unmarshaler].
func (t *ScrollOrientation) UnmarshalJSON(buf []byte) error {
	s := string(buf)
	s = strings.TrimSuffix(strings.TrimPrefix(s, `"`), `"`)

	switch ScrollOrientation(s) {
	case ScrollOrientationHorizontal:
		*t = ScrollOrientationHorizontal
	case ScrollOrientationVertical:
		*t = ScrollOrientationVertical
	default:
		return fmt.Errorf("unknown ScrollOrientation value: %v", s)
	}
	return nil
}

// DetachedElementInfo a structure to hold the top-level node of a detached
// tree and an array of its retained descendants.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#type-DetachedElementInfo
type DetachedElementInfo struct {
	TreeNode        *cdp.Node    `json:"treeNode"`
	RetainedNodeIDs []cdp.NodeID `json:"retainedNodeIds"`
}

// Quad an array of quad vertices, x immediately followed by y for each
// point, points clock-wise.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#type-Quad
type Quad []float64

// BoxModel box model.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#type-BoxModel
type BoxModel struct {
	Content      Quad              `json:"content"`                         // Content box
	Padding      Quad              `json:"padding"`                         // Padding box
	Border       Quad              `json:"border"`                          // Border box
	Margin       Quad              `json:"margin"`                          // Margin box
	Width        int64             `json:"width"`                           // Node width
	Height       int64             `json:"height"`                          // Node height
	ShapeOutside *ShapeOutsideInfo `json:"shapeOutside,omitempty,omitzero"` // Shape outside coordinates
}

// ShapeOutsideInfo CSS Shape Outside details.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#type-ShapeOutsideInfo
type ShapeOutsideInfo struct {
	Bounds      Quad             `json:"bounds"`      // Shape bounds
	Shape       []jsontext.Value `json:"shape"`       // Shape coordinate details
	MarginShape []jsontext.Value `json:"marginShape"` // Margin shape bounds
}

// Rect Rectangle.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#type-Rect
type Rect struct {
	X      float64 `json:"x"`      // X coordinate
	Y      float64 `json:"y"`      // Y coordinate
	Width  float64 `json:"width"`  // Rectangle width
	Height float64 `json:"height"` // Rectangle height
}

// CSSComputedStyleProperty [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#type-CSSComputedStyleProperty
type CSSComputedStyleProperty struct {
	Name  string `json:"name"`  // Computed style property name.
	Value string `json:"value"` // Computed style property value.
}

// EnableIncludeWhitespace whether to include whitespaces in the children
// array of returned Nodes.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#method-enable
type EnableIncludeWhitespace string

// String returns the EnableIncludeWhitespace as string value.
func (t EnableIncludeWhitespace) String() string {
	return string(t)
}

// EnableIncludeWhitespace values.
const (
	EnableIncludeWhitespaceNone EnableIncludeWhitespace = "none"
	EnableIncludeWhitespaceAll  EnableIncludeWhitespace = "all"
)

// UnmarshalJSON satisfies [json.Unmarshaler].
func (t *EnableIncludeWhitespace) UnmarshalJSON(buf []byte) error {
	s := string(buf)
	s = strings.TrimSuffix(strings.TrimPrefix(s, `"`), `"`)

	switch EnableIncludeWhitespace(s) {
	case EnableIncludeWhitespaceNone:
		*t = EnableIncludeWhitespaceNone
	case EnableIncludeWhitespaceAll:
		*t = EnableIncludeWhitespaceAll
	default:
		return fmt.Errorf("unknown EnableIncludeWhitespace value: %v", s)
	}
	return nil
}

// GetElementByRelationRelation type of relation to get.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOM#method-getElementByRelation
type GetElementByRelationRelation string

// String returns the GetElementByRelationRelation as string value.
func (t GetElementByRelationRelation) String() string {
	return string(t)
}

// GetElementByRelationRelation values.
const (
	GetElementByRelationRelationPopoverTarget  GetElementByRelationRelation = "PopoverTarget"
	GetElementByRelationRelationInterestTarget GetElementByRelationRelation = "InterestTarget"
	GetElementByRelationRelationCommandFor     GetElementByRelationRelation = "CommandFor"
)

// UnmarshalJSON satisfies [json.Unmarshaler].
func (t *GetElementByRelationRelation) UnmarshalJSON(buf []byte) error {
	s := string(buf)
	s = strings.TrimSuffix(strings.TrimPrefix(s, `"`), `"`)

	switch GetElementByRelationRelation(s) {
	case GetElementByRelationRelationPopoverTarget:
		*t = GetElementByRelationRelationPopoverTarget
	case GetElementByRelationRelationInterestTarget:
		*t = GetElementByRelationRelationInterestTarget
	case GetElementByRelationRelationCommandFor:
		*t = GetElementByRelationRelationCommandFor
	default:
		return fmt.Errorf("unknown GetElementByRelationRelation value: %v", s)
	}
	return nil
}
