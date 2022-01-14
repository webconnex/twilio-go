/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.25.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"

	"strings"
	"time"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateMessage'
type CreateMessageParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// Determines if the address can be stored or obfuscated based on privacy settings
	AddressRetention *string `json:"AddressRetention,omitempty"`
	// The SID of the application that should receive message status. We POST a `message_sid` parameter and a `message_status` parameter with a value of `sent` or `failed` to the [application](https://www.twilio.com/docs/usage/api/applications)'s `message_status_callback`. If a `status_callback` parameter is also passed, it will be ignored and the application's `message_status_callback` parameter will be used.
	ApplicationSid *string `json:"ApplicationSid,omitempty"`
	// Total number of attempts made ( including this ) to send out the message regardless of the provider used
	Attempt *int `json:"Attempt,omitempty"`
	// The text of the message you want to send. Can be up to 1,600 characters in length.
	Body *string `json:"Body,omitempty"`
	// Determines if the message content can be stored or redacted based on privacy settings
	ContentRetention *string `json:"ContentRetention,omitempty"`
	// Reserved
	ForceDelivery *bool `json:"ForceDelivery,omitempty"`
	// A Twilio phone number in [E.164](https://www.twilio.com/docs/glossary/what-e164) format, an [alphanumeric sender ID](https://www.twilio.com/docs/sms/send-messages#use-an-alphanumeric-sender-id), or a [Channel Endpoint address](https://www.twilio.com/docs/sms/channels#channel-addresses) that is enabled for the type of message you want to send. Phone numbers or [short codes](https://www.twilio.com/docs/sms/api/short-code) purchased from Twilio also work here. You cannot, for example, spoof messages from a private cell phone number. If you are using `messaging_service_sid`, this parameter must be empty.
	From *string `json:"From,omitempty"`
	// The maximum total price in US dollars that you will pay for the message to be delivered. Can be a decimal value that has up to 4 decimal places. All messages are queued for delivery and the message cost is checked before the message is sent. If the cost exceeds `max_price`, the message will fail and a status of `Failed` is sent to the status callback. If `MaxPrice` is not set, the message cost is not checked.
	MaxPrice *float32 `json:"MaxPrice,omitempty"`
	// The URL of the media to send with the message. The media can be of type `gif`, `png`, and `jpeg` and will be formatted correctly on the recipient's device. The media size limit is 5MB for supported file types (JPEG, PNG, GIF) and 500KB for [other types](https://www.twilio.com/docs/sms/accepted-mime-types) of accepted media. To send more than one image in the message body, provide multiple `media_url` parameters in the POST request. You can include up to 10 `media_url` parameters per message. You can send images in an SMS message in only the US and Canada.
	MediaUrl *[]string `json:"MediaUrl,omitempty"`
	// The SID of the [Messaging Service](https://www.twilio.com/docs/sms/services#send-a-message-with-copilot) you want to associate with the Message. Set this parameter to use the [Messaging Service Settings and Copilot Features](https://www.twilio.com/console/sms/services) you have configured and leave the `from` parameter empty. When only this parameter is set, Twilio will use your enabled Copilot Features to select the `from` phone number for delivery.
	MessagingServiceSid *string `json:"MessagingServiceSid,omitempty"`
	// Rich actions for Channels Messages.
	PersistentAction *[]string `json:"PersistentAction,omitempty"`
	// Whether to confirm delivery of the message. Set this value to `true` if you are sending messages that have a trackable user action and you intend to confirm delivery of the message using the [Message Feedback API](https://www.twilio.com/docs/sms/api/message-feedback-resource). This parameter is `false` by default.
	ProvideFeedback *bool `json:"ProvideFeedback,omitempty"`
	// Indicates your intent to schedule a message. Pass the value `fixed` to schedule a message at a fixed time.
	ScheduleType *string `json:"ScheduleType,omitempty"`
	// If set to True, Twilio will deliver the message as a single MMS message, regardless of the presence of media. This is a Beta Feature.
	SendAsMms *bool `json:"SendAsMms,omitempty"`
	// The time that Twilio will send the message. Must be in ISO 8601 format.
	SendAt *time.Time `json:"SendAt,omitempty"`
	// Whether to detect Unicode characters that have a similar GSM-7 character and replace them. Can be: `true` or `false`.
	SmartEncoded *bool `json:"SmartEncoded,omitempty"`
	// The URL we should call using the `status_callback_method` to send status information to your application. If specified, we POST these message status changes to the URL: `queued`, `failed`, `sent`, `delivered`, or `undelivered`. Twilio will POST its [standard request parameters](https://www.twilio.com/docs/sms/twiml#request-parameters) as well as some additional parameters including `MessageSid`, `MessageStatus`, and `ErrorCode`. If you include this parameter with the `messaging_service_sid`, we use this URL instead of the Status Callback URL of the [Messaging Service](https://www.twilio.com/docs/sms/services/api). URLs must contain a valid hostname and underscores are not allowed.
	StatusCallback *string `json:"StatusCallback,omitempty"`
	// The destination phone number in [E.164](https://www.twilio.com/docs/glossary/what-e164) format for SMS/MMS or [Channel user address](https://www.twilio.com/docs/sms/channels#channel-addresses) for other 3rd-party channels.
	To *string `json:"To,omitempty"`
	// How long in seconds the message can remain in our outgoing message queue. After this period elapses, the message fails and we call your status callback. Can be between 1 and the default value of 14,400 seconds. After a message has been accepted by a carrier, however, we cannot guarantee that the message will not be queued after this period. We recommend that this value be at least 5 seconds.
	ValidityPeriod *int `json:"ValidityPeriod,omitempty"`
}

func (params *CreateMessageParams) SetPathAccountSid(PathAccountSid string) *CreateMessageParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *CreateMessageParams) SetAddressRetention(AddressRetention string) *CreateMessageParams {
	params.AddressRetention = &AddressRetention
	return params
}
func (params *CreateMessageParams) SetApplicationSid(ApplicationSid string) *CreateMessageParams {
	params.ApplicationSid = &ApplicationSid
	return params
}
func (params *CreateMessageParams) SetAttempt(Attempt int) *CreateMessageParams {
	params.Attempt = &Attempt
	return params
}
func (params *CreateMessageParams) SetBody(Body string) *CreateMessageParams {
	params.Body = &Body
	return params
}
func (params *CreateMessageParams) SetContentRetention(ContentRetention string) *CreateMessageParams {
	params.ContentRetention = &ContentRetention
	return params
}
func (params *CreateMessageParams) SetForceDelivery(ForceDelivery bool) *CreateMessageParams {
	params.ForceDelivery = &ForceDelivery
	return params
}
func (params *CreateMessageParams) SetFrom(From string) *CreateMessageParams {
	params.From = &From
	return params
}
func (params *CreateMessageParams) SetMaxPrice(MaxPrice float32) *CreateMessageParams {
	params.MaxPrice = &MaxPrice
	return params
}
func (params *CreateMessageParams) SetMediaUrl(MediaUrl []string) *CreateMessageParams {
	params.MediaUrl = &MediaUrl
	return params
}
func (params *CreateMessageParams) SetMessagingServiceSid(MessagingServiceSid string) *CreateMessageParams {
	params.MessagingServiceSid = &MessagingServiceSid
	return params
}
func (params *CreateMessageParams) SetPersistentAction(PersistentAction []string) *CreateMessageParams {
	params.PersistentAction = &PersistentAction
	return params
}
func (params *CreateMessageParams) SetProvideFeedback(ProvideFeedback bool) *CreateMessageParams {
	params.ProvideFeedback = &ProvideFeedback
	return params
}
func (params *CreateMessageParams) SetScheduleType(ScheduleType string) *CreateMessageParams {
	params.ScheduleType = &ScheduleType
	return params
}
func (params *CreateMessageParams) SetSendAsMms(SendAsMms bool) *CreateMessageParams {
	params.SendAsMms = &SendAsMms
	return params
}
func (params *CreateMessageParams) SetSendAt(SendAt time.Time) *CreateMessageParams {
	params.SendAt = &SendAt
	return params
}
func (params *CreateMessageParams) SetSmartEncoded(SmartEncoded bool) *CreateMessageParams {
	params.SmartEncoded = &SmartEncoded
	return params
}
func (params *CreateMessageParams) SetStatusCallback(StatusCallback string) *CreateMessageParams {
	params.StatusCallback = &StatusCallback
	return params
}
func (params *CreateMessageParams) SetTo(To string) *CreateMessageParams {
	params.To = &To
	return params
}
func (params *CreateMessageParams) SetValidityPeriod(ValidityPeriod int) *CreateMessageParams {
	params.ValidityPeriod = &ValidityPeriod
	return params
}

// Send a message from the account used to make the request
func (c *ApiService) CreateMessage(params *CreateMessageParams) (*ApiV2010Message, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Messages.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.AddressRetention != nil {
		data.Set("AddressRetention", *params.AddressRetention)
	}
	if params != nil && params.ApplicationSid != nil {
		data.Set("ApplicationSid", *params.ApplicationSid)
	}
	if params != nil && params.Attempt != nil {
		data.Set("Attempt", fmt.Sprint(*params.Attempt))
	}
	if params != nil && params.Body != nil {
		data.Set("Body", *params.Body)
	}
	if params != nil && params.ContentRetention != nil {
		data.Set("ContentRetention", *params.ContentRetention)
	}
	if params != nil && params.ForceDelivery != nil {
		data.Set("ForceDelivery", fmt.Sprint(*params.ForceDelivery))
	}
	if params != nil && params.From != nil {
		data.Set("From", *params.From)
	}
	if params != nil && params.MaxPrice != nil {
		data.Set("MaxPrice", fmt.Sprint(*params.MaxPrice))
	}
	if params != nil && params.MediaUrl != nil {
		for _, item := range *params.MediaUrl {
			data.Add("MediaUrl", item)
		}
	}
	if params != nil && params.MessagingServiceSid != nil {
		data.Set("MessagingServiceSid", *params.MessagingServiceSid)
	}
	if params != nil && params.PersistentAction != nil {
		for _, item := range *params.PersistentAction {
			data.Add("PersistentAction", item)
		}
	}
	if params != nil && params.ProvideFeedback != nil {
		data.Set("ProvideFeedback", fmt.Sprint(*params.ProvideFeedback))
	}
	if params != nil && params.ScheduleType != nil {
		data.Set("ScheduleType", *params.ScheduleType)
	}
	if params != nil && params.SendAsMms != nil {
		data.Set("SendAsMms", fmt.Sprint(*params.SendAsMms))
	}
	if params != nil && params.SendAt != nil {
		data.Set("SendAt", fmt.Sprint((*params.SendAt).Format(time.RFC3339)))
	}
	if params != nil && params.SmartEncoded != nil {
		data.Set("SmartEncoded", fmt.Sprint(*params.SmartEncoded))
	}
	if params != nil && params.StatusCallback != nil {
		data.Set("StatusCallback", *params.StatusCallback)
	}
	if params != nil && params.To != nil {
		data.Set("To", *params.To)
	}
	if params != nil && params.ValidityPeriod != nil {
		data.Set("ValidityPeriod", fmt.Sprint(*params.ValidityPeriod))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010Message{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteMessage'
type DeleteMessageParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Message resources to delete.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *DeleteMessageParams) SetPathAccountSid(PathAccountSid string) *DeleteMessageParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Deletes a message record from your account
func (c *ApiService) DeleteMessage(Sid string, params *DeleteMessageParams) error {
	path := "/2010-04-01/Accounts/{AccountSid}/Messages/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Optional parameters for the method 'FetchMessage'
type FetchMessageParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Message resource to fetch.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchMessageParams) SetPathAccountSid(PathAccountSid string) *FetchMessageParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Fetch a message belonging to the account used to make the request
func (c *ApiService) FetchMessage(Sid string, params *FetchMessageParams) (*ApiV2010Message, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Messages/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010Message{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListMessage'
type ListMessageParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Message resources to read.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// Read messages sent to only this phone number.
	To *string `json:"To,omitempty"`
	// Read messages sent from only this phone number or alphanumeric sender ID.
	From *string `json:"From,omitempty"`
	// The date of the messages to show. Specify a date as `YYYY-MM-DD` in GMT to read only messages sent on this date. For example: `2009-07-06`. You can also specify an inequality, such as `DateSent<=YYYY-MM-DD`, to read messages sent on or before midnight on a date, and `DateSent>=YYYY-MM-DD` to read messages sent on or after midnight on a date.
	DateSent *time.Time `json:"DateSent,omitempty"`
	// The date of the messages to show. Specify a date as `YYYY-MM-DD` in GMT to read only messages sent on this date. For example: `2009-07-06`. You can also specify an inequality, such as `DateSent<=YYYY-MM-DD`, to read messages sent on or before midnight on a date, and `DateSent>=YYYY-MM-DD` to read messages sent on or after midnight on a date.
	DateSentBefore *time.Time `json:"DateSent&lt;,omitempty"`
	// The date of the messages to show. Specify a date as `YYYY-MM-DD` in GMT to read only messages sent on this date. For example: `2009-07-06`. You can also specify an inequality, such as `DateSent<=YYYY-MM-DD`, to read messages sent on or before midnight on a date, and `DateSent>=YYYY-MM-DD` to read messages sent on or after midnight on a date.
	DateSentAfter *time.Time `json:"DateSent&gt;,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListMessageParams) SetPathAccountSid(PathAccountSid string) *ListMessageParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListMessageParams) SetTo(To string) *ListMessageParams {
	params.To = &To
	return params
}
func (params *ListMessageParams) SetFrom(From string) *ListMessageParams {
	params.From = &From
	return params
}
func (params *ListMessageParams) SetDateSent(DateSent time.Time) *ListMessageParams {
	params.DateSent = &DateSent
	return params
}
func (params *ListMessageParams) SetDateSentBefore(DateSentBefore time.Time) *ListMessageParams {
	params.DateSentBefore = &DateSentBefore
	return params
}
func (params *ListMessageParams) SetDateSentAfter(DateSentAfter time.Time) *ListMessageParams {
	params.DateSentAfter = &DateSentAfter
	return params
}
func (params *ListMessageParams) SetPageSize(PageSize int) *ListMessageParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListMessageParams) SetLimit(Limit int) *ListMessageParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Message records from the API. Request is executed immediately.
func (c *ApiService) PageMessage(params *ListMessageParams, pageToken, pageNumber string) (*ListMessageResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Messages.json"

	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.To != nil {
		data.Set("To", *params.To)
	}
	if params != nil && params.From != nil {
		data.Set("From", *params.From)
	}
	if params != nil && params.DateSent != nil {
		data.Set("DateSent", fmt.Sprint((*params.DateSent).Format(time.RFC3339)))
	}
	if params != nil && params.DateSentBefore != nil {
		data.Set("DateSent<", fmt.Sprint((*params.DateSentBefore).Format(time.RFC3339)))
	}
	if params != nil && params.DateSentAfter != nil {
		data.Set("DateSent>", fmt.Sprint((*params.DateSentAfter).Format(time.RFC3339)))
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListMessageResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Message records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListMessage(params *ListMessageParams) ([]ApiV2010Message, error) {
	if params == nil {
		params = &ListMessageParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageMessage(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ApiV2010Message

	for response != nil {
		records = append(records, response.Messages...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListMessageResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListMessageResponse)
	}

	return records, err
}

// Streams Message records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamMessage(params *ListMessageParams) (chan ApiV2010Message, error) {
	if params == nil {
		params = &ListMessageParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageMessage(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ApiV2010Message, 1)

	go func() {
		for response != nil {
			for item := range response.Messages {
				channel <- response.Messages[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListMessageResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListMessageResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListMessageResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListMessageResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateMessage'
type UpdateMessageParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Message resources to update.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// The text of the message you want to send. Can be up to 1,600 characters long.
	Body *string `json:"Body,omitempty"`
	// When set as `canceled`, allows a message cancelation request if a message has not yet been sent.
	Status *string `json:"Status,omitempty"`
}

func (params *UpdateMessageParams) SetPathAccountSid(PathAccountSid string) *UpdateMessageParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *UpdateMessageParams) SetBody(Body string) *UpdateMessageParams {
	params.Body = &Body
	return params
}
func (params *UpdateMessageParams) SetStatus(Status string) *UpdateMessageParams {
	params.Status = &Status
	return params
}

// To redact a message-body from a post-flight message record, post to the message instance resource with an empty body
func (c *ApiService) UpdateMessage(Sid string, params *UpdateMessageParams) (*ApiV2010Message, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Messages/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Body != nil {
		data.Set("Body", *params.Body)
	}
	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010Message{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
