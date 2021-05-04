// Package gobot provides an easy to use interface to interact with
// the BotAPI.
package gobot

import "encoding/json"

// Body is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#making-requests.
type Body struct {
	Ok                 bool                `json:"ok"`
	Description        string              `json:"description"`
	ErrorCode          int                 `json:"error_code"`
	Result             json.RawMessage     `json:"result"`
	ResponseParameters *ResponseParameters `json:"parameters"`
}

// Error is the wrapper for the BotAPI's request errors
type Error struct {
	Description        string
	ErrorCode          int
	ResponseParameters *ResponseParameters
}

func (err *Error) Error() string {
	return err.Description
}

// Update is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#update.
type Update struct {
	UpdateID          int                `json:"update_id"`
	Message           *Message           `json:"message"`
	EditedMessage     *Message           `json:"edited_message"`
	ChannelPost       *Message           `json:"channel_post"`
	EditedChannelPost *Message           `json:"edited_channel_post"`
	CallbackQuery     *CallbackQuery     `json:"callback_query"`
	ShippingQuery     *ShippingQuery     `json:"shipping_query"`
	PreCheckoutQuery  *PreCheckoutQuery  `json:"pre_checkout_query"`
	Poll              *Poll              `json:"poll"`
	PollAnswer        *PollAnswer        `json:"poll_answer"`
	MyChatMember      *ChatMemberUpdated `json:"my_chat_member"`
	ChatMember        *ChatMemberUpdated `json:"chat_member"`
	// TODO: InlineQuery       *InlineQuery       `json:"inline_query"`
	// TODO: ChosenInlineQuery *ChosenInlineResult `json:"chosen_inline_result"`
}

// WebhookInfo is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#webhookinfo.
type WebhookInfo struct {
	URL                  string   `json:"url"`
	HasCustomCertificate bool     `json:"has_custom_certificate"`
	PendingUpdateCount   int      `json:"pending_update_count"`
	IPAddress            string   `json:"ip_address"`
	LastErrorDate        int      `json:"last_error_date"`
	LastErrorMessage     string   `json:"last_error_message"`
	MaxConnections       int      `json:"max_connections"`
	AllowedUpdates       []string `json:"allowed_updates"`
}

// User is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#user.
type User struct {
	ID                      int    `json:"id"`
	IsBot                   bool   `json:"is_bot"`
	FirstName               string `json:"first_name"`
	LastName                string `json:"last_name"`
	Username                string `json:"username"`
	LanguageCode            string `json:"language_code"`
	CanJoinGroups           bool   `json:"can_join_groups"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"`
	SupportsInlineQueries   bool   `json:"supports_inline_queries"`
}

// Chat is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#chat.
type Chat struct {
	ID               int              `json:"id"`
	Type             string           `json:"type"`
	Title            string           `json:"title"`
	Username         string           `json:"username"`
	FirstName        string           `json:"first_name"`
	LastName         string           `json:"last_name"`
	Photo            *ChatPhoto       `json:"photo"`
	Bio              string           `json:"bio"`
	Description      string           `json:"description"`
	InviteLink       string           `json:"invite_link"`
	PinnedMessage    *Message         `json:"pinned_message"`
	Permissions      *ChatPermissions `json:"permissions"`
	SlowModeDelay    int              `json:"slow_mode_delay"`
	StickerSetName   string           `json:"sticker_set_name"`
	CanSetStickerSet bool             `json:"can_set_sticker_set"`
	LinkedChatID     int              `json:"linked_chat_id"`
	Location         *ChatLocation    `json:"location"`
}

// Message is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#message.
type Message struct {
	MessageID             int                `json:"message_id"`
	From                  *User              `json:"from"`
	SenderChat            *Chat              `json:"sender_chat"`
	Date                  int                `json:"date"`
	Chat                  *Chat              `json:"chat"`
	ForwardFrom           *User              `json:"forward_from"`
	ForwardFromChat       *Chat              `json:"forward_from_chat"`
	ForwardFromMessageID  int                `json:"forward_from_message_id"`
	ForwardSignature      string             `json:"forward_signature"`
	ForwardSenderName     string             `json:"forward_sender_name"`
	ForwardDate           int                `json:"forward_date"`
	ReplyToMessage        *Message           `json:"reply_to_message"`
	ViaBot                *User              `json:"via_bot"`
	EditDate              int                `json:"edit_date"`
	MediaGroupID          string             `json:"media_group_id"`
	AuthorSignature       string             `json:"author_signature"`
	Text                  string             `json:"text"`
	Entities              []*MessageEntity   `json:"entities"`
	Animation             *Animation         `json:"animation"`
	Audio                 *Audio             `json:"audio"`
	Document              *Document          `json:"document"`
	Photo                 []*PhotoSize       `json:"photo"`
	Sticker               *Sticker           `json:"sticker"`
	Video                 *Video             `json:"video"`
	VideoNote             *VideoNote         `json:"video_note"`
	Caption               string             `json:"caption"`
	CaptionEntities       []*MessageEntity   `json:"caption_entities"`
	Contact               *Contact           `json:"contact"`
	Dice                  *Dice              `json:"dice"`
	Game                  *Game              `json:"game"`
	Poll                  *Poll              `json:"poll"`
	Venue                 *Venue             `json:"venue"`
	Location              *Location          `json:"location"`
	NewChatMembers        []*User            `json:"new_chat_members"`
	LeftChatMember        *User              `json:"left_chat_member"`
	NewChatTitle          string             `json:"new_chat_title"`
	NewChatPhoto          []*PhotoSize       `json:"new_chat_photo"`
	DeleteChatPhoto       bool               `json:"delete_chat_photo"`
	GroupChatCreated      bool               `json:"group_chat_created"`
	SupergroupChatCreated bool               `json:"supergroup_chat_created"`
	ChannelChatCreated    bool               `json:"channel_chat_created"`
	MigrateToChatID       int                `json:"migrate_to_chat_id"`
	MigrateFromChatID     int                `json:"migrate_from_chat_id"`
	PinnedMessage         *Message           `json:"pinned_message"`
	Invoice               *Invoice           `json:"invoice"`
	SuccessfulPayment     *SuccessfulPayment `json:"successful_payment"`
	ConnectedWebsite      string             `json:"connected_website"`
	// PassportData            *PassportData            `json:"passport_data"`
	ProximityAlertTriggered *ProximityAlertTriggered `json:"proximity_alert_triggered"`
	ReplyMarkup             *InlineKeyboardMarkup    `json:"reply_markup"`
}

// MessageID is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#messageid.
type MessageID struct {
	MessageID int `json:"message_id"`
}

// MessageEntity is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#messageentity.
type MessageEntity struct {
	Type     string `json:"type"`
	Offset   int    `json:"offset"`
	Length   int    `json:"length"`
	URL      string `json:"url"`
	User     *User  `json:"user"`
	Language string `json:"language"`
}

// PhotoSize is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#photosize.
type PhotoSize struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	FileSize     int    `json:"file_size"`
}

// Animation is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#animation.
type Animation struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Width        int        `json:"width"`
	Height       int        `json:"height"`
	Duration     int        `json:"duration"`
	Thumb        *PhotoSize `json:"thumb"`
	FileName     string     `json:"filename"`
	MimeType     string     `json:"mime_type"`
	FileSize     int        `json:"file_size"`
}

// Audio is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#audio.
type Audio struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Duration     int        `json:"duration"`
	Performer    string     `json:"performer"`
	Title        string     `json:"title"`
	FileName     string     `json:"filename"`
	MimeType     string     `json:"mime_type"`
	FileSize     int        `json:"file_size"`
	Thumb        *PhotoSize `json:"thumb"`
}

// Document is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#document.
type Document struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Thumb        *PhotoSize `json:"thumb"`
	FileName     string     `json:"filename"`
	MimeType     string     `json:"mime_type"`
	FileSize     int        `json:"file_size"`
}

// Video is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#video.
type Video struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Width        int        `json:"width"`
	Height       int        `json:"height"`
	Duration     int        `json:"duration"`
	Thumb        *PhotoSize `json:"thumb"`
	FileName     string     `json:"filename"`
	MimeType     string     `json:"mime_type"`
	FileSize     int        `json:"file_size"`
}

// VideoNote is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#videonote.
type VideoNote struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Length       int        `json:"length"`
	Duration     int        `json:"duration"`
	Thumb        *PhotoSize `json:"thumb"`
	FileSize     int        `json:"file_size"`
}

// Voice is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#voice.
type Voice struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Duration     int    `json:"duration"`
	MimeType     string `json:"mime_type"`
	FileSize     int    `json:"file_size"`
}

// Contact is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#contact.
type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	UserID      int    `json:"user_id"`
	VCard       string `json:"vcard"`
}

// Dice is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#dice.
type Dice struct {
	Emoji string `json:"emoji"`
	Value int    `json:"value"`
}

// PollOption is the wrapper for the BotAPI's typehttps://core.telegram.org/bots/api#polloption.
type PollOption struct {
	Text       string `json:"text"`
	VoterCount int    `json:"voter_count"`
}

// PollAnswer is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#pollanswer.
type PollAnswer struct {
	PollID    string `json:"poll_id"`
	User      *User  `json:"user"`
	OptionIDs []int  `json:"opotion_ids"`
}

// Poll is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#poll.
type Poll struct {
	ID                   string           `json:"id"`
	Question             string           `json:"question"`
	Options              []*PollOption    `json:"options"`
	TotalVoterCount      int              `json:"total_voter_count"`
	IsClosed             bool             `json:"is_closed"`
	IsAnonymous          bool             `json:"is_anonymous"`
	Type                 string           `json:"type"`
	AllowMultipleAnswers bool             `json:"allows_multiple_answers"`
	CorrectOptionID      int              `json:"correct_option_id"`
	Explanation          string           `json:"explanation"`
	ExplanationEntities  []*MessageEntity `json:"explanation_entities"`
	OpenPeriod           int              `json:"open_period"`
	CloseDate            int              `json:"close_date"`
}

// Location is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#location.
type Location struct {
	Longitude            float32 `json:"longitude"`
	Latitude             float32 `json:"latitude"`
	HorizontalAccuracy   float32 `json:"horizontal_accuracy"`
	LivePeriod           int     `json:"live_period"`
	Heading              int     `json:"heading"`
	ProximityAlertRadius int     `json:"proximity_alert_radius"`
}

// Venue is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#venue.
type Venue struct {
	Location        *Location `json:"location"`
	Title           string    `json:"title"`
	Address         string    `json:"address"`
	FoursquareID    string    `json:"foursquare_id"`
	FoursquareType  string    `json:"foursquare_type"`
	GooglePlaceID   string    `json:"google_place_id"`
	GooglePlaceType string    `json:"google_place_type"`
}

// ProximityAlertTriggered is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#proximityalerttriggered.
type ProximityAlertTriggered struct {
	Traveler *User `json:"traveler"`
	Watcher  *User `json:"watcher"`
	Distance int   `json:"distance"`
}

// UserProfilePhotos is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#userprofilephotos.
type UserProfilePhotos struct {
	TotalCount int            `json:"total_count"`
	Photos     [][]*PhotoSize `json:"photos"`
}

// File is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#file.
type File struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int    `json:"file_size"`
	FilePath     string `json:"file_path"`
}

// ReplyKeyboardMarkup is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#replykeyboardmarkup.
type ReplyKeyboardMarkup struct {
	Keyboard        [][]*KeyboardButton `json:"keyboard"`
	ResizeKeyboard  bool                `json:"resize_keyboard"`
	OneTimeKeyboard bool                `json:"one_time_keyboard"`
	Selective       bool                `json:"selective"`
}

// KeyboardButton is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#keyboardbutton.
type KeyboardButton struct {
	Text           string                  `json:"text"`
	RequestContact bool                    `json:"request_contact"`
	RequestLocatio bool                    `json:"request_location"`
	RequestPoll    *KeyboardButtonPollType `json:"request_poll"`
}

// KeyboardButtonPollType is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#keyboardbutton.
type KeyboardButtonPollType struct {
	Type string `json:"type"`
}

// ReplyKeyboardRemove is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#replykeyboardremove.
type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
	Selective      bool `json:"selective"`
}

// InlineKeyboardMarkup is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#inlinekeyboardmarkup.
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]*InlineKeyboardButton `json:"inline_keyboard"`
}

// InlineKeyboardButton is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#inlinekeyboardbutton.
type InlineKeyboardButton struct {
	Text                         string        `json:"text"`
	URL                          string        `json:"url"`
	LoginURL                     *LoginURL     `json:"login_url"`
	CallbackData                 string        `json:"callback_data"`
	SwitchInlineQuery            string        `json:"switch_inline_query"`
	SwitchInlineQueryCurrentChat string        `json:"switch_inline_query_current_chat"`
	CallbackGame                 *CallbackGame `json:"callback_game"`
	Pay                          bool          `json:"pay"`
}

// LoginURL is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#loginurl.
type LoginURL struct {
	URL                string `json:"url"`
	ForwardText        string `json:"forward_text"`
	BotUsername        string `json:"bot_username"`
	RequestWriteAccess bool   `json:"request_write_access"`
}

// CallbackQuery is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#callbackquery.
type CallbackQuery struct {
	ID              string   `json:"id"`
	From            *User    `json:"from"`
	Message         *Message `json:"message"`
	InlineMessageID string   `json:"inline_message_id"`
	ChatInstance    string   `json:"chat_instance"`
	Data            string   `json:"data"`
	GameShortName   bool     `json:"game_short_name"`
}

// ForceReply is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#forcereply.
type ForceReply struct {
	ForceReply bool `json:"force_reply"`
	Selective  bool `json:"selective"`
}

// ChatPhoto is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#chatphoto.
type ChatPhoto struct {
	SmallFileID       string `json:"small_file_id"`
	SmallFileUniqueID string `json:"small_file_unique_id"`
	BigFileID         string `json:"big_file_id"`
	BigFileUniqueID   string `json:"big_file_unique_id"`
}

type ChatInviteLink struct {
	InviteLink  string `json:"invite_link"`
	Creator     *User  `json:"creator"`
	IsPrimary   bool   `json:"is_primary"`
	IsRevoked   bool   `json:"is_revoked"`
	ExpireDate  int    `json:"expire_date"`
	MemberLimit int    `json:"member_limit"`
}

// ChatMember is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#chatmember.
type ChatMember struct {
	User                  *User  `json:"user"`
	Status                string `json:"status"`
	CustomTitle           string `json:"custom_title"`
	IsAnonymous           bool   `json:"is_anonymous"`
	CanBeEdited           bool   `json:"can_be_edited"`
	CanPostMessages       bool   `json:"can_post_messages"`
	CanEditMessages       bool   `json:"can_edit_messages"`
	CanDeleteMessages     bool   `json:"can_delete_messages"`
	CanRestrictMembers    bool   `json:"can_restrict_members"`
	CanPromoteMembers     bool   `json:"can_promote_members"`
	CanChangeInfo         bool   `json:"can_change_info"`
	CanInviteUsers        bool   `json:"can_invite_users"`
	CanPinMessages        bool   `json:"can_pin_messages"`
	IsMember              bool   `json:"is_member"`
	CanSendMessages       bool   `json:"can_send_messages"`
	CanSendMediaMessages  bool   `json:"can_send_media_messages"`
	CanSendPolls          bool   `json:"can_send_polls"`
	CanSendOtherMessages  bool   `json:"can_send_other_messages"`
	CanAddWebPagePreviews bool   `json:"can_add_web_page_previews"`
	UntilDate             int    `json:"until_date"`
}

type ChatMemberUpdated struct {
	Chat          *Chat           `json:"chat"`
	From          *User           `json:"from"`
	Date          int             `json:"date"`
	OldChatMember *ChatMember     `json:"old_chat_member"`
	NewChatMember *ChatMember     `json:"new_chat_member"`
	InviteLink    *ChatInviteLink `json:"invite_link"`
}

// ChatPermissions is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#chatpermissions.
type ChatPermissions struct {
	CanSendMessages       bool `json:"can_send_messages"`
	CanSendMediaMessages  bool `json:"can_send_media_messages"`
	CanSendPolls          bool `json:"can_send_polls"`
	CanSendOtherMessages  bool `json:"can_send_other_messages"`
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews"`
	CanChangeInfo         bool `json:"can_change_info"`
	CanInviteUsers        bool `json:"can_invite_users"`
	CanPinMessages        bool `json:"can_pin_messages"`
}

// ChatLocation is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#chatlocation.
type ChatLocation struct {
	Location Location `json:"location"`
	Address  string   `json:"address"`
}

// BotCommand is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#botcommand.
type BotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

// ResponseParameters is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#responseparameters.
type ResponseParameters struct {
	MigrateToChatID int `json:"migrate_to_chat_id"`
	RetryAfter      int `json:"retry_after"`
}

// InputMedia is a group containing the InputMedia types
// <https://core.telegram.org/bots/api#inputmedia>
type InputMedia struct {
	InputMediaAnimation *InputMediaAnimation
	InputMediaDocument  *InputMediaDocument
	InputMediaAudio     *InputMediaAudio
	InputMediaPhoto     *InputMediaPhoto
	InputMediaVideo     *InputMediaVideo
}

// InputMediaPhoto is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#inputmediaphoto.
type InputMediaPhoto struct {
	Type            string           `json:"type"`
	Media           string           `json:"media"`
	Caption         string           `json:"caption"`
	ParseMode       string           `json:"parse_mode"`
	CaptionEntities []*MessageEntity `json:"caption_entities"`
}

// InputMediaVideo is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#inputmediavideo.
type InputMediaVideo struct {
	Type              string           `json:"type"`
	Media             string           `json:"media"`
	Thumb             string           `json:"thumb"`
	Caption           string           `json:"caption"`
	ParseMode         string           `json:"parse_mode"`
	CaptionEntities   []*MessageEntity `json:"caption_entities"`
	Width             int              `json:"width"`
	Height            int              `json:"height"`
	Duration          int              `json:"duration"`
	SupportsStreaming bool             `json:"supports_streaming"`
}

// InputMediaAnimation is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#inputmediaanimation.
type InputMediaAnimation struct {
	Type              string           `json:"type"`
	Media             string           `json:"media"`
	Thumb             string           `json:"thumb"`
	Caption           string           `json:"caption"`
	ParseMode         string           `json:"parse_mode"`
	CaptionEntities   []*MessageEntity `json:"caption_entities"`
	Width             int              `json:"width"`
	Height            int              `json:"height"`
	Duration          int              `json:"duration"`
	SupportsStreaming bool             `json:"supports_streaming"`
}

// InputMediaAudio is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#inputmediaudio.
type InputMediaAudio struct {
	Type            string           `json:"type"`
	Media           string           `json:"media"`
	Thumb           string           `json:"thumb"`
	Caption         string           `json:"caption"`
	ParseMode       string           `json:"parse_mode"`
	CaptionEntities []*MessageEntity `json:"caption_entities"`
	Duration        int              `json:"duration"`
	Performer       string           `json:"performer"`
	Title           string           `json:"title"`
}

// InputMediaDocument is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#inputmediadocument.
type InputMediaDocument struct {
	Type                        string           `json:"type"`
	Media                       string           `json:"media"`
	Thumb                       string           `json:"thumb"`
	Caption                     string           `json:"caption"`
	ParseMode                   string           `json:"parse_mode"`
	CaptionEntities             []*MessageEntity `json:"caption_entities"`
	DisableContentTypeDetection bool             `json:"disable_content_type_detection"`
}

// Sticker is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#sticker.
type Sticker struct {
	FileID       string        `json:"file_id"`
	FileUniqueID string        `json:"file_unique_id"`
	Width        int           `json:"width"`
	Height       int           `json:"height"`
	IsAnimated   bool          `json:"is_animated"`
	Thumb        *PhotoSize    `json:"thumb"`
	Emoji        string        `json:"emoji"`
	SetName      string        `json:"set_name"`
	MaskPosition *MaskPosition `json:"mask_position"`
	FileSize     int           `json:"file_size"`
}

// StickerSet is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#stickerset.
type StickerSet struct {
	Name          string     `json:"name"`
	Title         string     `json:"title"`
	IsAnimated    bool       `json:"is_animated"`
	ContainsMasks bool       `json:"contains_masks"`
	Stickers      []*Sticker `json:"stickers"`
	Thumb         *PhotoSize `json:"thumb"`
}

// MaskPosition is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#maskposition.
type MaskPosition struct {
	Point  string  `json:"point"`
	XShift float32 `json:"x_shift"`
	YShift float32 `json:"y_shift"`
	Scale  float32 `json:"scale"`
}

// TODO: <https://core.telegram.org/bots/api#inlinequery>
// TODO: <https://core.telegram.org/bots/api#inputmessagecontent>
// TODO: <https://core.telegram.org/bots/api#choseninlineresult>

// LabeledPrice is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#labeledprice.
type LabeledPrice struct {
	Label  string `json:"label"`
	Amount int    `json:"amount"`
}

// Invoice is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#invoice.
type Invoice struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	StartParameter string `json:"start_parameter"`
	Currency       string `json:"currency"`
	TotalAmount    int    `json:"total_amount"`
}

// ShippingAddress is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#shippingaddress.
type ShippingAddress struct {
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	City        string `json:"city"`
	StreetLine1 string `json:"street_line1"`
	StreetLine2 string `json:"street_line2"`
	PostCode    string `json:"post_code"`
}

// OrderInfo is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#orderinfo.
type OrderInfo struct {
	Name            string          `json:"name"`
	PhoneNumber     string          `json:"phone_number"`
	Email           string          `json:"email"`
	ShippingAddress ShippingAddress `json:"shipping_address"`
}

// ShippingOption is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#shippingoption.
type ShippingOption struct {
	ID     string          `json:"id"`
	Title  string          `json:"title"`
	Prices []*LabeledPrice `json:"prices"`
}

// SuccessfulPayment is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#succsessfulpayment.
type SuccessfulPayment struct {
	Currency                string     `json:"currency"`
	TotalAmount             int        `json:"total_amount"`
	InvoicePayload          string     `json:"invoice_payload"`
	ShippingOptionID        string     `json:"shipping_option_id"`
	OrderInfo               *OrderInfo `json:"order_info"`
	TelegramPaymentChargeID string     `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeID string     `json:"provider_payment_charge_id"`
}

// ShippingQuery is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#shippingquery.
type ShippingQuery struct {
	ID              string           `json:"id"`
	From            *User            `json:"from"`
	InvoicePayload  string           `json:"invoice_payload"`
	ShippingAddress *ShippingAddress `json:"shipping_address"`
}

// PreCheckoutQuery is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#precheckoutquery.
type PreCheckoutQuery struct {
	ID               string     `json:"id"`
	From             *User      `json:"from"`
	Currency         string     `json:"currency"`
	TotalAmount      int        `json:"total_amount"`
	InvoicePayload   string     `json:"invoice_payload"`
	ShippingOptionID string     `json:"shipping_option_id"`
	OrderInfo        *OrderInfo `json:"order_info"`
}

// TODO: <https://core.telegram.org/bots/api#passportdata>
// TODO: <https://core.telegram.org/bots/api#passportfile>
// TODO: <https://core.telegram.org/bots/api#encryptedpassportelement>
// TODO: <https://core.telegram.org/bots/api#encryptedcredentials>
// TODO: <https://core.telegram.org/bots/api#passportelementerror>

// Game is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#game.
type Game struct {
	Title        string           `json:"title"`
	Description  string           `json:"description"`
	Photo        []*PhotoSize     `json:"photo"`
	Text         string           `json:"text"`
	TextEntities []*MessageEntity `json:"text_entities"`
	Animation    Animation        `json:"animation"`
}

// CallbackGame is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#callbackgame.
type CallbackGame struct {
}

// GameHighScore is the wrapper for the BotAPI's type https://core.telegram.org/bots/api#gamehighscore.
type GameHighScore struct {
	Position int   `json:"position"`
	User     *User `json:"user"`
	Score    int   `json:"score"`
}
