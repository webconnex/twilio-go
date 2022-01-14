/*
 * Twilio - Voice
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.25.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// VoiceV1ConnectionPolicyTarget struct for VoiceV1ConnectionPolicyTarget
type VoiceV1ConnectionPolicyTarget struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the Connection Policy that owns the Target
	ConnectionPolicySid *string `json:"connection_policy_sid,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// Whether the target is enabled
	Enabled *bool `json:"enabled,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The relative importance of the target
	Priority *int `json:"priority,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The SIP address you want Twilio to route your calls to
	Target *string `json:"target,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"url,omitempty"`
	// The value that determines the relative load the Target should receive compared to others with the same priority
	Weight *int `json:"weight,omitempty"`
}
