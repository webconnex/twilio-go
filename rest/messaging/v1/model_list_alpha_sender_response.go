/*
 * Twilio - Messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.25.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListAlphaSenderResponse struct for ListAlphaSenderResponse
type ListAlphaSenderResponse struct {
	AlphaSenders []MessagingV1AlphaSender `json:"alpha_senders,omitempty"`
	Meta         ListServiceResponseMeta  `json:"meta,omitempty"`
}
