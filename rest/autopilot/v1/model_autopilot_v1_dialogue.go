/*
 * Twilio - Autopilot
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.25.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// AutopilotV1Dialogue struct for AutopilotV1Dialogue
type AutopilotV1Dialogue struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the Assistant that is the parent of the resource
	AssistantSid *string `json:"assistant_sid,omitempty"`
	// The JSON string that describes the dialogue session object
	Data *map[string]interface{} `json:"data,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The absolute URL of the Dialogue resource
	Url *string `json:"url,omitempty"`
}
