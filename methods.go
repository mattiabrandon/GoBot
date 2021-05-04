// GoBot provides an easy to use interface to interact with
// the BotAPI.
package gobot

import (
	"encoding/json"
	"strconv"
)

// GetUpdatesParams contains all the parameters used to call the GetUpdates method.
type GetUpdatesParams struct {
	Offset         int      `json:"offset"`
	Limit          int      `json:"limit"`
	Timeout        int      `json:"timeout"`
	AllowedUpdates []string `json:"allow_updates"`
}

// GetUpdates is the wrapper for the BotAPI's method https://core.telegram.org/bots/api#getupdates.
func (bot GoBot) GetUpdates(params GetUpdatesParams) ([]*Update, error) {
	result, err := bot.Request("getUpdates", params)

	if err != nil {
		return nil, err
	}
	updates := []*Update{}
	err = json.Unmarshal(result, &updates)
	return updates, err
}

// SetWebhookParams contains all the parameters used to call the SetWebhook method.
type SetWebhookParams struct {
	URL                string   `json:"url"`
	Certificate        string   `json:"certificate"`
	IPAddress          string   `json:"ip_address"`
	MaxConnections     int      `json:"max_connections"`
	AllowedUpdates     []string `json:"allowed_updates"`
	DropPendingUpdates bool     `json:"drop_pending_updates"`
}

// SetWebhook is the wrapper for the BotAPI's method https://core.telegram.org/bots/api#setwebhook.
func (bot GoBot) SetWebhook(params SetWebhookParams) (bool, error) {
	_, err := bot.Request("setWebhook", params)

	if err != nil {
		return false, err
	}
	return true, err
}

// DeleteWebhookParams contains all the parameters used to call the DeleteWebhook method.
type DeleteWebhookParams struct {
	DropPendingUpdates bool `json:"drop_pending_updates"`
}

// DeleteWebhook is the wrapper for the BotAPI's method https://core.telegram.org/bots/api#deletewebhook.
func (bot GoBot) DeleteWebhook(params DeleteWebhookParams) (bool, error) {
	_, err := bot.Request("deleteWebhook", params)

	if err != nil {
		return false, err
	}
	return true, err
}

// GetWebhookInfo is the wrapper for the BotAPI's method https://core.telegram.org/bots/api#getwebhookinfo.
func (bot GoBot) GetWebhookInfo() (WebhookInfo, error) {
	result, err := bot.Request("getWebhookInfo", nil)
	webhookInfo := WebhookInfo{}

	if err != nil {
		return webhookInfo, err
	}
	err = json.Unmarshal(result, &webhookInfo)
	return webhookInfo, err
}

// GetMe is the wrapper for the BotAPI's method https://core.telegram.org/bots/api#getme.
func (bot GoBot) GetMe() (User, error) {
	result, err := bot.Request("getMe", nil)
	user := User{}

	if err != nil {
		return user, err
	}
	err = json.Unmarshal(result, &user)
	return user, err
}

// LogOut is the wrapper for the BotAPI's method https://core.telegram.org/bots/api#logout.
func (bot GoBot) LogOut() (bool, error) {
	_, err := bot.Request("logOut", nil)

	if err != nil {
		return false, err
	}
	return true, err
}

// Close is the wrapper for the BotAPI's method https://core.telegram.org/bots/api#close.
func (bot GoBot) Close() (bool, error) {
	_, err := bot.Request("close", nil)

	if err != nil {
		return false, err
	}
	return true, err
}

// SendMessageParams contains all the parameters used to call the SendMessage method.
type SendMessageParams struct {
	ChatID                   interface{}      `json:"chat_id"`
	Text                     string           `json:"text"`
	ParseMode                string           `json:"parse_mode"`
	Entities                 []*MessageEntity `json:"entities"`
	DisableWebPagePreview    bool             `json:"disable_web_page_preview"`
	DisableNotification      bool             `json:"disable_notification"`
	ReplyToMessageID         int              `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply"`
	ReplyMarkup              interface{}      `json:"reply_markup,omitempty"`
}

// SendMessage is the wrapper for the BotAPI's method https://core.telegram.org/bots/api#sendmessage.
func (bot GoBot) SendMessage(params SendMessageParams) (*Message, error) {
	result, err := bot.Request("sendMessage", params)

	if err != nil {
		return nil, err
	}
	message := &Message{}
	err = json.Unmarshal(result, &message)
	return message, err
}

// ForwardMessageParams contains all the parameters used to call the ForwardMessage method.
type ForwardMessageParams struct {
	ChatID              interface{} `json:"chat_id"`
	FromChatID          int         `json:"from_chat_id"`
	MessageID           int         `json:"message_id"`
	DisableNotification bool        `json:"disable_notification"`
}

// ForwardMessage is the wrapper for the BotAPI's method https://core.telegram.org/bots/api#forwardmessage.
func (bot GoBot) ForwardMessage(params ForwardMessageParams) (*Message, error) {
	result, err := bot.Request("forwardMessage", params)

	if err != nil {
		return nil, err
	}
	message := &Message{}
	err = json.Unmarshal(result, &message)
	return message, err
}

// CopyMessageParams contains all the parameters used to call the CopyMessage method.
type CopyMessageParams struct {
	ChatID                   interface{}     `json:"chat_id"`
	FromChatID               int             `json:"from_chat_id"`
	MessageID                int             `json:"message_id"`
	Caption                  string          `json:"caption"`
	ParseMode                string          `json:"parse_mode"`
	CaptionEntities          []MessageEntity `json:"caption_entities"`
	DisableNotification      bool            `json:"disable_notification"`
	ReplyToMessageID         int             `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool            `json:"allow_sending_without_reply"`
	ReplyMarkup              interface{}     `json:"reply_markup,omitempty"`
}

// CopyMessage is the wrapper for the BotAPI's method https://core.telegram.org/bots/api#copymessage.
func (bot GoBot) CopyMessage(params CopyMessageParams) (*Message, error) {
	result, err := bot.Request("forwardMessage", params)

	if err != nil {
		return nil, err
	}
	message := &Message{}
	err = json.Unmarshal(result, &message)
	return message, err
}

// SendPhotoParams contains all the parameters used to call the SendMedia method.
type SendPhotoParams struct {
	ChatID                   interface{}      `json:"chat_id"`
	Photo                    string           `json:"photo"`
	Caption                  string           `json:"caption"`
	ParseMode                string           `json:"parse_mode"`
	CaptionEntities          []*MessageEntity `json:"caption_entities"`
	DisableNotification      bool             `json:"disable_notification"`
	ReplyToMessageID         int              `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply"`
	ReplyMarkup              interface{}      `json:"reply_markup,omitempty"`
}

// SendAudioParams contains all the parameters used to call the SendMedia method.
type SendAudioParams struct {
	ChatID                   interface{}      `json:"chat_id"`
	Audio                    string           `json:"audio"`
	Caption                  string           `json:"caption"`
	ParseMode                string           `json:"parse_mode"`
	CaptionEntities          []*MessageEntity `json:"caption_entities"`
	Duration                 int              `json:"duration"`
	Performer                string           `json:"performer"`
	Title                    string           `json:"title"`
	Thumb                    string           `json:"thumb"`
	DisableNotification      bool             `json:"disable_notification"`
	ReplyToMessageID         int              `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply"`
	ReplyMarkup              interface{}      `json:"reply_markup,omitempty"`
}

// SendDocumentParams contains all the parameters used to call the SendMedia method.
type SendDocumentParams struct {
	ChatID                      interface{}      `json:"chat_id"`
	Document                    string           `json:"document"`
	Thumb                       string           `json:"thumb"`
	Caption                     string           `json:"caption"`
	ParseMode                   string           `json:"parse_mode"`
	CaptionEntities             []*MessageEntity `json:"caption_entities"`
	DisableContentTypeDetection bool             `json:"disable_content_type_detection"`
	DisableNotification         bool             `json:"disable_notification"`
	ReplyToMessageID            int              `json:"reply_to_message_id"`
	AllowSendingWithoutReply    bool             `json:"allow_sending_without_reply"`
	ReplyMarkup                 interface{}      `json:"reply_markup,omitempty"`
}

// SendVideoParams contains all the parameters used to call the SendMedia method.
type SendVideoParams struct {
	ChatID                   interface{}      `json:"chat_id"`
	Video                    string           `json:"video"`
	Duration                 string           `json:"duration"`
	Width                    int              `json:"width"`
	Height                   int              `json:"height"`
	Thumb                    string           `json:"thumb"`
	Caption                  string           `json:"caption"`
	ParseMode                string           `json:"parse_mode"`
	CaptionEntities          []*MessageEntity `json:"caption_entities"`
	SupportsStreaming        bool             `json:"supports_streaming"`
	DisableNotification      bool             `json:"disable_notification"`
	ReplyToMessageID         int              `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply"`
	ReplyMarkup              interface{}      `json:"reply_markup,omitempty"`
}

// SendAnimationParams contains all the parameters used to call the SendMedia method.
type SendAnimationParams struct {
	ChatID                   interface{}      `json:"chat_id"`
	Animation                string           `json:"animation"`
	Duration                 string           `json:"duration"`
	Width                    int              `json:"width"`
	Height                   int              `json:"height"`
	Thumb                    string           `json:"thumb"`
	Caption                  string           `json:"caption"`
	ParseMode                string           `json:"parse_mode"`
	CaptionEntities          []*MessageEntity `json:"caption_entities"`
	DisableNotification      bool             `json:"disable_notification"`
	ReplyToMessageID         int              `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply"`
	ReplyMarkup              interface{}      `json:"reply_markup,omitempty"`
}

// SendVoiceParams contains all the parameters used to call the SendMedia method.
type SendVoiceParams struct {
	ChatID                   interface{}      `json:"chat_id"`
	Voice                    string           `json:"voice"`
	Caption                  string           `json:"caption"`
	ParseMode                string           `json:"parse_mode"`
	CaptionEntities          []*MessageEntity `json:"caption_entities"`
	Duration                 int              `json:"duration"`
	DisableNotification      bool             `json:"disable_notification"`
	ReplyToMessageID         int              `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply"`
	ReplyMarkup              interface{}      `json:"reply_markup,omitempty"`
}

// SendVideoNoteParams contains all the parameters used to call the SendMedia method.
type SendVideoNoteParams struct {
	ChatID                   interface{}      `json:"chat_id"`
	VideoNote                string           `json:"video_note"`
	Duration                 int              `json:"duration"`
	Length                   int              `json:"length"`
	Caption                  string           `json:"caption"`
	ParseMode                string           `json:"parse_mode"`
	CaptionEntities          []*MessageEntity `json:"caption_entities"`
	DisableNotification      bool             `json:"disable_notification"`
	ReplyToMessageID         int              `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply"`
	ReplyMarkup              interface{}      `json:"reply_markup,omitempty"`
}

// SendMediaGroupParams contains all the parameters used to call the SendMedia method.
type SendMediaGroupParams struct {
	ChatID                   interface{}   `json:"chat_id"`
	Media                    []interface{} `json:"media"`
	DisableNotification      bool          `json:"disable_notification"`
	ReplyToMessageID         int           `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool          `json:"allow_sending_without_reply"`
}

// SendLocationParams contains all the parameters used to call the SendMedia method.
type SendLocationParams struct {
	ChatID                   interface{} `json:"chat_id"`
	Latitude                 float32     `json:"latitude"`
	Longitude                float32     `json:"longitude"`
	HorizontalAccuracy       float32     `json:"horizontal_accuracy"`
	LivePeriod               int         `json:"live_period"`
	Heading                  int         `json:"heading"`
	ProximityAlertRadius     int         `json:"proximity_alert_radius"`
	DisableNotification      bool        `json:"disable_notification"`
	ReplyToMessageID         int         `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool        `json:"allow_sending_without_reply"`
	ReplyMarkup              interface{} `json:"reply_markup,omitempty"`
}

// SendVenueParams contains all the parameters used to call the SendMedia method.
type SendVenueParams struct {
	ChatID                   interface{} `json:"chat_id"`
	Latitude                 float32     `json:"latitude"`
	Longitude                float32     `json:"longitude"`
	Title                    string      `json:"title"`
	Address                  string      `json:"address"`
	ForesquareID             string      `json:"foresquare_id"`
	ForesquareType           string      `json:"foresquare_type"`
	GooglePlaceID            string      `json:"google_place_id"`
	GooglePlaceType          string      `json:"google_place_type"`
	DisableNotification      bool        `json:"disable_notification"`
	ReplyToMessageID         int         `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool        `json:"allow_sending_without_reply"`
	ReplyMarkup              interface{} `json:"reply_markup,omitempty"`
}

// SendContactParams contains all the parameters used to call the SendMedia method.
type SendContactParams struct {
	ChatID                   interface{} `json:"chat_id"`
	PhoneNumber              string      `json:"phone_number"`
	FirstName                string      `json:"first_name"`
	LastName                 string      `json:"last_name"`
	VCard                    string      `json:"v_card"`
	DisableNotification      bool        `json:"disable_notification"`
	ReplyToMessageID         int         `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool        `json:"allow_sending_without_reply"`
	ReplyMarkup              interface{} `json:"reply_markup,omitempty"`
}

// SendPoll contains all the parameters used to call the SendMedia method.
type SendPollParams struct {
	ChatID                   interface{}      `json:"chat_id"`
	Question                 string           `json:"question"`
	Options                  []string         `json:"options"`
	IsAnonymus               bool             `json:"is_anonymous"`
	Type                     string           `json:"type"`
	AllowsMultipleAnswers    bool             `json:"allows_multiple_answers"`
	CorrectOptionID          int              `json:"correct_option_id"`
	Explaination             string           `json:"explaination"`
	ExplainationParseMode    string           `json:"explaination_parse_mode"`
	ExplainationEntities     []*MessageEntity `json:"explaination_entities"`
	OpenPeriod               int              `json:"open_period"`
	CloseDate                int              `json:"close_date"`
	IsClosed                 bool             `json:"is_closed"`
	DisableNotification      bool             `json:"disable_notification"`
	ReplyToMessageID         int              `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply"`
	ReplyMarkup              interface{}      `json:"reply_markup,omitempty"`
}

// SendDice contains all the parameters used to call the SendMedia method.
type SendDiceParams struct {
	ChatID                   interface{} `json:"chat_id"`
	Emoji                    string      `json:"emoji"`
	DisableNotification      bool        `json:"disable_notification"`
	ReplyToMessageID         int         `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool        `json:"allow_sending_without_reply"`
	ReplyMarkup              interface{} `json:"reply_markup,omitempty"`
}

// SendMedia is the wrapper for the following BotAPI's methods:
// https://core.telegram.org/bots/api#sendphoto;
// https://core.telegram.org/bots/api#sendaudio;
// https://core.telegram.org/bots/api#senddocument;
// https://core.telegram.org/bots/api#sendvideo;
// https://core.telegram.org/bots/api#sendanimation;
// https://core.telegram.org/bots/api#sendvoice;
// https://core.telegram.org/bots/api#sendvideonote;
// https://core.telegram.org/bots/api#sendmediagroup;
// https://core.telegram.org/bots/api#sendlocation;
// https://core.telegram.org/bots/api#sendvenue;
// https://core.telegram.org/bots/api#sendcontact;
// https://core.telegram.org/bots/api#sendpoll;
// https://core.telegram.org/bots/api#senddice;
func (bot GoBot) SendMedia(params interface{}) (*Message, error) {
	var method string

	switch params.(type) {
	case SendPhotoParams:
		method = "sendPhoto"
	case SendAudioParams:
		method = "sendAudio"
	case SendDocumentParams:
		method = "sendDocument"
	case SendVideoParams:
		method = "sendVideo"
	case SendAnimationParams:
		method = "sendAnimation"
	case SendVoiceParams:
		method = "sendVoice"
	case SendVideoNoteParams:
		method = "sendVideoNote"
	case SendMediaGroupParams:
		method = "sendMediaGroup"
	case SendLocationParams:
		method = "sendLocation"
	case SendVenueParams:
		method = "sendVenue"
	case SendContactParams:
		method = "sendContact"
	case SendPollParams:
		method = "sendPoll"
	case SendDiceParams:
		method = "sendDice"
	}
	result, err := bot.Request(method, params)

	if err != nil {
		return nil, err
	}
	message := &Message{}
	err = json.Unmarshal(result, &message)
	return message, err
}

// SendChatAction is the wrapper for the BotAPI's method https://core.telegram.org/bots/api#sendchataction.
func (bot GoBot) SendChatAction(chatID interface{}, action string) (bool, error) {
	_, err := bot.Request("sendAudio", map[string]interface{}{
		"chat_id": chatID,
		"action":  action,
	})

	if err != nil {
		return false, err
	}
	return true, nil
}

// GetUserProfilePhotosParams contains all the parameters used to call the GetUserProfilePhotos method.
type GetUserProfilePhotosParams struct {
	UserID int `json:"user_id"`
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

// GetUserProfilePhotos is the wrapper for the BotAPI's method https://core.telegram.org/bots/api#getuserprofilephotos.
func (bot GoBot) GetUserProfilePhotos(params GetUserProfilePhotosParams) (*UserProfilePhotos, error) {
	result, err := bot.Request("getUserProfilePhotos", params)

	if err != nil {
		return nil, err
	}
	userProfilePhotos := &UserProfilePhotos{}
	err = json.Unmarshal(result, &userProfilePhotos)
	return userProfilePhotos, err
}

// GetFile is the wrapper for the BotAPI's method https://core.telegram.org/bots/api#getfile.
func (bot GoBot) GetFile(fileID string) (*File, error) {
	result, err := bot.Request("getUserProfilePhotos", map[string]interface{}{"file_id": fileID})

	if err != nil {
		return nil, err
	}
	file := &File{}
	err = json.Unmarshal(result, &file)
	return file, err
}

// KickChatMemberParams contains all the parameters used to call the KickChatMember method.
type KickChatMemberParams struct {
	ChatID         interface{} `json:"chat_id"`
	UserID         int         `json:"user_id"`
	UntilDate      int         `json:"until_date"`
	RevokeMessages bool        `json:"revoke_messages"`
}

// KickChatMember is the wrapper for the BotAPI's method https://core.telegram.org/bots/api#kickchatmember.
func (bot GoBot) KickChatMember(params KickChatMemberParams) (bool, error) {
	_, err := bot.Request("kickChatMember", params)

	if err != nil {
		return false, err
	}
	return true, err
}

// UnbanChatMemberParams contains all the parameters used to call the UnbanChatMember method.
type UnbanChatMemberParams struct {
	ChatID       interface{} `json:"chat_id"`
	UserID       int         `json:"user_id"`
	OnlyIfBanned bool        `json:"only_if_banned"`
}

// UnbanChatMember is the wrapper for the BotAPI's method https://core.telegram.org/bots/api#unbanchatmember.
func (bot GoBot) UnbanChatMember(params UnbanChatMemberParams) (bool, error) {
	_, err := bot.Request("unbanChatMember", params)

	if err != nil {
		return false, err
	}
	return true, err
}

// RestrictChatMemberParams contains all the parameters used to call the RestrictChatMember method.
type RestrictChatMemberParams struct {
	ChatID      interface{}     `json:"chat_id"`
	UserID      int             `json:"user_id"`
	Offset      int             `json:"offset"`
	Permissions ChatPermissions `json:"permissions"`
	UntilDate   int             `json:"until_date"`
}

// RestrictChatMember is the wrapper for the BotAPI's method https://core.telegram.org/bots/api#restrictchatmember.
func (bot GoBot) RestrictChatMember(params RestrictChatMemberParams) (bool, error) {
	_, err := bot.Request("restrictChatMember", params)

	if err != nil {
		return false, err
	}
	return true, err
}

// PromoteChatMemberParams contains all the parameters used to call the PromoteChatMember method.
type PromoteChatMemberParams struct {
	ChatID              interface{} `json:"chat_id"`
	UserID              int         `json:"user_id"`
	IsAnonymous         bool        `json:"is_anonymous"`
	CanmanageChat       bool        `json:"can_manage_chat"`
	CanPostMessages     bool        `json:"can_post_messages"`
	CanEditMessages     bool        `json:"can_edit_messages"`
	CanDeleteMessages   bool        `json:"can_delete_messages"`
	CanManageVoiceChats bool        `json:"can_manage_voice_chats"`
	CanRestrictMembers  bool        `json:"can_restrict_members"`
	CanPromoteMembers   bool        `json:"can_promote_members"`
	CanChangeInfo       bool        `json:"can_change_info"`
	CanInviteUsers      bool        `json:"can_invite_users"`
	CanPinMessages      bool        `json:"can_pin_messages"`
}

// PromoteChatMember is the wrapper for the BotAPI's method https://core.telegram.org/bots/api#promotechatmember.
func (bot GoBot) PromoteChatMember(params PromoteChatMemberParams) (bool, error) {
	_, err := bot.Request("promoteChatMember", params)

	if err != nil {
		return false, err
	}
	return true, err
}

// TODO: <https://core.telegram.org/bots/api#setchatadministratorcustomtitle>
// TODO: <https://core.telegram.org/bots/api#setchatpermissions>
// TODO: <https://core.telegram.org/bots/api#exportchatinvitelink>
// TODO: <https://core.telegram.org/bots/api#setchatphoto>
// TODO: <https://core.telegram.org/bots/api#deletechatphoto>
// TODO: <https://core.telegram.org/bots/api#setchattitle>
// TODO: <https://core.telegram.org/bots/api#setchatdescription>
// TODO: <https://core.telegram.org/bots/api#pinchatmessage>
// TODO: <https://core.telegram.org/bots/api#unpinchatmessage>
// TODO: <https://core.telegram.org/bots/api#unpinallchatmessages>
// TODO: <https://core.telegram.org/bots/api#leavechat>

// GetChat is the wrapper for the BotAPI's method https://core.telegram.org/bots/api#getchat.
func (bot GoBot) GetChat(chatID interface{}) (Chat, error) {
	result, err := bot.Request("getChat", map[string]interface{}{"chat_id": chatID})
	chat := Chat{}

	if err != nil {
		return chat, err
	}
	err = json.Unmarshal(result, &chat)
	return chat, err
}

// GetChatAdministrators is the wrapper for the BotAPI's method https://core.telegram.org/bots/api#getchatadministrators.
func (bot GoBot) GetChatAdministrators(chatID interface{}) ([]*ChatMember, error) {
	result, err := bot.Request("getChatAdministrators", map[string]interface{}{"chat_id": chatID})

	if err != nil {
		return nil, err
	}
	chatMembers := []*ChatMember{}
	err = json.Unmarshal(result, &chatMembers)
	return chatMembers, err
}

// GetChatMembersCount is the wrapper for the BotAPI's method https://core.telegram.org/bots/api#getchatmemberscount.
func (bot GoBot) GetChatMembersCount(chatID interface{}) (int, error) {
	result, err := bot.Request("getChatMembersCount", map[string]interface{}{"chat_id": chatID})

	if err != nil {
		return 0, err
	}
	return strconv.Atoi(string(result))
}

// TODO: <https://core.telegram.org/bots/api#getchatmember>
// TODO: <https://core.telegram.org/bots/api#setchatstickerset>
// TODO: <https://core.telegram.org/bots/api#deletechatstickerset>

// AnswerCallbackQueryParams contains all the parameters used to call the AnswerCallbackQuery method.
type AnswerCallbackQueryParams struct {
	CallbackQueryID string `json:"chat_id"`
	Text            string `json:"text"`
	ShowAlert       bool   `json:"show_alert"`
	URL             string `json:"url"`
	CacheTime       int    `json:"cache_time"`
}

// AnswerCallbackQuery is the wrapper for the BotAPI's method https://core.telegram.org/bots/api#answercallbackquery.
func (bot GoBot) AnswerCallbackQuery(params AnswerCallbackQueryParams) (bool, error) {
	_, err := bot.Request("answerCallbackQuery", params)

	if err != nil {
		return false, err
	}
	return true, err
}

// TODO: <https://core.telegram.org/bots/api#setmycommands>
// TODO: <https://core.telegram.org/bots/api#getmycommands>

// EditMessageTextParams contains all the parameters used to call the EditMessageText method.
type EditMessageTextParams struct {
	ChatID                interface{}     `json:"chat_id"`
	MessageID             int             `json:"message_id"`
	InlineMessageID       string          `json:"inline_message_id"`
	Text                  string          `json:"text"`
	ParseMode             string          `json:"parse_mode"`
	Entities              []MessageEntity `json:"entities"`
	DisableWebPagePreview bool            `json:"disable_web_page_preview"`
	ReplyMarkup           interface{}     `json:"reply_markup,omitempty"`
}

// EditMessageText is the wrapper for the BotAPI's method https://core.telegram.org/bots/api#editmessagetext.
func (bot GoBot) EditMessageText(params EditMessageTextParams) (*Message, error) {
	result, err := bot.Request("editMessageText", params)

	if err != nil {
		return nil, err
	}
	message := &Message{}
	err = json.Unmarshal(result, &message)
	return message, err
}

// EditMessageCaptionParams contains all the parameters used to call the EditMessageMedia method.
type EditMessageCaptionParams struct {
	ChatID          interface{}     `json:"chat_id"`
	MessageID       int             `json:"message_id"`
	InlineMessageID string          `json:"inline_message_id"`
	Caption         string          `json:"caption"`
	ParseMode       string          `json:"parse_mode"`
	CaptionEntities []MessageEntity `json:"caption_entities"`
	ReplyMarkup     interface{}     `json:"reply_markup,omitempty"`
}

// EditMessageCaption is the wrapper for the BotAPI's method https://core.telegram.org/bots/api#editmessagecaption.
func (bot GoBot) EditMessageCaption(params EditMessageCaptionParams) (*Message, error) {
	result, err := bot.Request("editMessageCaption", params)

	if err != nil {
		return nil, err
	}
	message := &Message{}
	err = json.Unmarshal(result, &message)
	return message, err
}

// EditMessageMediaParams contains all the parameters used to call the EditMessageMedia method.
type EditMessageMediaParams struct {
	ChatID          interface{} `json:"chat_id"`
	MessageID       int         `json:"message_id"`
	InlineMessageID string      `json:"inline_message_id"`
	Media           InputMedia  `json:"media"`
	ReplyMarkup     interface{} `json:"reply_markup,omitempty"`
}

// EditMessageMedia is the wrapper for the BotAPI's method https://core.telegram.org/bots/api#editmessagemedia.
func (bot GoBot) EditMessageMedia(params EditMessageMediaParams) (*Message, error) {
	result, err := bot.Request("editMessageMedia", params)

	if err != nil {
		return nil, err
	}
	message := &Message{}
	err = json.Unmarshal(result, &message)
	return message, err
}

// EditMessageReplyMarkupParams contains all the parameters used to call the EditMessageReplyMarkup method.
type EditMessageReplyMarkupParams struct {
	ChatID          interface{} `json:"chat_id"`
	MessageID       int         `json:"message_id"`
	InlineMessageID string      `json:"inline_message_id"`
	ReplyMarkup     interface{} `json:"reply_markup,omitempty"`
}

// EditMessageReplyMarkup is the wrapper for the BotAPI's method https://core.telegram.org/bots/api#editmessagecaption.
func (bot GoBot) EditMessageReplyMarkup(params EditMessageReplyMarkupParams) (*Message, error) {
	result, err := bot.Request("editMessageReplyMarkup", params)

	if err != nil {
		return nil, err
	}
	message := &Message{}
	err = json.Unmarshal(result, &message)
	return message, err
}

// TODO: <https://core.telegram.org/bots/api#stoppoll>

// DeleteMessageParams contains all the parameters used to call the DeleteMessage method.
type DeleteMessageParams struct {
	ChatID    interface{} `json:"chat_id"`
	MessageID int         `json:"message_id"`
}

// DeleteMessage is the wrapper for the BotAPI's method https://core.telegram.org/bots/api#deletemessage.
func (bot GoBot) DeleteMessage(params DeleteMessageParams) (bool, error) {
	_, err := bot.Request("deleteMessage", params)

	if err != nil {
		return false, err
	}
	return true, err
}

// TODO: <https://core.telegram.org/bots/api#sendsticker>
// TODO: <https://core.telegram.org/bots/api#getstickerset>
// TODO: <https://core.telegram.org/bots/api#uploadstickerfile>
// TODO: <https://core.telegram.org/bots/api#createnewstickerset>
// TODO: <https://core.telegram.org/bots/api#addstickertoset>
// TODO: <https://core.telegram.org/bots/api#setstickerpositioninset>
// TODO: <https://core.telegram.org/bots/api#deletestickerfromset>
// TODO: <https://core.telegram.org/bots/api#setstickersetthumb>
// TODO: <https://core.telegram.org/bots/api#answerinlinequery>
// TODO: <https://core.telegram.org/bots/api#sendinvoice>
// TODO: <https://core.telegram.org/bots/api#answershippingquery>
// TODO: <https://core.telegram.org/bots/api#answerprecheckoutquery>
// TODO: <https://core.telegram.org/bots/api#sendgame>
// TODO: <https://core.telegram.org/bots/api#setpassportdataerrors>
// TODO: <https://core.telegram.org/bots/api#setgamescore>
// TODO: <https://core.telegram.org/bots/api#getgamescore>
// TODO: <https://core.telegram.org/bots/api#getgamehighscores>
