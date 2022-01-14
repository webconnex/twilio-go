/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.25.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListTaskQueueResponse struct for ListTaskQueueResponse
type ListTaskQueueResponse struct {
	Meta       ListWorkspaceResponseMeta `json:"meta,omitempty"`
	TaskQueues []TaskrouterV1TaskQueue   `json:"task_queues,omitempty"`
}
