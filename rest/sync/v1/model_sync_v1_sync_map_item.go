/*
 * Twilio - Sync
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

// SyncV1SyncMapItem struct for SyncV1SyncMapItem
type SyncV1SyncMapItem struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The identity of the Map Item's creator
	CreatedBy *string `json:"created_by,omitempty"`
	// An arbitrary, schema-less object that the Map Item stores
	Data *map[string]interface{} `json:"data,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the Map Item expires
	DateExpires *time.Time `json:"date_expires,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The unique, user-defined key for the Map Item
	Key *string `json:"key,omitempty"`
	// The SID of the Sync Map that contains the Map Item
	MapSid *string `json:"map_sid,omitempty"`
	// The current revision of the Map Item, represented as a string
	Revision *string `json:"revision,omitempty"`
	// The SID of the Sync Service that the resource is associated with
	ServiceSid *string `json:"service_sid,omitempty"`
	// The absolute URL of the Map Item resource
	Url *string `json:"url,omitempty"`
}
