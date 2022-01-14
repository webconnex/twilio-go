/*
 * Twilio - Verify
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

// VerifyV2Webhook struct for VerifyV2Webhook
type VerifyV2Webhook struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The array of events that this Webhook is subscribed to.
	EventTypes *[]string `json:"event_types,omitempty"`
	// The string that you assigned to describe the webhook
	FriendlyName *string `json:"friendly_name,omitempty"`
	// Service Sid.
	ServiceSid *string `json:"service_sid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The webhook status
	Status *string `json:"status,omitempty"`
	// The absolute URL of the Webhook resource
	Url *string `json:"url,omitempty"`
	// The webhook version
	Version *string `json:"version,omitempty"`
	// The method used when calling the webhook's URL.
	WebhookMethod *string `json:"webhook_method,omitempty"`
	// The URL associated with this Webhook.
	WebhookUrl *string `json:"webhook_url,omitempty"`
}
