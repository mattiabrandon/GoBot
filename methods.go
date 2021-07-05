package gobot

import "encoding/json"

type GetUpdatesParams struct {
	Offset         int      `json:"offset"`          // Identifier of the first update to be returned. Must be greater by one than the highest among the identifiers of previously received updates. By default, updates starting with the earliest unconfirmed update are returned. An update is considered confirmed as soon as getUpdates is called with an offset higher than its update_id. The negative offset can be specified to retrieve updates starting from -offset update from the end of the updates queue. All previous updates will forgotten.
	Limit          int      `json:"limit"`           // Limits the number of updates to be retrieved. Values between 1-100 are accepted. Defaults to 100.
	Timeout        int      `json:"timeout"`         // Timeout in seconds for long polling. Defaults to 0, i.e. usual short polling. Should be positive, short polling should be used for testing purposes only.
	AllowedUpdates []string `json:"allowed_updates"` // A JSON-serialized list of the update types you want your bot to receive. For example, specify ["message", "edited_channel_post", "callback_query"] to only receive updates of these types. See Update for a complete list of available update types. Specify an empty list to receive all update types except chat_member (default). If not specified, the previous setting will be used.Please note that this parameter doesn't affect updates created before the call to the getUpdates, so unwanted updates may be received for a short period of time.
}

func (bot GoBot) GetUpdates(params GetUpdatesParams) ([]*Update, error) {
	response, err := bot.Request("getUpdates", params)

	if err != nil {
		return nil, err
	}
	var parsedResponse []*Update
	return parsedResponse, json.Unmarshal(response, &parsedResponse)
}

type SetWebhookParams struct {
	Url                string      `json:"url"`                  // HTTPS url to send updates to. Use an empty string to remove webhook integration
	Certificate        interface{} `json:"certificate"`          // Upload your public key certificate so that the root certificate in use can be checked. See our self-signed guide for details.
	IpAddress          string      `json:"ip_address"`           // The fixed IP address which will be used to send webhook requests instead of the IP address resolved through DNS
	MaxConnections     int         `json:"max_connections"`      // Maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery, 1-100. Defaults to 40. Use lower values to limit the load on your bot's server, and higher values to increase your bot's throughput.
	AllowedUpdates     []string    `json:"allowed_updates"`      // A JSON-serialized list of the update types you want your bot to receive. For example, specify ["message", "edited_channel_post", "callback_query"] to only receive updates of these types. See Update for a complete list of available update types. Specify an empty list to receive all update types except chat_member (default). If not specified, the previous setting will be used.Please note that this parameter doesn't affect updates created before the call to the setWebhook, so unwanted updates may be received for a short period of time.
	DropPendingUpdates bool        `json:"drop_pending_updates"` // Pass True to drop all pending updates
}

func (bot GoBot) SetWebhook(params SetWebhookParams) (bool, error) {
	_, err := bot.Request("setWebhook", params)

	if err != nil {
		return false, err
	}
	return true, nil
}

type DeleteWebhookParams struct {
	DropPendingUpdates bool `json:"drop_pending_updates"` // Pass True to drop all pending updates
}

func (bot GoBot) DeleteWebhook(params DeleteWebhookParams) (bool, error) {
	_, err := bot.Request("deleteWebhook", params)

	if err != nil {
		return false, err
	}
	return true, nil
}

func (bot GoBot) GetWebhookInfo() (*WebhookInfo, error) {
	response, err := bot.Request("getWebhookInfo", nil)

	if err != nil {
		return nil, err
	}
	var parsedResponse *WebhookInfo
	return parsedResponse, json.Unmarshal(response, &parsedResponse)
}

func (bot GoBot) GetMe() (*User, error) {
	response, err := bot.Request("getMe", nil)

	if err != nil {
		return nil, err
	}
	var parsedResponse *User
	return parsedResponse, json.Unmarshal(response, &parsedResponse)
}

func (bot GoBot) LogOut() (bool, error) {
	_, err := bot.Request("logOut", nil)

	if err != nil {
		return false, err
	}
	return true, nil
}

func (bot GoBot) Close() (bool, error) {
	_, err := bot.Request("close", nil)

	if err != nil {
		return false, err
	}
	return true, nil
}

type SendMessageParams struct {
	ChatId                   interface{}      `json:"chat_id"`                     // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Text                     string           `json:"text"`                        // Text of the message to be sent, 1-4096 characters after entities parsing
	ParseMode                string           `json:"parse_mode"`                  // Mode for parsing entities in the message text. See formatting options for more details.
	Entities                 []*MessageEntity `json:"entities"`                    // List of special entities that appear in message text, which can be specified instead of parse_mode
	DisableWebPagePreview    bool             `json:"disable_web_page_preview"`    // Disables link previews for links in this message
	DisableNotification      bool             `json:"disable_notification"`        // Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageId         int              `json:"reply_to_message_id"`         // If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply"` // Pass True, if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              interface{}      `json:"reply_markup"`                // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

func (bot GoBot) SendMessage(params SendMessageParams) (*Message, error) {
	response, err := bot.Request("sendMessage", params)

	if err != nil {
		return nil, err
	}
	var parsedResponse *Message
	return parsedResponse, json.Unmarshal(response, &parsedResponse)
}

type ForwardMessageParams struct {
	ChatId              interface{} `json:"chat_id"`              // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	FromChatId          interface{} `json:"from_chat_id"`         // Unique identifier for the chat where the original message was sent (or channel username in the format @channelusername)
	DisableNotification bool        `json:"disable_notification"` // Sends the message silently. Users will receive a notification with no sound.
	MessageId           int         `json:"message_id"`           // Message identifier in the chat specified in from_chat_id
}

func (bot GoBot) ForwardMessage(params ForwardMessageParams) (*Message, error) {
	response, err := bot.Request("forwardMessage", params)

	if err != nil {
		return nil, err
	}
	var parsedResponse *Message
	return parsedResponse, json.Unmarshal(response, &parsedResponse)
}

type CopyMessageParams struct {
	ChatId                   interface{}      `json:"chat_id"`                     // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	FromChatId               interface{}      `json:"from_chat_id"`                // Unique identifier for the chat where the original message was sent (or channel username in the format @channelusername)
	MessageId                int              `json:"message_id"`                  // Message identifier in the chat specified in from_chat_id
	Caption                  string           `json:"caption"`                     // New caption for media, 0-1024 characters after entities parsing. If not specified, the original caption is kept
	ParseMode                string           `json:"parse_mode"`                  // Mode for parsing entities in the new caption. See formatting options for more details.
	CaptionEntities          []*MessageEntity `json:"caption_entities"`            // List of special entities that appear in the new caption, which can be specified instead of parse_mode
	DisableNotification      bool             `json:"disable_notification"`        // Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageId         int              `json:"reply_to_message_id"`         // If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply"` // Pass True, if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              interface{}      `json:"reply_markup"`                // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

func (bot GoBot) CopyMessage(params CopyMessageParams) (*MessageId, error) {
	response, err := bot.Request("copyMessage", params)

	if err != nil {
		return nil, err
	}
	var parsedResponse *MessageId
	return parsedResponse, json.Unmarshal(response, &parsedResponse)
}

type SendPhotoParams struct {
	ChatId                   interface{}      `json:"chat_id"`                     // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Photo                    interface{}      `json:"photo"`                       // Photo to send. Pass a file_id as String to send a photo that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a photo from the Internet, or upload a new photo using multipart/form-data. The photo must be at most 10 MB in size. The photo's width and height must not exceed 10000 in total. Width and height ratio must be at most 20. More info on Sending Files »
	Caption                  string           `json:"caption"`                     // Photo caption (may also be used when resending photos by file_id), 0-1024 characters after entities parsing
	ParseMode                string           `json:"parse_mode"`                  // Mode for parsing entities in the photo caption. See formatting options for more details.
	CaptionEntities          []*MessageEntity `json:"caption_entities"`            // List of special entities that appear in the caption, which can be specified instead of parse_mode
	DisableNotification      bool             `json:"disable_notification"`        // Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageId         int              `json:"reply_to_message_id"`         // If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply"` // Pass True, if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              interface{}      `json:"reply_markup"`                // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

func (bot GoBot) SendPhoto(params SendPhotoParams) (*Message, error) {
	response, err := bot.Request("sendPhoto", params)

	if err != nil {
		return nil, err
	}
	var parsedResponse *Message
	return parsedResponse, json.Unmarshal(response, &parsedResponse)
}

type SendAudioParams struct {
	ChatId                   interface{}      `json:"chat_id"`                     // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Audio                    interface{}      `json:"audio"`                       // Audio file to send. Pass a file_id as String to send an audio file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get an audio file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
	Caption                  string           `json:"caption"`                     // Audio caption, 0-1024 characters after entities parsing
	ParseMode                string           `json:"parse_mode"`                  // Mode for parsing entities in the audio caption. See formatting options for more details.
	CaptionEntities          []*MessageEntity `json:"caption_entities"`            // List of special entities that appear in the caption, which can be specified instead of parse_mode
	Duration                 int              `json:"duration"`                    // Duration of the audio in seconds
	Performer                string           `json:"performer"`                   // Performer
	Title                    string           `json:"title"`                       // Track name
	Thumb                    interface{}      `json:"thumb"`                       // Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass "attach://<file_attach_name>" if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
	DisableNotification      bool             `json:"disable_notification"`        // Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageId         int              `json:"reply_to_message_id"`         // If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply"` // Pass True, if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              interface{}      `json:"reply_markup"`                // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

func (bot GoBot) SendAudio(params SendAudioParams) (*Message, error) {
	response, err := bot.Request("sendAudio", params)

	if err != nil {
		return nil, err
	}
	var parsedResponse *Message
	return parsedResponse, json.Unmarshal(response, &parsedResponse)
}

type SendDocumentParams struct {
	ChatId                      interface{}      `json:"chat_id"`                        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Document                    interface{}      `json:"document"`                       // File to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
	Thumb                       interface{}      `json:"thumb"`                          // Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass "attach://<file_attach_name>" if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
	Caption                     string           `json:"caption"`                        // Document caption (may also be used when resending documents by file_id), 0-1024 characters after entities parsing
	ParseMode                   string           `json:"parse_mode"`                     // Mode for parsing entities in the document caption. See formatting options for more details.
	CaptionEntities             []*MessageEntity `json:"caption_entities"`               // List of special entities that appear in the caption, which can be specified instead of parse_mode
	DisableContentTypeDetection bool             `json:"disable_content_type_detection"` // Disables automatic server-side content type detection for files uploaded using multipart/form-data
	DisableNotification         bool             `json:"disable_notification"`           // Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageId            int              `json:"reply_to_message_id"`            // If the message is a reply, ID of the original message
	AllowSendingWithoutReply    bool             `json:"allow_sending_without_reply"`    // Pass True, if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup                 interface{}      `json:"reply_markup"`                   // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

func (bot GoBot) SendDocument(params SendDocumentParams) (*Message, error) {
	response, err := bot.Request("sendDocument", params)

	if err != nil {
		return nil, err
	}
	var parsedResponse *Message
	return parsedResponse, json.Unmarshal(response, &parsedResponse)
}

type SendVideoParams struct {
	ChatId                   interface{}      `json:"chat_id"`                     // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Video                    interface{}      `json:"video"`                       // Video to send. Pass a file_id as String to send a video that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a video from the Internet, or upload a new video using multipart/form-data. More info on Sending Files »
	Duration                 int              `json:"duration"`                    // Duration of sent video in seconds
	Width                    int              `json:"width"`                       // Video width
	Height                   int              `json:"height"`                      // Video height
	Thumb                    interface{}      `json:"thumb"`                       // Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass "attach://<file_attach_name>" if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
	Caption                  string           `json:"caption"`                     // Video caption (may also be used when resending videos by file_id), 0-1024 characters after entities parsing
	ParseMode                string           `json:"parse_mode"`                  // Mode for parsing entities in the video caption. See formatting options for more details.
	CaptionEntities          []*MessageEntity `json:"caption_entities"`            // List of special entities that appear in the caption, which can be specified instead of parse_mode
	SupportsStreaming        bool             `json:"supports_streaming"`          // Pass True, if the uploaded video is suitable for streaming
	DisableNotification      bool             `json:"disable_notification"`        // Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageId         int              `json:"reply_to_message_id"`         // If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply"` // Pass True, if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              interface{}      `json:"reply_markup"`                // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

func (bot GoBot) SendVideo(params SendVideoParams) (*Message, error) {
	response, err := bot.Request("sendVideo", params)

	if err != nil {
		return nil, err
	}
	var parsedResponse *Message
	return parsedResponse, json.Unmarshal(response, &parsedResponse)
}

type SendAnimationParams struct {
	ChatId                   interface{}      `json:"chat_id"`                     // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Animation                interface{}      `json:"animation"`                   // Animation to send. Pass a file_id as String to send an animation that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get an animation from the Internet, or upload a new animation using multipart/form-data. More info on Sending Files »
	Duration                 int              `json:"duration"`                    // Duration of sent animation in seconds
	Width                    int              `json:"width"`                       // Animation width
	Height                   int              `json:"height"`                      // Animation height
	Thumb                    interface{}      `json:"thumb"`                       // Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass "attach://<file_attach_name>" if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
	Caption                  string           `json:"caption"`                     // Animation caption (may also be used when resending animation by file_id), 0-1024 characters after entities parsing
	ParseMode                string           `json:"parse_mode"`                  // Mode for parsing entities in the animation caption. See formatting options for more details.
	CaptionEntities          []*MessageEntity `json:"caption_entities"`            // List of special entities that appear in the caption, which can be specified instead of parse_mode
	DisableNotification      bool             `json:"disable_notification"`        // Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageId         int              `json:"reply_to_message_id"`         // If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply"` // Pass True, if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              interface{}      `json:"reply_markup"`                // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

func (bot GoBot) SendAnimation(params SendAnimationParams) (*Message, error) {
	response, err := bot.Request("sendAnimation", params)

	if err != nil {
		return nil, err
	}
	var parsedResponse *Message
	return parsedResponse, json.Unmarshal(response, &parsedResponse)
}

type SendVoiceParams struct {
	ChatId                   interface{}      `json:"chat_id"`                     // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Voice                    interface{}      `json:"voice"`                       // Audio file to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
	Caption                  string           `json:"caption"`                     // Voice message caption, 0-1024 characters after entities parsing
	ParseMode                string           `json:"parse_mode"`                  // Mode for parsing entities in the voice message caption. See formatting options for more details.
	CaptionEntities          []*MessageEntity `json:"caption_entities"`            // List of special entities that appear in the caption, which can be specified instead of parse_mode
	Duration                 int              `json:"duration"`                    // Duration of the voice message in seconds
	DisableNotification      bool             `json:"disable_notification"`        // Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageId         int              `json:"reply_to_message_id"`         // If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply"` // Pass True, if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              interface{}      `json:"reply_markup"`                // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

func (bot GoBot) SendVoice(params SendVoiceParams) (*Message, error) {
	response, err := bot.Request("sendVoice", params)

	if err != nil {
		return nil, err
	}
	var parsedResponse *Message
	return parsedResponse, json.Unmarshal(response, &parsedResponse)
}

type SendVideoNoteParams struct {
	ChatId                   interface{} `json:"chat_id"`                     // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	VideoNote                interface{} `json:"video_note"`                  // Video note to send. Pass a file_id as String to send a video note that exists on the Telegram servers (recommended) or upload a new video using multipart/form-data. More info on Sending Files ». Sending video notes by a URL is currently unsupported
	Duration                 int         `json:"duration"`                    // Duration of sent video in seconds
	Length                   int         `json:"length"`                      // Video width and height, i.e. diameter of the video message
	Thumb                    interface{} `json:"thumb"`                       // Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass "attach://<file_attach_name>" if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
	DisableNotification      bool        `json:"disable_notification"`        // Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageId         int         `json:"reply_to_message_id"`         // If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool        `json:"allow_sending_without_reply"` // Pass True, if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              interface{} `json:"reply_markup"`                // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

func (bot GoBot) SendVideoNote(params SendVideoNoteParams) (*Message, error) {
	response, err := bot.Request("sendVideoNote", params)

	if err != nil {
		return nil, err
	}
	var parsedResponse *Message
	return parsedResponse, json.Unmarshal(response, &parsedResponse)
}

type SendMediaGroupParams struct {
	ChatId                   interface{}   `json:"chat_id"`                     // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Media                    []interface{} `json:"media"`                       // A JSON-serialized array describing messages to be sent, must include 2-10 items
	DisableNotification      bool          `json:"disable_notification"`        // Sends messages silently. Users will receive a notification with no sound.
	ReplyToMessageId         int           `json:"reply_to_message_id"`         // If the messages are a reply, ID of the original message
	AllowSendingWithoutReply bool          `json:"allow_sending_without_reply"` // Pass True, if the message should be sent even if the specified replied-to message is not found
}

func (bot GoBot) SendMediaGroup(params SendMediaGroupParams) ([]*Message, error) {
	response, err := bot.Request("sendMediaGroup", params)

	if err != nil {
		return nil, err
	}
	var parsedResponse []*Message
	return parsedResponse, json.Unmarshal(response, &parsedResponse)
}

type SendLocationParams struct {
	ChatId                   interface{} `json:"chat_id"`                     // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Latitude                 float64     `json:"latitude"`                    // Latitude of the location
	Longitude                float64     `json:"longitude"`                   // Longitude of the location
	HorizontalAccuracy       float64     `json:"horizontal_accuracy"`         // The radius of uncertainty for the location, measured in meters; 0-1500
	LivePeriod               int         `json:"live_period"`                 // Period in seconds for which the location will be updated (see Live Locations, should be between 60 and 86400.
	Heading                  int         `json:"heading"`                     // For live locations, a direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.
	ProximityAlertRadius     int         `json:"proximity_alert_radius"`      // For live locations, a maximum distance for proximity alerts about approaching another chat member, in meters. Must be between 1 and 100000 if specified.
	DisableNotification      bool        `json:"disable_notification"`        // Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageId         int         `json:"reply_to_message_id"`         // If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool        `json:"allow_sending_without_reply"` // Pass True, if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              interface{} `json:"reply_markup"`                // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

func (bot GoBot) SendLocation(params SendLocationParams) (*Message, error) {
	response, err := bot.Request("sendLocation", params)

	if err != nil {
		return nil, err
	}
	var parsedResponse *Message
	return parsedResponse, json.Unmarshal(response, &parsedResponse)
}

type EditMessageLiveLocationParams struct {
	ChatId               interface{}           `json:"chat_id"`                // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId            int                   `json:"message_id"`             // Required if inline_message_id is not specified. Identifier of the message to edit
	InlineMessageId      string                `json:"inline_message_id"`      // Required if chat_id and message_id are not specified. Identifier of the inline message
	Latitude             float64               `json:"latitude"`               // Latitude of new location
	Longitude            float64               `json:"longitude"`              // Longitude of new location
	HorizontalAccuracy   float64               `json:"horizontal_accuracy"`    // The radius of uncertainty for the location, measured in meters; 0-1500
	Heading              int                   `json:"heading"`                // Direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.
	ProximityAlertRadius int                   `json:"proximity_alert_radius"` // Maximum distance for proximity alerts about approaching another chat member, in meters. Must be between 1 and 100000 if specified.
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup"`           // A JSON-serialized object for a new inline keyboard.
}

func (bot GoBot) EditMessageLiveLocation(params EditMessageLiveLocationParams) (interface{}, error) {
	response, err := bot.Request("editMessageLiveLocation", params)

	if err != nil {
		return nil, err
	}
	var parsedResponse interface{}
	return parsedResponse, json.Unmarshal(response, &parsedResponse)
}

type StopMessageLiveLocationParams struct {
	ChatId          interface{}           `json:"chat_id"`           // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId       int                   `json:"message_id"`        // Required if inline_message_id is not specified. Identifier of the message with live location to stop
	InlineMessageId string                `json:"inline_message_id"` // Required if chat_id and message_id are not specified. Identifier of the inline message
	ReplyMarkup     *InlineKeyboardMarkup `json:"reply_markup"`      // A JSON-serialized object for a new inline keyboard.
}

func (bot GoBot) StopMessageLiveLocation(params StopMessageLiveLocationParams) (interface{}, error) {
	response, err := bot.Request("stopMessageLiveLocation", params)

	if err != nil {
		return nil, err
	}
	var parsedResponse interface{}
	return parsedResponse, json.Unmarshal(response, &parsedResponse)
}

type SendVenueParams struct {
	ChatId                   interface{} `json:"chat_id"`                     // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Latitude                 float64     `json:"latitude"`                    // Latitude of the venue
	Longitude                float64     `json:"longitude"`                   // Longitude of the venue
	Title                    string      `json:"title"`                       // Name of the venue
	Address                  string      `json:"address"`                     // Address of the venue
	FoursquareId             string      `json:"foursquare_id"`               // Foursquare identifier of the venue
	FoursquareType           string      `json:"foursquare_type"`             // Foursquare type of the venue, if known. (For example, "arts_entertainment/default", "arts_entertainment/aquarium" or "food/icecream".)
	GooglePlaceId            string      `json:"google_place_id"`             // Google Places identifier of the venue
	GooglePlaceType          string      `json:"google_place_type"`           // Google Places type of the venue. (See supported types.)
	DisableNotification      bool        `json:"disable_notification"`        // Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageId         int         `json:"reply_to_message_id"`         // If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool        `json:"allow_sending_without_reply"` // Pass True, if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              interface{} `json:"reply_markup"`                // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

func (bot GoBot) SendVenue(params SendVenueParams) (*Message, error) {
	response, err := bot.Request("sendVenue", params)

	if err != nil {
		return nil, err
	}
	var parsedResponse *Message
	return parsedResponse, json.Unmarshal(response, &parsedResponse)
}

type SendContactParams struct {
	ChatId                   interface{} `json:"chat_id"`                     // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	PhoneNumber              string      `json:"phone_number"`                // Contact's phone number
	FirstName                string      `json:"first_name"`                  // Contact's first name
	LastName                 string      `json:"last_name"`                   // Contact's last name
	Vcard                    string      `json:"vcard"`                       // Additional data about the contact in the form of a vCard, 0-2048 bytes
	DisableNotification      bool        `json:"disable_notification"`        // Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageId         int         `json:"reply_to_message_id"`         // If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool        `json:"allow_sending_without_reply"` // Pass True, if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              interface{} `json:"reply_markup"`                // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove keyboard or to force a reply from the user.
}

func (bot GoBot) SendContact(params SendContactParams) (*Message, error) {
	response, err := bot.Request("sendContact", params)

	if err != nil {
		return nil, err
	}
	var parsedResponse *Message
	return parsedResponse, json.Unmarshal(response, &parsedResponse)
}

type SendPollParams struct {
	ChatId                   interface{}      `json:"chat_id"`                     // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Question                 string           `json:"question"`                    // Poll question, 1-300 characters
	Options                  []string         `json:"options"`                     // A JSON-serialized list of answer options, 2-10 strings 1-100 characters each
	IsAnonymous              bool             `json:"is_anonymous"`                // True, if the poll needs to be anonymous, defaults to True
	Type                     string           `json:"type"`                        // Poll type, "quiz" or "regular", defaults to "regular"
	AllowsMultipleAnswers    bool             `json:"allows_multiple_answers"`     // True, if the poll allows multiple answers, ignored for polls in quiz mode, defaults to False
	CorrectOptionId          int              `json:"correct_option_id"`           // 0-based identifier of the correct answer option, required for polls in quiz mode
	Explanation              string           `json:"explanation"`                 // Text that is shown when a user chooses an incorrect answer or taps on the lamp icon in a quiz-style poll, 0-200 characters with at most 2 line feeds after entities parsing
	ExplanationParseMode     string           `json:"explanation_parse_mode"`      // Mode for parsing entities in the explanation. See formatting options for more details.
	ExplanationEntities      []*MessageEntity `json:"explanation_entities"`        // List of special entities that appear in the poll explanation, which can be specified instead of parse_mode
	OpenPeriod               int              `json:"open_period"`                 // Amount of time in seconds the poll will be active after creation, 5-600. Can't be used together with close_date.
	CloseDate                int              `json:"close_date"`                  // Point in time (Unix timestamp) when the poll will be automatically closed. Must be at least 5 and no more than 600 seconds in the future. Can't be used together with open_period.
	IsClosed                 bool             `json:"is_closed"`                   // Pass True, if the poll needs to be immediately closed. This can be useful for poll preview.
	DisableNotification      bool             `json:"disable_notification"`        // Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageId         int              `json:"reply_to_message_id"`         // If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply"` // Pass True, if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              interface{}      `json:"reply_markup"`                // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

func (bot GoBot) SendPoll(params SendPollParams) (*Message, error) {
	response, err := bot.Request("sendPoll", params)

	if err != nil {
		return nil, err
	}
	var parsedResponse *Message
	return parsedResponse, json.Unmarshal(response, &parsedResponse)
}

type SendDiceParams struct {
	ChatId                   interface{} `json:"chat_id"`                     // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Emoji                    string      `json:"emoji"`                       // Emoji on which the dice throw animation is based. Currently, must be one of "", "", "", "", "", or "". Dice can have values 1-6 for "", "" and "", values 1-5 for "" and "", and values 1-64 for "". Defaults to ""
	DisableNotification      bool        `json:"disable_notification"`        // Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageId         int         `json:"reply_to_message_id"`         // If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool        `json:"allow_sending_without_reply"` // Pass True, if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              interface{} `json:"reply_markup"`                // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

func (bot GoBot) SendDice(params SendDiceParams) (*Message, error) {
	response, err := bot.Request("sendDice", params)

	if err != nil {
		return nil, err
	}
	var parsedResponse *Message
	return parsedResponse, json.Unmarshal(response, &parsedResponse)
}

type SendChatActionParams struct {
	ChatId interface{} `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Action string      `json:"action"`  // Type of action to broadcast. Choose one, depending on what the user is about to receive: typing for text messages, upload_photo for photos, record_video or upload_video for videos, record_voice or upload_voice for voice notes, upload_document for general files, find_location for location data, record_video_note or upload_video_note for video notes.
}

func (bot GoBot) SendChatAction(params SendChatActionParams) (bool, error) {
	_, err := bot.Request("sendChatAction", params)

	if err != nil {
		return false, err
	}
	return true, nil
}

type GetUserProfilePhotosParams struct {
	UserId int `json:"user_id"` // Unique identifier of the target user
	Offset int `json:"offset"`  // Sequential number of the first photo to be returned. By default, all photos are returned.
	Limit  int `json:"limit"`   // Limits the number of photos to be retrieved. Values between 1-100 are accepted. Defaults to 100.
}

func (bot GoBot) GetUserProfilePhotos(params GetUserProfilePhotosParams) (*UserProfilePhotos, error) {
	response, err := bot.Request("getUserProfilePhotos", params)

	if err != nil {
		return nil, err
	}
	var parsedResponse *UserProfilePhotos
	return parsedResponse, json.Unmarshal(response, &parsedResponse)
}

type GetFileParams struct {
	FileId string `json:"file_id"` // File identifier to get info about
}

func (bot GoBot) GetFile(params GetFileParams) (*File, error) {
	response, err := bot.Request("getFile", params)

	if err != nil {
		return nil, err
	}
	var parsedResponse *File
	return parsedResponse, json.Unmarshal(response, &parsedResponse)
}

type BanChatMemberParams struct {
	ChatId         interface{} `json:"chat_id"`         // Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername)
	UserId         int         `json:"user_id"`         // Unique identifier of the target user
	UntilDate      int         `json:"until_date"`      // Date when the user will be unbanned, unix time. If user is banned for more than 366 days or less than 30 seconds from the current time they are considered to be banned forever. Applied for supergroups and channels only.
	RevokeMessages bool        `json:"revoke_messages"` // Pass True to delete all messages from the chat for the user that is being removed. If False, the user will be able to see messages in the group that were sent before the user was removed. Always True for supergroups and channels.
}

func (bot GoBot) BanChatMember(params BanChatMemberParams) (bool, error) {
	_, err := bot.Request("banChatMember", params)

	if err != nil {
		return false, err
	}
	return true, nil
}

type UnbanChatMemberParams struct {
	ChatId       interface{} `json:"chat_id"`        // Unique identifier for the target group or username of the target supergroup or channel (in the format @username)
	UserId       int         `json:"user_id"`        // Unique identifier of the target user
	OnlyIfBanned bool        `json:"only_if_banned"` // Do nothing if the user is not banned
}

func (bot GoBot) UnbanChatMember(params UnbanChatMemberParams) (bool, error) {
	_, err := bot.Request("unbanChatMember", params)

	if err != nil {
		return false, err
	}
	return true, nil
}

type RestrictChatMemberParams struct {
	ChatId      interface{}      `json:"chat_id"`     // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	UserId      int              `json:"user_id"`     // Unique identifier of the target user
	Permissions *ChatPermissions `json:"permissions"` // A JSON-serialized object for new user permissions
	UntilDate   int              `json:"until_date"`  // Date when restrictions will be lifted for the user, unix time. If user is restricted for more than 366 days or less than 30 seconds from the current time, they are considered to be restricted forever
}

func (bot GoBot) RestrictChatMember(params RestrictChatMemberParams) (bool, error) {
	_, err := bot.Request("restrictChatMember", params)

	if err != nil {
		return false, err
	}
	return true, nil
}

type PromoteChatMemberParams struct {
	ChatId              interface{} `json:"chat_id"`                // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	UserId              int         `json:"user_id"`                // Unique identifier of the target user
	IsAnonymous         bool        `json:"is_anonymous"`           // Pass True, if the administrator's presence in the chat is hidden
	CanManageChat       bool        `json:"can_manage_chat"`        // Pass True, if the administrator can access the chat event log, chat statistics, message statistics in channels, see channel members, see anonymous administrators in supergroups and ignore slow mode. Implied by any other administrator privilege
	CanPostMessages     bool        `json:"can_post_messages"`      // Pass True, if the administrator can create channel posts, channels only
	CanEditMessages     bool        `json:"can_edit_messages"`      // Pass True, if the administrator can edit messages of other users and can pin messages, channels only
	CanDeleteMessages   bool        `json:"can_delete_messages"`    // Pass True, if the administrator can delete messages of other users
	CanManageVoiceChats bool        `json:"can_manage_voice_chats"` // Pass True, if the administrator can manage voice chats
	CanRestrictMembers  bool        `json:"can_restrict_members"`   // Pass True, if the administrator can restrict, ban or unban chat members
	CanPromoteMembers   bool        `json:"can_promote_members"`    // Pass True, if the administrator can add new administrators with a subset of their own privileges or demote administrators that he has promoted, directly or indirectly (promoted by administrators that were appointed by him)
	CanChangeInfo       bool        `json:"can_change_info"`        // Pass True, if the administrator can change chat title, photo and other settings
	CanInviteUsers      bool        `json:"can_invite_users"`       // Pass True, if the administrator can invite new users to the chat
	CanPinMessages      bool        `json:"can_pin_messages"`       // Pass True, if the administrator can pin messages, supergroups only
}

func (bot GoBot) PromoteChatMember(params PromoteChatMemberParams) (bool, error) {
	_, err := bot.Request("promoteChatMember", params)

	if err != nil {
		return false, err
	}
	return true, nil
}

type SetChatAdministratorCustomTitleParams struct {
	ChatId      interface{} `json:"chat_id"`      // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	UserId      int         `json:"user_id"`      // Unique identifier of the target user
	CustomTitle string      `json:"custom_title"` // New custom title for the administrator; 0-16 characters, emoji are not allowed
}

func (bot GoBot) SetChatAdministratorCustomTitle(params SetChatAdministratorCustomTitleParams) (bool, error) {
	_, err := bot.Request("setChatAdministratorCustomTitle", params)

	if err != nil {
		return false, err
	}
	return true, nil
}

type SetChatPermissionsParams struct {
	ChatId      interface{}      `json:"chat_id"`     // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	Permissions *ChatPermissions `json:"permissions"` // New default chat permissions
}

func (bot GoBot) SetChatPermissions(params SetChatPermissionsParams) (bool, error) {
	_, err := bot.Request("setChatPermissions", params)

	if err != nil {
		return false, err
	}
	return true, nil
}

type ExportChatInviteLinkParams struct {
	ChatId interface{} `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
}

func (bot GoBot) ExportChatInviteLink(params ExportChatInviteLinkParams) (string, error) {
	response, err := bot.Request("exportChatInviteLink", params)

	if err != nil {
		return "", err
	}
	var parsedResponse string
	return parsedResponse, json.Unmarshal(response, &parsedResponse)
}

type CreateChatInviteLinkParams struct {
	ChatId      interface{} `json:"chat_id"`      // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ExpireDate  int         `json:"expire_date"`  // Point in time (Unix timestamp) when the link will expire
	MemberLimit int         `json:"member_limit"` // Maximum number of users that can be members of the chat simultaneously after joining the chat via this invite link; 1-99999
}

func (bot GoBot) CreateChatInviteLink(params CreateChatInviteLinkParams) (*ChatInviteLink, error) {
	response, err := bot.Request("createChatInviteLink", params)

	if err != nil {
		return nil, err
	}
	var parsedResponse *ChatInviteLink
	return parsedResponse, json.Unmarshal(response, &parsedResponse)
}

type EditChatInviteLinkParams struct {
	ChatId      interface{} `json:"chat_id"`      // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	InviteLink  string      `json:"invite_link"`  // The invite link to edit
	ExpireDate  int         `json:"expire_date"`  // Point in time (Unix timestamp) when the link will expire
	MemberLimit int         `json:"member_limit"` // Maximum number of users that can be members of the chat simultaneously after joining the chat via this invite link; 1-99999
}

func (bot GoBot) EditChatInviteLink(params EditChatInviteLinkParams) (*ChatInviteLink, error) {
	response, err := bot.Request("editChatInviteLink", params)

	if err != nil {
		return nil, err
	}
	var parsedResponse *ChatInviteLink
	return parsedResponse, json.Unmarshal(response, &parsedResponse)
}

type RevokeChatInviteLinkParams struct {
	ChatId     interface{} `json:"chat_id"`     // Unique identifier of the target chat or username of the target channel (in the format @channelusername)
	InviteLink string      `json:"invite_link"` // The invite link to revoke
}

func (bot GoBot) RevokeChatInviteLink(params RevokeChatInviteLinkParams) (*ChatInviteLink, error) {
	response, err := bot.Request("revokeChatInviteLink", params)

	if err != nil {
		return nil, err
	}
	var parsedResponse *ChatInviteLink
	return parsedResponse, json.Unmarshal(response, &parsedResponse)
}

type SetChatPhotoParams struct {
	ChatId interface{} `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Photo  interface{} `json:"photo"`   // New chat photo, uploaded using multipart/form-data
}

func (bot GoBot) SetChatPhoto(params SetChatPhotoParams) (bool, error) {
	_, err := bot.Request("setChatPhoto", params)

	if err != nil {
		return false, err
	}
	return true, nil
}

type DeleteChatPhotoParams struct {
	ChatId interface{} `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
}

func (bot GoBot) DeleteChatPhoto(params DeleteChatPhotoParams) (bool, error) {
	_, err := bot.Request("deleteChatPhoto", params)

	if err != nil {
		return false, err
	}
	return true, nil
}

type SetChatTitleParams struct {
	ChatId interface{} `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Title  string      `json:"title"`   // New chat title, 1-255 characters
}

func (bot GoBot) SetChatTitle(params SetChatTitleParams) (bool, error) {
	_, err := bot.Request("setChatTitle", params)

	if err != nil {
		return false, err
	}
	return true, nil
}

type SetChatDescriptionParams struct {
	ChatId      interface{} `json:"chat_id"`     // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Description string      `json:"description"` // New chat description, 0-255 characters
}

func (bot GoBot) SetChatDescription(params SetChatDescriptionParams) (bool, error) {
	_, err := bot.Request("setChatDescription", params)

	if err != nil {
		return false, err
	}
	return true, nil
}

type PinChatMessageParams struct {
	ChatId              interface{} `json:"chat_id"`              // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId           int         `json:"message_id"`           // Identifier of a message to pin
	DisableNotification bool        `json:"disable_notification"` // Pass True, if it is not necessary to send a notification to all chat members about the new pinned message. Notifications are always disabled in channels and private chats.
}

func (bot GoBot) PinChatMessage(params PinChatMessageParams) (bool, error) {
	_, err := bot.Request("pinChatMessage", params)

	if err != nil {
		return false, err
	}
	return true, nil
}

type UnpinChatMessageParams struct {
	ChatId    interface{} `json:"chat_id"`    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId int         `json:"message_id"` // Identifier of a message to unpin. If not specified, the most recent pinned message (by sending date) will be unpinned.
}

func (bot GoBot) UnpinChatMessage(params UnpinChatMessageParams) (bool, error) {
	_, err := bot.Request("unpinChatMessage", params)

	if err != nil {
		return false, err
	}
	return true, nil
}

type UnpinAllChatMessagesParams struct {
	ChatId interface{} `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
}

func (bot GoBot) UnpinAllChatMessages(params UnpinAllChatMessagesParams) (bool, error) {
	_, err := bot.Request("unpinAllChatMessages", params)

	if err != nil {
		return false, err
	}
	return true, nil
}

type LeaveChatParams struct {
	ChatId interface{} `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
}

func (bot GoBot) LeaveChat(params LeaveChatParams) (bool, error) {
	_, err := bot.Request("leaveChat", params)

	if err != nil {
		return false, err
	}
	return true, nil
}

type GetChatParams struct {
	ChatId interface{} `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
}

func (bot GoBot) GetChat(params GetChatParams) (*Chat, error) {
	response, err := bot.Request("getChat", params)

	if err != nil {
		return nil, err
	}
	var parsedResponse *Chat
	return parsedResponse, json.Unmarshal(response, &parsedResponse)
}

type GetChatAdministratorsParams struct {
	ChatId interface{} `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
}

func (bot GoBot) GetChatAdministrators(params GetChatAdministratorsParams) ([]interface{}, error) {
	response, err := bot.Request("getChatAdministrators", params)

	if err != nil {
		return nil, err
	}
	var parsedResponse []interface{}
	return parsedResponse, json.Unmarshal(response, &parsedResponse)
}

type GetChatMemberCountParams struct {
	ChatId interface{} `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
}

func (bot GoBot) GetChatMemberCount(params GetChatMemberCountParams) (int, error) {
	response, err := bot.Request("getChatMemberCount", params)

	if err != nil {
		return 0, err
	}
	var parsedResponse int
	return parsedResponse, json.Unmarshal(response, &parsedResponse)
}

type GetChatMemberParams struct {
	ChatId interface{} `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
	UserId int         `json:"user_id"` // Unique identifier of the target user
}

func (bot GoBot) GetChatMember(params GetChatMemberParams) (interface{}, error) {
	response, err := bot.Request("getChatMember", params)

	if err != nil {
		return nil, err
	}
	var parsedResponse interface{}
	return parsedResponse, json.Unmarshal(response, &parsedResponse)
}

type SetChatStickerSetParams struct {
	ChatId         interface{} `json:"chat_id"`          // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	StickerSetName string      `json:"sticker_set_name"` // Name of the sticker set to be set as the group sticker set
}

func (bot GoBot) SetChatStickerSet(params SetChatStickerSetParams) (bool, error) {
	_, err := bot.Request("setChatStickerSet", params)

	if err != nil {
		return false, err
	}
	return true, nil
}

type DeleteChatStickerSetParams struct {
	ChatId interface{} `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
}

func (bot GoBot) DeleteChatStickerSet(params DeleteChatStickerSetParams) (bool, error) {
	_, err := bot.Request("deleteChatStickerSet", params)

	if err != nil {
		return false, err
	}
	return true, nil
}

type AnswerCallbackQueryParams struct {
	CallbackQueryId string `json:"callback_query_id"` // Unique identifier for the query to be answered
	Text            string `json:"text"`              // Text of the notification. If not specified, nothing will be shown to the user, 0-200 characters
	ShowAlert       bool   `json:"show_alert"`        // If true, an alert will be shown by the client instead of a notification at the top of the chat screen. Defaults to false.
	Url             string `json:"url"`               // URL that will be opened by the user's client. If you have created a Game and accepted the conditions via @Botfather, specify the URL that opens your game — note that this will only work if the query comes from a callback_game button.Otherwise, you may use links like t.me/your_bot?start=XXXX that open your bot with a parameter.
	CacheTime       int    `json:"cache_time"`        // The maximum amount of time in seconds that the result of the callback query may be cached client-side. Telegram apps will support caching starting in version 3.14. Defaults to 0.
}

func (bot GoBot) AnswerCallbackQuery(params AnswerCallbackQueryParams) (bool, error) {
	_, err := bot.Request("answerCallbackQuery", params)

	if err != nil {
		return false, err
	}
	return true, nil
}

type SetMyCommandsParams struct {
	Commands     []*BotCommand `json:"commands"`      // A JSON-serialized list of bot commands to be set as the list of the bot's commands. At most 100 commands can be specified.
	Scope        interface{}   `json:"scope"`         // A JSON-serialized object, describing scope of users for which the commands are relevant. Defaults to BotCommandScopeDefault.
	LanguageCode string        `json:"language_code"` // A two-letter ISO 639-1 language code. If empty, commands will be applied to all users from the given scope, for whose language there are no dedicated commands
}

func (bot GoBot) SetMyCommands(params SetMyCommandsParams) (bool, error) {
	_, err := bot.Request("setMyCommands", params)

	if err != nil {
		return false, err
	}
	return true, nil
}

type DeleteMyCommandsParams struct {
	Scope        interface{} `json:"scope"`         // A JSON-serialized object, describing scope of users for which the commands are relevant. Defaults to BotCommandScopeDefault.
	LanguageCode string      `json:"language_code"` // A two-letter ISO 639-1 language code. If empty, commands will be applied to all users from the given scope, for whose language there are no dedicated commands
}

func (bot GoBot) DeleteMyCommands(params DeleteMyCommandsParams) (bool, error) {
	_, err := bot.Request("deleteMyCommands", params)

	if err != nil {
		return false, err
	}
	return true, nil
}

type GetMyCommandsParams struct {
	Scope        interface{} `json:"scope"`         // A JSON-serialized object, describing scope of users. Defaults to BotCommandScopeDefault.
	LanguageCode string      `json:"language_code"` // A two-letter ISO 639-1 language code or an empty string
}

func (bot GoBot) GetMyCommands(params GetMyCommandsParams) ([]*BotCommand, error) {
	response, err := bot.Request("getMyCommands", params)

	if err != nil {
		return nil, err
	}
	var parsedResponse []*BotCommand
	return parsedResponse, json.Unmarshal(response, &parsedResponse)
}

type EditMessageTextParams struct {
	ChatId                interface{}           `json:"chat_id"`                  // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId             int                   `json:"message_id"`               // Required if inline_message_id is not specified. Identifier of the message to edit
	InlineMessageId       string                `json:"inline_message_id"`        // Required if chat_id and message_id are not specified. Identifier of the inline message
	Text                  string                `json:"text"`                     // New text of the message, 1-4096 characters after entities parsing
	ParseMode             string                `json:"parse_mode"`               // Mode for parsing entities in the message text. See formatting options for more details.
	Entities              []*MessageEntity      `json:"entities"`                 // List of special entities that appear in message text, which can be specified instead of parse_mode
	DisableWebPagePreview bool                  `json:"disable_web_page_preview"` // Disables link previews for links in this message
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup"`             // A JSON-serialized object for an inline keyboard.
}

func (bot GoBot) EditMessageText(params EditMessageTextParams) (interface{}, error) {
	response, err := bot.Request("editMessageText", params)

	if err != nil {
		return nil, err
	}
	var parsedResponse interface{}
	return parsedResponse, json.Unmarshal(response, &parsedResponse)
}

type EditMessageCaptionParams struct {
	ChatId          interface{}           `json:"chat_id"`           // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId       int                   `json:"message_id"`        // Required if inline_message_id is not specified. Identifier of the message to edit
	InlineMessageId string                `json:"inline_message_id"` // Required if chat_id and message_id are not specified. Identifier of the inline message
	Caption         string                `json:"caption"`           // New caption of the message, 0-1024 characters after entities parsing
	ParseMode       string                `json:"parse_mode"`        // Mode for parsing entities in the message caption. See formatting options for more details.
	CaptionEntities []*MessageEntity      `json:"caption_entities"`  // List of special entities that appear in the caption, which can be specified instead of parse_mode
	ReplyMarkup     *InlineKeyboardMarkup `json:"reply_markup"`      // A JSON-serialized object for an inline keyboard.
}

func (bot GoBot) EditMessageCaption(params EditMessageCaptionParams) (interface{}, error) {
	response, err := bot.Request("editMessageCaption", params)

	if err != nil {
		return nil, err
	}
	var parsedResponse interface{}
	return parsedResponse, json.Unmarshal(response, &parsedResponse)
}

type EditMessageMediaParams struct {
	ChatId          interface{}           `json:"chat_id"`           // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId       int                   `json:"message_id"`        // Required if inline_message_id is not specified. Identifier of the message to edit
	InlineMessageId string                `json:"inline_message_id"` // Required if chat_id and message_id are not specified. Identifier of the inline message
	Media           interface{}           `json:"media"`             // A JSON-serialized object for a new media content of the message
	ReplyMarkup     *InlineKeyboardMarkup `json:"reply_markup"`      // A JSON-serialized object for a new inline keyboard.
}

func (bot GoBot) EditMessageMedia(params EditMessageMediaParams) (interface{}, error) {
	response, err := bot.Request("editMessageMedia", params)

	if err != nil {
		return nil, err
	}
	var parsedResponse interface{}
	return parsedResponse, json.Unmarshal(response, &parsedResponse)
}

type EditMessageReplyMarkupParams struct {
	ChatId          interface{}           `json:"chat_id"`           // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId       int                   `json:"message_id"`        // Required if inline_message_id is not specified. Identifier of the message to edit
	InlineMessageId string                `json:"inline_message_id"` // Required if chat_id and message_id are not specified. Identifier of the inline message
	ReplyMarkup     *InlineKeyboardMarkup `json:"reply_markup"`      // A JSON-serialized object for an inline keyboard.
}

func (bot GoBot) EditMessageReplyMarkup(params EditMessageReplyMarkupParams) (interface{}, error) {
	response, err := bot.Request("editMessageReplyMarkup", params)

	if err != nil {
		return nil, err
	}
	var parsedResponse interface{}
	return parsedResponse, json.Unmarshal(response, &parsedResponse)
}

type StopPollParams struct {
	ChatId      interface{}           `json:"chat_id"`      // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId   int                   `json:"message_id"`   // Identifier of the original message with the poll
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"` // A JSON-serialized object for a new message inline keyboard.
}

func (bot GoBot) StopPoll(params StopPollParams) (*Poll, error) {
	response, err := bot.Request("stopPoll", params)

	if err != nil {
		return nil, err
	}
	var parsedResponse *Poll
	return parsedResponse, json.Unmarshal(response, &parsedResponse)
}

type DeleteMessageParams struct {
	ChatId    interface{} `json:"chat_id"`    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId int         `json:"message_id"` // Identifier of the message to delete
}

func (bot GoBot) DeleteMessage(params DeleteMessageParams) (bool, error) {
	_, err := bot.Request("deleteMessage", params)

	if err != nil {
		return false, err
	}
	return true, nil
}

type SendStickerParams struct {
	ChatId                   interface{} `json:"chat_id"`                     // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Sticker                  interface{} `json:"sticker"`                     // Sticker to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a .WEBP file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
	DisableNotification      bool        `json:"disable_notification"`        // Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageId         int         `json:"reply_to_message_id"`         // If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool        `json:"allow_sending_without_reply"` // Pass True, if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              interface{} `json:"reply_markup"`                // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

func (bot GoBot) SendSticker(params SendStickerParams) (*Message, error) {
	response, err := bot.Request("sendSticker", params)

	if err != nil {
		return nil, err
	}
	var parsedResponse *Message
	return parsedResponse, json.Unmarshal(response, &parsedResponse)
}

type GetStickerSetParams struct {
	Name string `json:"name"` // Name of the sticker set
}

func (bot GoBot) GetStickerSet(params GetStickerSetParams) (*StickerSet, error) {
	response, err := bot.Request("getStickerSet", params)

	if err != nil {
		return nil, err
	}
	var parsedResponse *StickerSet
	return parsedResponse, json.Unmarshal(response, &parsedResponse)
}

type UploadStickerFileParams struct {
	UserId     int         `json:"user_id"`     // User identifier of sticker file owner
	PngSticker interface{} `json:"png_sticker"` // PNG image with the sticker, must be up to 512 kilobytes in size, dimensions must not exceed 512px, and either width or height must be exactly 512px. More info on Sending Files »
}

func (bot GoBot) UploadStickerFile(params UploadStickerFileParams) (*File, error) {
	response, err := bot.Request("uploadStickerFile", params)

	if err != nil {
		return nil, err
	}
	var parsedResponse *File
	return parsedResponse, json.Unmarshal(response, &parsedResponse)
}

type CreateNewStickerSetParams struct {
	UserId        int           `json:"user_id"`        // User identifier of created sticker set owner
	Name          string        `json:"name"`           // Short name of sticker set, to be used in t.me/addstickers/ URLs (e.g., animals). Can contain only english letters, digits and underscores. Must begin with a letter, can't contain consecutive underscores and must end in "_by_<bot username>". <bot_username> is case insensitive. 1-64 characters.
	Title         string        `json:"title"`          // Sticker set title, 1-64 characters
	PngSticker    interface{}   `json:"png_sticker"`    // PNG image with the sticker, must be up to 512 kilobytes in size, dimensions must not exceed 512px, and either width or height must be exactly 512px. Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
	TgsSticker    interface{}   `json:"tgs_sticker"`    // TGS animation with the sticker, uploaded using multipart/form-data. See https://core.telegram.org/animated_stickers#technical-requirements for technical requirements
	Emojis        string        `json:"emojis"`         // One or more emoji corresponding to the sticker
	ContainsMasks bool          `json:"contains_masks"` // Pass True, if a set of mask stickers should be created
	MaskPosition  *MaskPosition `json:"mask_position"`  // A JSON-serialized object for position where the mask should be placed on faces
}

func (bot GoBot) CreateNewStickerSet(params CreateNewStickerSetParams) (bool, error) {
	_, err := bot.Request("createNewStickerSet", params)

	if err != nil {
		return false, err
	}
	return true, nil
}

type AddStickerToSetParams struct {
	UserId       int           `json:"user_id"`       // User identifier of sticker set owner
	Name         string        `json:"name"`          // Sticker set name
	PngSticker   interface{}   `json:"png_sticker"`   // PNG image with the sticker, must be up to 512 kilobytes in size, dimensions must not exceed 512px, and either width or height must be exactly 512px. Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
	TgsSticker   interface{}   `json:"tgs_sticker"`   // TGS animation with the sticker, uploaded using multipart/form-data. See https://core.telegram.org/animated_stickers#technical-requirements for technical requirements
	Emojis       string        `json:"emojis"`        // One or more emoji corresponding to the sticker
	MaskPosition *MaskPosition `json:"mask_position"` // A JSON-serialized object for position where the mask should be placed on faces
}

func (bot GoBot) AddStickerToSet(params AddStickerToSetParams) (bool, error) {
	_, err := bot.Request("addStickerToSet", params)

	if err != nil {
		return false, err
	}
	return true, nil
}

type SetStickerPositionInSetParams struct {
	Sticker  string `json:"sticker"`  // File identifier of the sticker
	Position int    `json:"position"` // New sticker position in the set, zero-based
}

func (bot GoBot) SetStickerPositionInSet(params SetStickerPositionInSetParams) (bool, error) {
	_, err := bot.Request("setStickerPositionInSet", params)

	if err != nil {
		return false, err
	}
	return true, nil
}

type DeleteStickerFromSetParams struct {
	Sticker string `json:"sticker"` // File identifier of the sticker
}

func (bot GoBot) DeleteStickerFromSet(params DeleteStickerFromSetParams) (bool, error) {
	_, err := bot.Request("deleteStickerFromSet", params)

	if err != nil {
		return false, err
	}
	return true, nil
}

type SetStickerSetThumbParams struct {
	Name   string      `json:"name"`    // Sticker set name
	UserId int         `json:"user_id"` // User identifier of the sticker set owner
	Thumb  interface{} `json:"thumb"`   // A PNG image with the thumbnail, must be up to 128 kilobytes in size and have width and height exactly 100px, or a TGS animation with the thumbnail up to 32 kilobytes in size; see https://core.telegram.org/animated_stickers#technical-requirements for animated sticker technical requirements. Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files ». Animated sticker set thumbnail can't be uploaded via HTTP URL.
}

func (bot GoBot) SetStickerSetThumb(params SetStickerSetThumbParams) (bool, error) {
	_, err := bot.Request("setStickerSetThumb", params)

	if err != nil {
		return false, err
	}
	return true, nil
}

type AnswerInlineQueryParams struct {
	InlineQueryId     string        `json:"inline_query_id"`     // Unique identifier for the answered query
	Results           []interface{} `json:"results"`             // A JSON-serialized array of results for the inline query
	CacheTime         int           `json:"cache_time"`          // The maximum amount of time in seconds that the result of the inline query may be cached on the server. Defaults to 300.
	IsPersonal        bool          `json:"is_personal"`         // Pass True, if results may be cached on the server side only for the user that sent the query. By default, results may be returned to any user who sends the same query
	NextOffset        string        `json:"next_offset"`         // Pass the offset that a client should send in the next query with the same text to receive more results. Pass an empty string if there are no more results or if you don't support pagination. Offset length can't exceed 64 bytes.
	SwitchPmText      string        `json:"switch_pm_text"`      // If passed, clients will display a button with specified text that switches the user to a private chat with the bot and sends the bot a start message with the parameter switch_pm_parameter
	SwitchPmParameter string        `json:"switch_pm_parameter"` // Deep-linking parameter for the /start message sent to the bot when user presses the switch button. 1-64 characters, only A-Z, a-z, 0-9, _ and - are allowed.Example: An inline bot that sends YouTube videos can ask the user to connect the bot to their YouTube account to adapt search results accordingly. To do this, it displays a 'Connect your YouTube account' button above the results, or even before showing any. The user presses the button, switches to a private chat with the bot and, in doing so, passes a start parameter that instructs the bot to return an oauth link. Once done, the bot can offer a switch_inline button so that the user can easily return to the chat where they wanted to use the bot's inline capabilities.
}

func (bot GoBot) AnswerInlineQuery(params AnswerInlineQueryParams) (bool, error) {
	_, err := bot.Request("answerInlineQuery", params)

	if err != nil {
		return false, err
	}
	return true, nil
}

type SendInvoiceParams struct {
	ChatId                    interface{}           `json:"chat_id"`                       // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Title                     string                `json:"title"`                         // Product name, 1-32 characters
	Description               string                `json:"description"`                   // Product description, 1-255 characters
	Payload                   string                `json:"payload"`                       // Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use for your internal processes.
	ProviderToken             string                `json:"provider_token"`                // Payments provider token, obtained via Botfather
	Currency                  string                `json:"currency"`                      // Three-letter ISO 4217 currency code, see more on currencies
	Prices                    []*LabeledPrice       `json:"prices"`                        // Price breakdown, a JSON-serialized list of components (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.)
	MaxTipAmount              int                   `json:"max_tip_amount"`                // The maximum accepted amount for tips in the smallest units of the currency (integer, not float/double). For example, for a maximum tip of US$ 1.45 pass max_tip_amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies). Defaults to 0
	SuggestedTipAmounts       []int                 `json:"suggested_tip_amounts"`         // A JSON-serialized array of suggested amounts of tips in the smallest units of the currency (integer, not float/double). At most 4 suggested tip amounts can be specified. The suggested tip amounts must be positive, passed in a strictly increased order and must not exceed max_tip_amount.
	StartParameter            string                `json:"start_parameter"`               // Unique deep-linking parameter. If left empty, forwarded copies of the sent message will have a Pay button, allowing multiple users to pay directly from the forwarded message, using the same invoice. If non-empty, forwarded copies of the sent message will have a URL button with a deep link to the bot (instead of a Pay button), with the value used as the start parameter
	ProviderData              string                `json:"provider_data"`                 // A JSON-serialized data about the invoice, which will be shared with the payment provider. A detailed description of required fields should be provided by the payment provider.
	PhotoUrl                  string                `json:"photo_url"`                     // URL of the product photo for the invoice. Can be a photo of the goods or a marketing image for a service. People like it better when they see what they are paying for.
	PhotoSize                 int                   `json:"photo_size"`                    // Photo size
	PhotoWidth                int                   `json:"photo_width"`                   // Photo width
	PhotoHeight               int                   `json:"photo_height"`                  // Photo height
	NeedName                  bool                  `json:"need_name"`                     // Pass True, if you require the user's full name to complete the order
	NeedPhoneNumber           bool                  `json:"need_phone_number"`             // Pass True, if you require the user's phone number to complete the order
	NeedEmail                 bool                  `json:"need_email"`                    // Pass True, if you require the user's email address to complete the order
	NeedShippingAddress       bool                  `json:"need_shipping_address"`         // Pass True, if you require the user's shipping address to complete the order
	SendPhoneNumberToProvider bool                  `json:"send_phone_number_to_provider"` // Pass True, if user's phone number should be sent to provider
	SendEmailToProvider       bool                  `json:"send_email_to_provider"`        // Pass True, if user's email address should be sent to provider
	IsFlexible                bool                  `json:"is_flexible"`                   // Pass True, if the final price depends on the shipping method
	DisableNotification       bool                  `json:"disable_notification"`          // Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageId          int                   `json:"reply_to_message_id"`           // If the message is a reply, ID of the original message
	AllowSendingWithoutReply  bool                  `json:"allow_sending_without_reply"`   // Pass True, if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup               *InlineKeyboardMarkup `json:"reply_markup"`                  // A JSON-serialized object for an inline keyboard. If empty, one 'Pay total price' button will be shown. If not empty, the first button must be a Pay button.
}

func (bot GoBot) SendInvoice(params SendInvoiceParams) (*Message, error) {
	response, err := bot.Request("sendInvoice", params)

	if err != nil {
		return nil, err
	}
	var parsedResponse *Message
	return parsedResponse, json.Unmarshal(response, &parsedResponse)
}

type AnswerShippingQueryParams struct {
	ShippingQueryId string            `json:"shipping_query_id"` // Unique identifier for the query to be answered
	Ok              bool              `json:"ok"`                // Specify True if delivery to the specified address is possible and False if there are any problems (for example, if delivery to the specified address is not possible)
	ShippingOptions []*ShippingOption `json:"shipping_options"`  // Required if ok is True. A JSON-serialized array of available shipping options.
	ErrorMessage    string            `json:"error_message"`     // Required if ok is False. Error message in human readable form that explains why it is impossible to complete the order (e.g. "Sorry, delivery to your desired address is unavailable'). Telegram will display this message to the user.
}

func (bot GoBot) AnswerShippingQuery(params AnswerShippingQueryParams) (bool, error) {
	_, err := bot.Request("answerShippingQuery", params)

	if err != nil {
		return false, err
	}
	return true, nil
}

type AnswerPreCheckoutQueryParams struct {
	PreCheckoutQueryId string `json:"pre_checkout_query_id"` // Unique identifier for the query to be answered
	Ok                 bool   `json:"ok"`                    // Specify True if everything is alright (goods are available, etc.) and the bot is ready to proceed with the order. Use False if there are any problems.
	ErrorMessage       string `json:"error_message"`         // Required if ok is False. Error message in human readable form that explains the reason for failure to proceed with the checkout (e.g. "Sorry, somebody just bought the last of our amazing black T-shirts while you were busy filling out your payment details. Please choose a different color or garment!"). Telegram will display this message to the user.
}

func (bot GoBot) AnswerPreCheckoutQuery(params AnswerPreCheckoutQueryParams) (bool, error) {
	_, err := bot.Request("answerPreCheckoutQuery", params)

	if err != nil {
		return false, err
	}
	return true, nil
}

type SetPassportDataErrorsParams struct {
	UserId int           `json:"user_id"` // User identifier
	Errors []interface{} `json:"errors"`  // A JSON-serialized array describing the errors
}

func (bot GoBot) SetPassportDataErrors(params SetPassportDataErrorsParams) (bool, error) {
	_, err := bot.Request("setPassportDataErrors", params)

	if err != nil {
		return false, err
	}
	return true, nil
}

type SendGameParams struct {
	ChatId                   int                   `json:"chat_id"`                     // Unique identifier for the target chat
	GameShortName            string                `json:"game_short_name"`             // Short name of the game, serves as the unique identifier for the game. Set up your games via Botfather.
	DisableNotification      bool                  `json:"disable_notification"`        // Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageId         int                   `json:"reply_to_message_id"`         // If the message is a reply, ID of the original message
	AllowSendingWithoutReply bool                  `json:"allow_sending_without_reply"` // Pass True, if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              *InlineKeyboardMarkup `json:"reply_markup"`                // A JSON-serialized object for an inline keyboard. If empty, one 'Play game_title' button will be shown. If not empty, the first button must launch the game.
}

func (bot GoBot) SendGame(params SendGameParams) (*Message, error) {
	response, err := bot.Request("sendGame", params)

	if err != nil {
		return nil, err
	}
	var parsedResponse *Message
	return parsedResponse, json.Unmarshal(response, &parsedResponse)
}

type SetGameScoreParams struct {
	UserId             int    `json:"user_id"`              // User identifier
	Score              int    `json:"score"`                // New score, must be non-negative
	Force              bool   `json:"force"`                // Pass True, if the high score is allowed to decrease. This can be useful when fixing mistakes or banning cheaters
	DisableEditMessage bool   `json:"disable_edit_message"` // Pass True, if the game message should not be automatically edited to include the current scoreboard
	ChatId             int    `json:"chat_id"`              // Required if inline_message_id is not specified. Unique identifier for the target chat
	MessageId          int    `json:"message_id"`           // Required if inline_message_id is not specified. Identifier of the sent message
	InlineMessageId    string `json:"inline_message_id"`    // Required if chat_id and message_id are not specified. Identifier of the inline message
}

func (bot GoBot) SetGameScore(params SetGameScoreParams) (interface{}, error) {
	response, err := bot.Request("setGameScore", params)

	if err != nil {
		return nil, err
	}
	var parsedResponse interface{}
	return parsedResponse, json.Unmarshal(response, &parsedResponse)
}

type GetGameHighScoresParams struct {
	UserId          int    `json:"user_id"`           // Target user id
	ChatId          int    `json:"chat_id"`           // Required if inline_message_id is not specified. Unique identifier for the target chat
	MessageId       int    `json:"message_id"`        // Required if inline_message_id is not specified. Identifier of the sent message
	InlineMessageId string `json:"inline_message_id"` // Required if chat_id and message_id are not specified. Identifier of the inline message
}

func (bot GoBot) GetGameHighScores(params GetGameHighScoresParams) ([]*GameHighScore, error) {
	response, err := bot.Request("getGameHighScores", params)

	if err != nil {
		return nil, err
	}
	var parsedResponse []*GameHighScore
	return parsedResponse, json.Unmarshal(response, &parsedResponse)
}
