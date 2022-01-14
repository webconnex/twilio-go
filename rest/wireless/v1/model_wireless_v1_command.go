/*
 * Twilio - Wireless
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

// WirelessV1Command struct for WirelessV1Command
type WirelessV1Command struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The message being sent to or from the SIM
	Command *string `json:"command,omitempty"`
	// The mode used to send the SMS message
	CommandMode *string `json:"command_mode,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated format
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// Whether to request a delivery receipt
	DeliveryReceiptRequested *bool `json:"delivery_receipt_requested,omitempty"`
	// The direction of the Command
	Direction *string `json:"direction,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The SID of the Sim resource that the Command was sent to or from
	SimSid *string `json:"sim_sid,omitempty"`
	// The status of the Command
	Status *string `json:"status,omitempty"`
	// The type of transport used
	Transport *string `json:"transport,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"url,omitempty"`
}
