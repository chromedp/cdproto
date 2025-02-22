package backgroundservice

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"fmt"
	"strings"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/serviceworker"
)

// ServiceName the Background Service that will be associated with the
// commands/events. Every Background Service operates independently, but they
// share the same API.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/BackgroundService#type-ServiceName
type ServiceName string

// String returns the ServiceName as string value.
func (t ServiceName) String() string {
	return string(t)
}

// ServiceName values.
const (
	ServiceNameBackgroundFetch        ServiceName = "backgroundFetch"
	ServiceNameBackgroundSync         ServiceName = "backgroundSync"
	ServiceNamePushMessaging          ServiceName = "pushMessaging"
	ServiceNameNotifications          ServiceName = "notifications"
	ServiceNamePaymentHandler         ServiceName = "paymentHandler"
	ServiceNamePeriodicBackgroundSync ServiceName = "periodicBackgroundSync"
)

// UnmarshalJSON satisfies [json.Unmarshaler].
func (t *ServiceName) UnmarshalJSON(buf []byte) error {
	s := string(buf)
	s = strings.TrimSuffix(strings.TrimPrefix(s, `"`), `"`)

	switch ServiceName(s) {
	case ServiceNameBackgroundFetch:
		*t = ServiceNameBackgroundFetch
	case ServiceNameBackgroundSync:
		*t = ServiceNameBackgroundSync
	case ServiceNamePushMessaging:
		*t = ServiceNamePushMessaging
	case ServiceNameNotifications:
		*t = ServiceNameNotifications
	case ServiceNamePaymentHandler:
		*t = ServiceNamePaymentHandler
	case ServiceNamePeriodicBackgroundSync:
		*t = ServiceNamePeriodicBackgroundSync
	default:
		return fmt.Errorf("unknown ServiceName value: %v", s)
	}
	return nil
}

// EventMetadata a key-value pair for additional event information to pass
// along.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/BackgroundService#type-EventMetadata
type EventMetadata struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Event [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/BackgroundService#type-BackgroundServiceEvent
type Event struct {
	Timestamp                   *cdp.TimeSinceEpoch          `json:"timestamp"`                   // Timestamp of the event (in seconds).
	Origin                      string                       `json:"origin"`                      // The origin this event belongs to.
	ServiceWorkerRegistrationID serviceworker.RegistrationID `json:"serviceWorkerRegistrationId"` // The Service Worker ID that initiated the event.
	Service                     ServiceName                  `json:"service"`                     // The Background Service this event belongs to.
	EventName                   string                       `json:"eventName"`                   // A description of the event.
	InstanceID                  string                       `json:"instanceId"`                  // An identifier that groups related events together.
	EventMetadata               []*EventMetadata             `json:"eventMetadata"`               // A list of event-specific information.
	StorageKey                  string                       `json:"storageKey"`                  // Storage key this event belongs to.
}
