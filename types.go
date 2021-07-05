package gobot

import "encoding/json"

type Body struct {
	Ok                 bool                `json:"ok"`
	Description        string              `json:"description"`
	ErrorCode          int                 `json:"error_code"`
	Result             json.RawMessage     `json:"result"`
	ResponseParameters *ResponseParameters `json:"parameters"`
}

type Error struct {
	Description        string
	ErrorCode          int
	ResponseParameters *ResponseParameters
}

func (err *Error) Error() string {
	return err.Description
}

type Update struct {
	UpdateId           int                 `json:"update_id"`            // The update's unique identifier. Update identifiers start from a certain positive number and increase sequentially. This ID becomes especially handy if you're using Webhooks, since it allows you to ignore repeated updates or to restore the correct update sequence, should they get out of order. If there are no new updates for at least a week, then identifier of the next update will be chosen randomly instead of sequentially.
	Message            *Message            `json:"message"`              // Optional. New incoming message of any kind — text, photo, sticker, etc.
	EditedMessage      *Message            `json:"edited_message"`       // Optional. New version of a message that is known to the bot and was edited
	ChannelPost        *Message            `json:"channel_post"`         // Optional. New incoming channel post of any kind — text, photo, sticker, etc.
	EditedChannelPost  *Message            `json:"edited_channel_post"`  // Optional. New version of a channel post that is known to the bot and was edited
	InlineQuery        *InlineQuery        `json:"inline_query"`         // Optional. New incoming inline query
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result"` // Optional. The result of an inline query that was chosen by a user and sent to their chat partner. Please see our documentation on the feedback collecting for details on how to enable these updates for your bot.
	CallbackQuery      *CallbackQuery      `json:"callback_query"`       // Optional. New incoming callback query
	ShippingQuery      *ShippingQuery      `json:"shipping_query"`       // Optional. New incoming shipping query. Only for invoices with flexible price
	PreCheckoutQuery   *PreCheckoutQuery   `json:"pre_checkout_query"`   // Optional. New incoming pre-checkout query. Contains full information about checkout
	Poll               *Poll               `json:"poll"`                 // Optional. New poll state. Bots receive only updates about stopped polls and polls, which are sent by the bot
	PollAnswer         *PollAnswer         `json:"poll_answer"`          // Optional. A user changed their answer in a non-anonymous poll. Bots receive new votes only in polls that were sent by the bot itself.
	MyChatMember       *ChatMemberUpdated  `json:"my_chat_member"`       // Optional. The bot's chat member status was updated in a chat. For private chats, this update is received only when the bot is blocked or unblocked by the user.
	ChatMember         *ChatMemberUpdated  `json:"chat_member"`          // Optional. A chat member's status was updated in a chat. The bot must be an administrator in the chat and must explicitly specify "chat_member" in the list of allowed_updates to receive these updates.
}

type WebhookInfo struct {
	Url                  string   `json:"url"`                    // Webhook URL, may be empty if webhook is not set up
	HasCustomCertificate bool     `json:"has_custom_certificate"` // True, if a custom certificate was provided for webhook certificate checks
	PendingUpdateCount   int      `json:"pending_update_count"`   // Number of updates awaiting delivery
	IpAddress            string   `json:"ip_address"`             // Optional. Currently used webhook IP address
	LastErrorDate        int      `json:"last_error_date"`        // Optional. Unix time for the most recent error that happened when trying to deliver an update via webhook
	LastErrorMessage     string   `json:"last_error_message"`     // Optional. Error message in human-readable format for the most recent error that happened when trying to deliver an update via webhook
	MaxConnections       int      `json:"max_connections"`        // Optional. Maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery
	AllowedUpdates       []string `json:"allowed_updates"`        // Optional. A list of update types the bot is subscribed to. Defaults to all update types except chat_member
}

type User struct {
	Id                      int    `json:"id"`                          // Unique identifier for this user or bot. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier.
	IsBot                   bool   `json:"is_bot"`                      // True, if this user is a bot
	FirstName               string `json:"first_name"`                  // User's or bot's first name
	LastName                string `json:"last_name"`                   // Optional. User's or bot's last name
	Username                string `json:"username"`                    // Optional. User's or bot's username
	LanguageCode            string `json:"language_code"`               // Optional. IETF language tag of the user's language
	CanJoinGroups           bool   `json:"can_join_groups"`             // Optional. True, if the bot can be invited to groups. Returned only in getMe.
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"` // Optional. True, if privacy mode is disabled for the bot. Returned only in getMe.
	SupportsInlineQueries   bool   `json:"supports_inline_queries"`     // Optional. True, if the bot supports inline queries. Returned only in getMe.
}

type Chat struct {
	Id                    int              `json:"id"`                       // Unique identifier for this chat. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	Type                  string           `json:"type"`                     // Type of chat, can be either "private", "group", "supergroup" or "channel"
	Title                 string           `json:"title"`                    // Optional. Title, for supergroups, channels and group chats
	Username              string           `json:"username"`                 // Optional. Username, for private chats, supergroups and channels if available
	FirstName             string           `json:"first_name"`               // Optional. First name of the other party in a private chat
	LastName              string           `json:"last_name"`                // Optional. Last name of the other party in a private chat
	Photo                 *ChatPhoto       `json:"photo"`                    // Optional. Chat photo. Returned only in getChat.
	Bio                   string           `json:"bio"`                      // Optional. Bio of the other party in a private chat. Returned only in getChat.
	Description           string           `json:"description"`              // Optional. Description, for groups, supergroups and channel chats. Returned only in getChat.
	InviteLink            string           `json:"invite_link"`              // Optional. Primary invite link, for groups, supergroups and channel chats. Returned only in getChat.
	PinnedMessage         *Message         `json:"pinned_message"`           // Optional. The most recent pinned message (by sending date). Returned only in getChat.
	Permissions           *ChatPermissions `json:"permissions"`              // Optional. Default chat member permissions, for groups and supergroups. Returned only in getChat.
	SlowModeDelay         int              `json:"slow_mode_delay"`          // Optional. For supergroups, the minimum allowed delay between consecutive messages sent by each unpriviledged user. Returned only in getChat.
	MessageAutoDeleteTime int              `json:"message_auto_delete_time"` // Optional. The time after which all messages sent to the chat will be automatically deleted; in seconds. Returned only in getChat.
	StickerSetName        string           `json:"sticker_set_name"`         // Optional. For supergroups, name of group sticker set. Returned only in getChat.
	CanSetStickerSet      bool             `json:"can_set_sticker_set"`      // Optional. True, if the bot can change the group sticker set. Returned only in getChat.
	LinkedChatId          int              `json:"linked_chat_id"`           // Optional. Unique identifier for the linked chat, i.e. the discussion group identifier for a channel and vice versa; for supergroups and channel chats. This identifier may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier. Returned only in getChat.
	Location              *ChatLocation    `json:"location"`                 // Optional. For supergroups, the location to which the supergroup is connected. Returned only in getChat.
}

type Message struct {
	MessageId                     int                            `json:"message_id"`                        // Unique message identifier inside this chat
	From                          *User                          `json:"from"`                              // Optional. Sender, empty for messages sent to channels
	SenderChat                    *Chat                          `json:"sender_chat"`                       // Optional. Sender of the message, sent on behalf of a chat. The channel itself for channel messages. The supergroup itself for messages from anonymous group administrators. The linked channel for messages automatically forwarded to the discussion group
	Date                          int                            `json:"date"`                              // Date the message was sent in Unix time
	Chat                          *Chat                          `json:"chat"`                              // Conversation the message belongs to
	ForwardFrom                   *User                          `json:"forward_from"`                      // Optional. For forwarded messages, sender of the original message
	ForwardFromChat               *Chat                          `json:"forward_from_chat"`                 // Optional. For messages forwarded from channels or from anonymous administrators, information about the original sender chat
	ForwardFromMessageId          int                            `json:"forward_from_message_id"`           // Optional. For messages forwarded from channels, identifier of the original message in the channel
	ForwardSignature              string                         `json:"forward_signature"`                 // Optional. For messages forwarded from channels, signature of the post author if present
	ForwardSenderName             string                         `json:"forward_sender_name"`               // Optional. Sender's name for messages forwarded from users who disallow adding a link to their account in forwarded messages
	ForwardDate                   int                            `json:"forward_date"`                      // Optional. For forwarded messages, date the original message was sent in Unix time
	ReplyToMessage                *Message                       `json:"reply_to_message"`                  // Optional. For replies, the original message. Note that the Message object in this field will not contain further reply_to_message fields even if it itself is a reply.
	ViaBot                        *User                          `json:"via_bot"`                           // Optional. Bot through which the message was sent
	EditDate                      int                            `json:"edit_date"`                         // Optional. Date the message was last edited in Unix time
	MediaGroupId                  string                         `json:"media_group_id"`                    // Optional. The unique identifier of a media message group this message belongs to
	AuthorSignature               string                         `json:"author_signature"`                  // Optional. Signature of the post author for messages in channels, or the custom title of an anonymous group administrator
	Text                          string                         `json:"text"`                              // Optional. For text messages, the actual UTF-8 text of the message, 0-4096 characters
	Entities                      []*MessageEntity               `json:"entities"`                          // Optional. For text messages, special entities like usernames, URLs, bot commands, etc. that appear in the text
	Animation                     *Animation                     `json:"animation"`                         // Optional. Message is an animation, information about the animation. For backward compatibility, when this field is set, the document field will also be set
	Audio                         *Audio                         `json:"audio"`                             // Optional. Message is an audio file, information about the file
	Document                      *Document                      `json:"document"`                          // Optional. Message is a general file, information about the file
	Photo                         []*PhotoSize                   `json:"photo"`                             // Optional. Message is a photo, available sizes of the photo
	Sticker                       *Sticker                       `json:"sticker"`                           // Optional. Message is a sticker, information about the sticker
	Video                         *Video                         `json:"video"`                             // Optional. Message is a video, information about the video
	VideoNote                     *VideoNote                     `json:"video_note"`                        // Optional. Message is a video note, information about the video message
	Voice                         *Voice                         `json:"voice"`                             // Optional. Message is a voice message, information about the file
	Caption                       string                         `json:"caption"`                           // Optional. Caption for the animation, audio, document, photo, video or voice, 0-1024 characters
	CaptionEntities               []*MessageEntity               `json:"caption_entities"`                  // Optional. For messages with a caption, special entities like usernames, URLs, bot commands, etc. that appear in the caption
	Contact                       *Contact                       `json:"contact"`                           // Optional. Message is a shared contact, information about the contact
	Dice                          *Dice                          `json:"dice"`                              // Optional. Message is a dice with random value
	Game                          *Game                          `json:"game"`                              // Optional. Message is a game, information about the game. More about games »
	Poll                          *Poll                          `json:"poll"`                              // Optional. Message is a native poll, information about the poll
	Venue                         *Venue                         `json:"venue"`                             // Optional. Message is a venue, information about the venue. For backward compatibility, when this field is set, the location field will also be set
	Location                      *Location                      `json:"location"`                          // Optional. Message is a shared location, information about the location
	NewChatMembers                []*User                        `json:"new_chat_members"`                  // Optional. New members that were added to the group or supergroup and information about them (the bot itself may be one of these members)
	LeftChatMember                *User                          `json:"left_chat_member"`                  // Optional. A member was removed from the group, information about them (this member may be the bot itself)
	NewChatTitle                  string                         `json:"new_chat_title"`                    // Optional. A chat title was changed to this value
	NewChatPhoto                  []*PhotoSize                   `json:"new_chat_photo"`                    // Optional. A chat photo was change to this value
	DeleteChatPhoto               bool                           `json:"delete_chat_photo"`                 // Optional. Service message: the chat photo was deleted
	GroupChatCreated              bool                           `json:"group_chat_created"`                // Optional. Service message: the group has been created
	SupergroupChatCreated         bool                           `json:"supergroup_chat_created"`           // Optional. Service message: the supergroup has been created. This field can't be received in a message coming through updates, because bot can't be a member of a supergroup when it is created. It can only be found in reply_to_message if someone replies to a very first message in a directly created supergroup.
	ChannelChatCreated            bool                           `json:"channel_chat_created"`              // Optional. Service message: the channel has been created. This field can't be received in a message coming through updates, because bot can't be a member of a channel when it is created. It can only be found in reply_to_message if someone replies to a very first message in a channel.
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed"` // Optional. Service message: auto-delete timer settings changed in the chat
	MigrateToChatId               int                            `json:"migrate_to_chat_id"`                // Optional. The group has been migrated to a supergroup with the specified identifier. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	MigrateFromChatId             int                            `json:"migrate_from_chat_id"`              // Optional. The supergroup has been migrated from a group with the specified identifier. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	PinnedMessage                 *Message                       `json:"pinned_message"`                    // Optional. Specified message was pinned. Note that the Message object in this field will not contain further reply_to_message fields even if it is itself a reply.
	Invoice                       *Invoice                       `json:"invoice"`                           // Optional. Message is an invoice for a payment, information about the invoice. More about payments »
	SuccessfulPayment             *SuccessfulPayment             `json:"successful_payment"`                // Optional. Message is a service message about a successful payment, information about the payment. More about payments »
	ConnectedWebsite              string                         `json:"connected_website"`                 // Optional. The domain name of the website on which the user has logged in. More about Telegram Login »
	PassportData                  *PassportData                  `json:"passport_data"`                     // Optional. Telegram Passport data
	ProximityAlertTriggered       *ProximityAlertTriggered       `json:"proximity_alert_triggered"`         // Optional. Service message. A user in the chat triggered another user's proximity alert while sharing Live Location.
	VoiceChatScheduled            *VoiceChatScheduled            `json:"voice_chat_scheduled"`              // Optional. Service message: voice chat scheduled
	VoiceChatStarted              *VoiceChatStarted              `json:"voice_chat_started"`                // Optional. Service message: voice chat started
	VoiceChatEnded                *VoiceChatEnded                `json:"voice_chat_ended"`                  // Optional. Service message: voice chat ended
	VoiceChatParticipantsInvited  *VoiceChatParticipantsInvited  `json:"voice_chat_participants_invited"`   // Optional. Service message: new participants invited to a voice chat
	ReplyMarkup                   *InlineKeyboardMarkup          `json:"reply_markup"`                      // Optional. Inline keyboard attached to the message. login_url buttons are represented as ordinary url buttons.
}

type MessageId struct {
	MessageId int `json:"message_id"` // Unique message identifier
}

type MessageEntity struct {
	Type     string `json:"type"`     // Type of the entity. Can be "mention" (@username), "hashtag" (#hashtag), "cashtag" ($USD), "bot_command" (/start@jobs_bot), "url" (https://telegram.org), "email" (do-not-reply@telegram.org), "phone_number" (+1-212-555-0123), "bold" (bold text), "italic" (italic text), "underline" (underlined text), "strikethrough" (strikethrough text), "code" (monowidth string), "pre" (monowidth block), "text_link" (for clickable text URLs), "text_mention" (for users without usernames)
	Offset   int    `json:"offset"`   // Offset in UTF-16 code units to the start of the entity
	Length   int    `json:"length"`   // Length of the entity in UTF-16 code units
	Url      string `json:"url"`      // Optional. For "text_link" only, url that will be opened after user taps on the text
	User     *User  `json:"user"`     // Optional. For "text_mention" only, the mentioned user
	Language string `json:"language"` // Optional. For "pre" only, the programming language of the entity text
}

type PhotoSize struct {
	FileId       string `json:"file_id"`        // Identifier for this file, which can be used to download or reuse the file
	FileUniqueId string `json:"file_unique_id"` // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Width        int    `json:"width"`          // Photo width
	Height       int    `json:"height"`         // Photo height
	FileSize     int    `json:"file_size"`      // Optional. File size
}

type Animation struct {
	FileId       string     `json:"file_id"`        // Identifier for this file, which can be used to download or reuse the file
	FileUniqueId string     `json:"file_unique_id"` // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Width        int        `json:"width"`          // Video width as defined by sender
	Height       int        `json:"height"`         // Video height as defined by sender
	Duration     int        `json:"duration"`       // Duration of the video in seconds as defined by sender
	Thumb        *PhotoSize `json:"thumb"`          // Optional. Animation thumbnail as defined by sender
	FileName     string     `json:"file_name"`      // Optional. Original animation filename as defined by sender
	MimeType     string     `json:"mime_type"`      // Optional. MIME type of the file as defined by sender
	FileSize     int        `json:"file_size"`      // Optional. File size
}

type Audio struct {
	FileId       string     `json:"file_id"`        // Identifier for this file, which can be used to download or reuse the file
	FileUniqueId string     `json:"file_unique_id"` // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Duration     int        `json:"duration"`       // Duration of the audio in seconds as defined by sender
	Performer    string     `json:"performer"`      // Optional. Performer of the audio as defined by sender or by audio tags
	Title        string     `json:"title"`          // Optional. Title of the audio as defined by sender or by audio tags
	FileName     string     `json:"file_name"`      // Optional. Original filename as defined by sender
	MimeType     string     `json:"mime_type"`      // Optional. MIME type of the file as defined by sender
	FileSize     int        `json:"file_size"`      // Optional. File size
	Thumb        *PhotoSize `json:"thumb"`          // Optional. Thumbnail of the album cover to which the music file belongs
}

type Document struct {
	FileId       string     `json:"file_id"`        // Identifier for this file, which can be used to download or reuse the file
	FileUniqueId string     `json:"file_unique_id"` // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Thumb        *PhotoSize `json:"thumb"`          // Optional. Document thumbnail as defined by sender
	FileName     string     `json:"file_name"`      // Optional. Original filename as defined by sender
	MimeType     string     `json:"mime_type"`      // Optional. MIME type of the file as defined by sender
	FileSize     int        `json:"file_size"`      // Optional. File size
}

type Video struct {
	FileId       string     `json:"file_id"`        // Identifier for this file, which can be used to download or reuse the file
	FileUniqueId string     `json:"file_unique_id"` // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Width        int        `json:"width"`          // Video width as defined by sender
	Height       int        `json:"height"`         // Video height as defined by sender
	Duration     int        `json:"duration"`       // Duration of the video in seconds as defined by sender
	Thumb        *PhotoSize `json:"thumb"`          // Optional. Video thumbnail
	FileName     string     `json:"file_name"`      // Optional. Original filename as defined by sender
	MimeType     string     `json:"mime_type"`      // Optional. Mime type of a file as defined by sender
	FileSize     int        `json:"file_size"`      // Optional. File size
}

type VideoNote struct {
	FileId       string     `json:"file_id"`        // Identifier for this file, which can be used to download or reuse the file
	FileUniqueId string     `json:"file_unique_id"` // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Length       int        `json:"length"`         // Video width and height (diameter of the video message) as defined by sender
	Duration     int        `json:"duration"`       // Duration of the video in seconds as defined by sender
	Thumb        *PhotoSize `json:"thumb"`          // Optional. Video thumbnail
	FileSize     int        `json:"file_size"`      // Optional. File size
}

type Voice struct {
	FileId       string `json:"file_id"`        // Identifier for this file, which can be used to download or reuse the file
	FileUniqueId string `json:"file_unique_id"` // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Duration     int    `json:"duration"`       // Duration of the audio in seconds as defined by sender
	MimeType     string `json:"mime_type"`      // Optional. MIME type of the file as defined by sender
	FileSize     int    `json:"file_size"`      // Optional. File size
}

type Contact struct {
	PhoneNumber string `json:"phone_number"` // Contact's phone number
	FirstName   string `json:"first_name"`   // Contact's first name
	LastName    string `json:"last_name"`    // Optional. Contact's last name
	UserId      int    `json:"user_id"`      // Optional. Contact's user identifier in Telegram. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier.
	Vcard       string `json:"vcard"`        // Optional. Additional data about the contact in the form of a vCard
}

type Dice struct {
	Emoji string `json:"emoji"` // Emoji on which the dice throw animation is based
	Value int    `json:"value"` // Value of the dice, 1-6 for "", "" and "" base emoji, 1-5 for "" and "" base emoji, 1-64 for "" base emoji
}

type PollOption struct {
	Text       string `json:"text"`        // Option text, 1-100 characters
	VoterCount int    `json:"voter_count"` // Number of users that voted for this option
}

type PollAnswer struct {
	PollId    string `json:"poll_id"`    // Unique poll identifier
	User      *User  `json:"user"`       // The user, who changed the answer to the poll
	OptionIds []int  `json:"option_ids"` // 0-based identifiers of answer options, chosen by the user. May be empty if the user retracted their vote.
}

type Poll struct {
	Id                    string           `json:"id"`                      // Unique poll identifier
	Question              string           `json:"question"`                // Poll question, 1-300 characters
	Options               []*PollOption    `json:"options"`                 // List of poll options
	TotalVoterCount       int              `json:"total_voter_count"`       // Total number of users that voted in the poll
	IsClosed              bool             `json:"is_closed"`               // True, if the poll is closed
	IsAnonymous           bool             `json:"is_anonymous"`            // True, if the poll is anonymous
	Type                  string           `json:"type"`                    // Poll type, currently can be "regular" or "quiz"
	AllowsMultipleAnswers bool             `json:"allows_multiple_answers"` // True, if the poll allows multiple answers
	CorrectOptionId       int              `json:"correct_option_id"`       // Optional. 0-based identifier of the correct answer option. Available only for polls in the quiz mode, which are closed, or was sent (not forwarded) by the bot or to the private chat with the bot.
	Explanation           string           `json:"explanation"`             // Optional. Text that is shown when a user chooses an incorrect answer or taps on the lamp icon in a quiz-style poll, 0-200 characters
	ExplanationEntities   []*MessageEntity `json:"explanation_entities"`    // Optional. Special entities like usernames, URLs, bot commands, etc. that appear in the explanation
	OpenPeriod            int              `json:"open_period"`             // Optional. Amount of time in seconds the poll will be active after creation
	CloseDate             int              `json:"close_date"`              // Optional. Point in time (Unix timestamp) when the poll will be automatically closed
}

type Location struct {
	Longitude            float64 `json:"longitude"`              // Longitude as defined by sender
	Latitude             float64 `json:"latitude"`               // Latitude as defined by sender
	HorizontalAccuracy   float64 `json:"horizontal_accuracy"`    // Optional. The radius of uncertainty for the location, measured in meters; 0-1500
	LivePeriod           int     `json:"live_period"`            // Optional. Time relative to the message sending date, during which the location can be updated, in seconds. For active live locations only.
	Heading              int     `json:"heading"`                // Optional. The direction in which user is moving, in degrees; 1-360. For active live locations only.
	ProximityAlertRadius int     `json:"proximity_alert_radius"` // Optional. Maximum distance for proximity alerts about approaching another chat member, in meters. For sent live locations only.
}

type Venue struct {
	Location        *Location `json:"location"`          // Venue location. Can't be a live location
	Title           string    `json:"title"`             // Name of the venue
	Address         string    `json:"address"`           // Address of the venue
	FoursquareId    string    `json:"foursquare_id"`     // Optional. Foursquare identifier of the venue
	FoursquareType  string    `json:"foursquare_type"`   // Optional. Foursquare type of the venue. (For example, "arts_entertainment/default", "arts_entertainment/aquarium" or "food/icecream".)
	GooglePlaceId   string    `json:"google_place_id"`   // Optional. Google Places identifier of the venue
	GooglePlaceType string    `json:"google_place_type"` // Optional. Google Places type of the venue. (See supported types.)
}

type ProximityAlertTriggered struct {
	Traveler *User `json:"traveler"` // User that triggered the alert
	Watcher  *User `json:"watcher"`  // User that set the alert
	Distance int   `json:"distance"` // The distance between the users
}

type MessageAutoDeleteTimerChanged struct {
	MessageAutoDeleteTime int `json:"message_auto_delete_time"` // New auto-delete time for messages in the chat
}

type VoiceChatScheduled struct {
	StartDate int `json:"start_date"` // Point in time (Unix timestamp) when the voice chat is supposed to be started by a chat administrator
}

type VoiceChatStarted struct {
}

type VoiceChatEnded struct {
	Duration int `json:"duration"` // Voice chat duration; in seconds
}

type VoiceChatParticipantsInvited struct {
	Users []*User `json:"users"` // Optional. New members that were invited to the voice chat
}

type UserProfilePhotos struct {
	TotalCount int            `json:"total_count"` // Total number of profile pictures the target user has
	Photos     [][]*PhotoSize `json:"photos"`      // Requested profile pictures (in up to 4 sizes each)
}

type File struct {
	FileId       string `json:"file_id"`        // Identifier for this file, which can be used to download or reuse the file
	FileUniqueId string `json:"file_unique_id"` // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	FileSize     int    `json:"file_size"`      // Optional. File size, if known
	FilePath     string `json:"file_path"`      // Optional. File path. Use https://api.telegram.org/file/bot<token>/<file_path> to get the file.
}

type ReplyKeyboardMarkup struct {
	Keyboard              [][]*KeyboardButton `json:"keyboard"`                // Array of button rows, each represented by an Array of KeyboardButton objects
	ResizeKeyboard        bool                `json:"resize_keyboard"`         // Optional. Requests clients to resize the keyboard vertically for optimal fit (e.g., make the keyboard smaller if there are just two rows of buttons). Defaults to false, in which case the custom keyboard is always of the same height as the app's standard keyboard.
	OneTimeKeyboard       bool                `json:"one_time_keyboard"`       // Optional. Requests clients to hide the keyboard as soon as it's been used. The keyboard will still be available, but clients will automatically display the usual letter-keyboard in the chat – the user can press a special button in the input field to see the custom keyboard again. Defaults to false.
	InputFieldPlaceholder string              `json:"input_field_placeholder"` // Optional. The placeholder to be shown in the input field when the keyboard is active; 1-64 characters
	Selective             bool                `json:"selective"`               // Optional. Use this parameter if you want to show the keyboard to specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.Example: A user requests to change the bot's language, bot replies to the request with a keyboard to select the new language. Other users in the group don't see the keyboard.
}

type KeyboardButton struct {
	Text            string                  `json:"text"`             // Text of the button. If none of the optional fields are used, it will be sent as a message when the button is pressed
	RequestContact  bool                    `json:"request_contact"`  // Optional. If True, the user's phone number will be sent as a contact when the button is pressed. Available in private chats only
	RequestLocation bool                    `json:"request_location"` // Optional. If True, the user's current location will be sent when the button is pressed. Available in private chats only
	RequestPoll     *KeyboardButtonPollType `json:"request_poll"`     // Optional. If specified, the user will be asked to create a poll and send it to the bot when the button is pressed. Available in private chats only
}

type KeyboardButtonPollType struct {
	Type string `json:"type"` // Optional. If quiz is passed, the user will be allowed to create only polls in the quiz mode. If regular is passed, only regular polls will be allowed. Otherwise, the user will be allowed to create a poll of any type.
}

type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"` // Requests clients to remove the custom keyboard (user will not be able to summon this keyboard; if you want to hide the keyboard from sight but keep it accessible, use one_time_keyboard in ReplyKeyboardMarkup)
	Selective      bool `json:"selective"`       // Optional. Use this parameter if you want to remove the keyboard for specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.Example: A user votes in a poll, bot returns confirmation message in reply to the vote and removes the keyboard for that user, while still showing the keyboard with poll options to users who haven't voted yet.
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]*InlineKeyboardButton `json:"inline_keyboard"` // Array of button rows, each represented by an Array of InlineKeyboardButton objects
}

type InlineKeyboardButton struct {
	Text                         string        `json:"text"`                             // Label text on the button
	Url                          string        `json:"url"`                              // Optional. HTTP or tg:// url to be opened when button is pressed
	LoginUrl                     *LoginUrl     `json:"login_url"`                        // Optional. An HTTP URL used to automatically authorize the user. Can be used as a replacement for the Telegram Login Widget.
	CallbackData                 string        `json:"callback_data"`                    // Optional. Data to be sent in a callback query to the bot when button is pressed, 1-64 bytes
	SwitchInlineQuery            string        `json:"switch_inline_query"`              // Optional. If set, pressing the button will prompt the user to select one of their chats, open that chat and insert the bot's username and the specified inline query in the input field. Can be empty, in which case just the bot's username will be inserted.Note: This offers an easy way for users to start using your bot in inline mode when they are currently in a private chat with it. Especially useful when combined with switch_pm… actions – in this case the user will be automatically returned to the chat they switched from, skipping the chat selection screen.
	SwitchInlineQueryCurrentChat string        `json:"switch_inline_query_current_chat"` // Optional. If set, pressing the button will insert the bot's username and the specified inline query in the current chat's input field. Can be empty, in which case only the bot's username will be inserted.This offers a quick way for the user to open your bot in inline mode in the same chat – good for selecting something from multiple options.
	CallbackGame                 *CallbackGame `json:"callback_game"`                    // Optional. Description of the game that will be launched when the user presses the button.NOTE: This type of button must always be the first button in the first row.
	Pay                          bool          `json:"pay"`                              // Optional. Specify True, to send a Pay button.NOTE: This type of button must always be the first button in the first row.
}

type LoginUrl struct {
	Url                string `json:"url"`                  // An HTTP URL to be opened with user authorization data added to the query string when the button is pressed. If the user refuses to provide authorization data, the original URL without information about the user will be opened. The data added is the same as described in Receiving authorization data.NOTE: You must always check the hash of the received data to verify the authentication and the integrity of the data as described in Checking authorization.
	ForwardText        string `json:"forward_text"`         // Optional. New text of the button in forwarded messages.
	BotUsername        string `json:"bot_username"`         // Optional. Username of a bot, which will be used for user authorization. See Setting up a bot for more details. If not specified, the current bot's username will be assumed. The url's domain must be the same as the domain linked with the bot. See Linking your domain to the bot for more details.
	RequestWriteAccess bool   `json:"request_write_access"` // Optional. Pass True to request the permission for your bot to send messages to the user.
}

type CallbackQuery struct {
	Id              string   `json:"id"`                // Unique identifier for this query
	From            *User    `json:"from"`              // Sender
	Message         *Message `json:"message"`           // Optional. Message with the callback button that originated the query. Note that message content and message date will not be available if the message is too old
	InlineMessageId string   `json:"inline_message_id"` // Optional. Identifier of the message sent via the bot in inline mode, that originated the query.
	ChatInstance    string   `json:"chat_instance"`     // Global identifier, uniquely corresponding to the chat to which the message with the callback button was sent. Useful for high scores in games.
	Data            string   `json:"data"`              // Optional. Data associated with the callback button. Be aware that a bad client can send arbitrary data in this field.
	GameShortName   string   `json:"game_short_name"`   // Optional. Short name of a Game to be returned, serves as the unique identifier for the game
}

type ForceReply struct {
	ForceReply            bool   `json:"force_reply"`             // Shows reply interface to the user, as if they manually selected the bot's message and tapped 'Reply'
	InputFieldPlaceholder string `json:"input_field_placeholder"` // Optional. The placeholder to be shown in the input field when the reply is active; 1-64 characters
	Selective             bool   `json:"selective"`               // Optional. Use this parameter if you want to force reply from specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
}

type ChatPhoto struct {
	SmallFileId       string `json:"small_file_id"`        // File identifier of small (160x160) chat photo. This file_id can be used only for photo download and only for as long as the photo is not changed.
	SmallFileUniqueId string `json:"small_file_unique_id"` // Unique file identifier of small (160x160) chat photo, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	BigFileId         string `json:"big_file_id"`          // File identifier of big (640x640) chat photo. This file_id can be used only for photo download and only for as long as the photo is not changed.
	BigFileUniqueId   string `json:"big_file_unique_id"`   // Unique file identifier of big (640x640) chat photo, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
}

type ChatInviteLink struct {
	InviteLink  string `json:"invite_link"`  // The invite link. If the link was created by another chat administrator, then the second part of the link will be replaced with "…".
	Creator     *User  `json:"creator"`      // Creator of the link
	IsPrimary   bool   `json:"is_primary"`   // True, if the link is primary
	IsRevoked   bool   `json:"is_revoked"`   // True, if the link is revoked
	ExpireDate  int    `json:"expire_date"`  // Optional. Point in time (Unix timestamp) when the link will expire or has been expired
	MemberLimit int    `json:"member_limit"` // Optional. Maximum number of users that can be members of the chat simultaneously after joining the chat via this invite link; 1-99999
}

type ChatMemberOwner struct {
	Status      string `json:"status"`       // The member's status in the chat, always "creator"
	User        *User  `json:"user"`         // Information about the user
	CustomTitle string `json:"custom_title"` // Custom title for this user
	IsAnonymous bool   `json:"is_anonymous"` // True, if the user's presence in the chat is hidden
}

type ChatMemberAdministrator struct {
	Status              string `json:"status"`                 // The member's status in the chat, always "administrator"
	User                *User  `json:"user"`                   // Information about the user
	CanBeEdited         bool   `json:"can_be_edited"`          // True, if the bot is allowed to edit administrator privileges of that user
	CustomTitle         string `json:"custom_title"`           // Custom title for this user
	IsAnonymous         bool   `json:"is_anonymous"`           // True, if the user's presence in the chat is hidden
	CanManageChat       bool   `json:"can_manage_chat"`        // True, if the administrator can access the chat event log, chat statistics, message statistics in channels, see channel members, see anonymous administrators in supergroups and ignore slow mode. Implied by any other administrator privilege
	CanPostMessages     bool   `json:"can_post_messages"`      // True, if the administrator can post in the channel; channels only
	CanEditMessages     bool   `json:"can_edit_messages"`      // True, if the administrator can edit messages of other users and can pin messages; channels only
	CanDeleteMessages   bool   `json:"can_delete_messages"`    // True, if the administrator can delete messages of other users
	CanManageVoiceChats bool   `json:"can_manage_voice_chats"` // True, if the administrator can manage voice chats
	CanRestrictMembers  bool   `json:"can_restrict_members"`   // True, if the administrator can restrict, ban or unban chat members
	CanPromoteMembers   bool   `json:"can_promote_members"`    // True, if the administrator can add new administrators with a subset of their own privileges or demote administrators that he has promoted, directly or indirectly (promoted by administrators that were appointed by the user)
	CanChangeInfo       bool   `json:"can_change_info"`        // True, if the user is allowed to change the chat title, photo and other settings
	CanInviteUsers      bool   `json:"can_invite_users"`       // True, if the user is allowed to invite new users to the chat
	CanPinMessages      bool   `json:"can_pin_messages"`       // True, if the user is allowed to pin messages; groups and supergroups only
}

type ChatMemberMember struct {
	Status string `json:"status"` // The member's status in the chat, always "member"
	User   *User  `json:"user"`   // Information about the user
}

type ChatMemberRestricted struct {
	Status                string `json:"status"`                    // The member's status in the chat, always "restricted"
	User                  *User  `json:"user"`                      // Information about the user
	IsMember              bool   `json:"is_member"`                 // True, if the user is a member of the chat at the moment of the request
	CanChangeInfo         bool   `json:"can_change_info"`           // True, if the user is allowed to change the chat title, photo and other settings
	CanInviteUsers        bool   `json:"can_invite_users"`          // True, if the user is allowed to invite new users to the chat
	CanPinMessages        bool   `json:"can_pin_messages"`          // True, if the user is allowed to pin messages; groups and supergroups only
	CanSendMessages       bool   `json:"can_send_messages"`         // True, if the user is allowed to send text messages, contacts, locations and venues
	CanSendMediaMessages  bool   `json:"can_send_media_messages"`   // True, if the user is allowed to send audios, documents, photos, videos, video notes and voice notes
	CanSendPolls          bool   `json:"can_send_polls"`            // True, if the user is allowed to send polls
	CanSendOtherMessages  bool   `json:"can_send_other_messages"`   // True, if the user is allowed to send animations, games, stickers and use inline bots
	CanAddWebPagePreviews bool   `json:"can_add_web_page_previews"` // True, if the user is allowed to add web page previews to their messages
	UntilDate             int    `json:"until_date"`                // Date when restrictions will be lifted for this user; unix time
}

type ChatMemberLeft struct {
	Status string `json:"status"` // The member's status in the chat, always "left"
	User   *User  `json:"user"`   // Information about the user
}

type ChatMemberBanned struct {
	Status    string `json:"status"`     // The member's status in the chat, always "kicked"
	User      *User  `json:"user"`       // Information about the user
	UntilDate int    `json:"until_date"` // Date when restrictions will be lifted for this user; unix time
}

type ChatMemberUpdated struct {
	Chat          *Chat           `json:"chat"`            // Chat the user belongs to
	From          *User           `json:"from"`            // Performer of the action, which resulted in the change
	Date          int             `json:"date"`            // Date the change was done in Unix time
	OldChatMember interface{}     `json:"old_chat_member"` // Previous information about the chat member
	NewChatMember interface{}     `json:"new_chat_member"` // New information about the chat member
	InviteLink    *ChatInviteLink `json:"invite_link"`     // Optional. Chat invite link, which was used by the user to join the chat; for joining by invite link events only.
}

type ChatPermissions struct {
	CanSendMessages       bool `json:"can_send_messages"`         // Optional. True, if the user is allowed to send text messages, contacts, locations and venues
	CanSendMediaMessages  bool `json:"can_send_media_messages"`   // Optional. True, if the user is allowed to send audios, documents, photos, videos, video notes and voice notes, implies can_send_messages
	CanSendPolls          bool `json:"can_send_polls"`            // Optional. True, if the user is allowed to send polls, implies can_send_messages
	CanSendOtherMessages  bool `json:"can_send_other_messages"`   // Optional. True, if the user is allowed to send animations, games, stickers and use inline bots, implies can_send_media_messages
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews"` // Optional. True, if the user is allowed to add web page previews to their messages, implies can_send_media_messages
	CanChangeInfo         bool `json:"can_change_info"`           // Optional. True, if the user is allowed to change the chat title, photo and other settings. Ignored in public supergroups
	CanInviteUsers        bool `json:"can_invite_users"`          // Optional. True, if the user is allowed to invite new users to the chat
	CanPinMessages        bool `json:"can_pin_messages"`          // Optional. True, if the user is allowed to pin messages. Ignored in public supergroups
}

type ChatLocation struct {
	Location *Location `json:"location"` // The location to which the supergroup is connected. Can't be a live location.
	Address  string    `json:"address"`  // Location address; 1-64 characters, as defined by the chat owner
}

type BotCommand struct {
	Command     string `json:"command"`     // Text of the command, 1-32 characters. Can contain only lowercase English letters, digits and underscores.
	Description string `json:"description"` // Description of the command, 3-256 characters.
}

type BotCommandScopeDefault struct {
	Type string `json:"type"` // Scope type, must be default
}

type BotCommandScopeAllPrivateChats struct {
	Type string `json:"type"` // Scope type, must be all_private_chats
}

type BotCommandScopeAllGroupChats struct {
	Type string `json:"type"` // Scope type, must be all_group_chats
}

type BotCommandScopeAllChatAdministrators struct {
	Type string `json:"type"` // Scope type, must be all_chat_administrators
}

type BotCommandScopeChat struct {
	Type   string      `json:"type"`    // Scope type, must be chat
	ChatId interface{} `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
}

type BotCommandScopeChatAdministrators struct {
	Type   string      `json:"type"`    // Scope type, must be chat_administrators
	ChatId interface{} `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
}

type BotCommandScopeChatMember struct {
	Type   string      `json:"type"`    // Scope type, must be chat_member
	ChatId interface{} `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	UserId int         `json:"user_id"` // Unique identifier of the target user
}

type ResponseParameters struct {
	MigrateToChatId int `json:"migrate_to_chat_id"` // Optional. The group has been migrated to a supergroup with the specified identifier. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	RetryAfter      int `json:"retry_after"`        // Optional. In case of exceeding flood control, the number of seconds left to wait before the request can be repeated
}

type InputMediaPhoto struct {
	Type            string           `json:"type"`             // Type of the result, must be photo
	Media           string           `json:"media"`            // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass "attach://<file_attach_name>" to upload a new one using multipart/form-data under <file_attach_name> name. More info on Sending Files »
	Caption         string           `json:"caption"`          // Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	ParseMode       string           `json:"parse_mode"`       // Optional. Mode for parsing entities in the photo caption. See formatting options for more details.
	CaptionEntities []*MessageEntity `json:"caption_entities"` // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
}

type InputMediaVideo struct {
	Type              string           `json:"type"`               // Type of the result, must be video
	Media             string           `json:"media"`              // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass "attach://<file_attach_name>" to upload a new one using multipart/form-data under <file_attach_name> name. More info on Sending Files »
	Thumb             interface{}      `json:"thumb"`              // Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass "attach://<file_attach_name>" if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
	Caption           string           `json:"caption"`            // Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
	ParseMode         string           `json:"parse_mode"`         // Optional. Mode for parsing entities in the video caption. See formatting options for more details.
	CaptionEntities   []*MessageEntity `json:"caption_entities"`   // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	Width             int              `json:"width"`              // Optional. Video width
	Height            int              `json:"height"`             // Optional. Video height
	Duration          int              `json:"duration"`           // Optional. Video duration
	SupportsStreaming bool             `json:"supports_streaming"` // Optional. Pass True, if the uploaded video is suitable for streaming
}

type InputMediaAnimation struct {
	Type            string           `json:"type"`             // Type of the result, must be animation
	Media           string           `json:"media"`            // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass "attach://<file_attach_name>" to upload a new one using multipart/form-data under <file_attach_name> name. More info on Sending Files »
	Thumb           interface{}      `json:"thumb"`            // Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass "attach://<file_attach_name>" if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
	Caption         string           `json:"caption"`          // Optional. Caption of the animation to be sent, 0-1024 characters after entities parsing
	ParseMode       string           `json:"parse_mode"`       // Optional. Mode for parsing entities in the animation caption. See formatting options for more details.
	CaptionEntities []*MessageEntity `json:"caption_entities"` // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	Width           int              `json:"width"`            // Optional. Animation width
	Height          int              `json:"height"`           // Optional. Animation height
	Duration        int              `json:"duration"`         // Optional. Animation duration
}

type InputMediaAudio struct {
	Type            string           `json:"type"`             // Type of the result, must be audio
	Media           string           `json:"media"`            // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass "attach://<file_attach_name>" to upload a new one using multipart/form-data under <file_attach_name> name. More info on Sending Files »
	Thumb           interface{}      `json:"thumb"`            // Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass "attach://<file_attach_name>" if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
	Caption         string           `json:"caption"`          // Optional. Caption of the audio to be sent, 0-1024 characters after entities parsing
	ParseMode       string           `json:"parse_mode"`       // Optional. Mode for parsing entities in the audio caption. See formatting options for more details.
	CaptionEntities []*MessageEntity `json:"caption_entities"` // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	Duration        int              `json:"duration"`         // Optional. Duration of the audio in seconds
	Performer       string           `json:"performer"`        // Optional. Performer of the audio
	Title           string           `json:"title"`            // Optional. Title of the audio
}

type InputMediaDocument struct {
	Type                        string           `json:"type"`                           // Type of the result, must be document
	Media                       string           `json:"media"`                          // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass "attach://<file_attach_name>" to upload a new one using multipart/form-data under <file_attach_name> name. More info on Sending Files »
	Thumb                       interface{}      `json:"thumb"`                          // Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass "attach://<file_attach_name>" if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
	Caption                     string           `json:"caption"`                        // Optional. Caption of the document to be sent, 0-1024 characters after entities parsing
	ParseMode                   string           `json:"parse_mode"`                     // Optional. Mode for parsing entities in the document caption. See formatting options for more details.
	CaptionEntities             []*MessageEntity `json:"caption_entities"`               // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	DisableContentTypeDetection bool             `json:"disable_content_type_detection"` // Optional. Disables automatic server-side content type detection for files uploaded using multipart/form-data. Always true, if the document is sent as part of an album.
}

type Sticker struct {
	FileId       string        `json:"file_id"`        // Identifier for this file, which can be used to download or reuse the file
	FileUniqueId string        `json:"file_unique_id"` // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Width        int           `json:"width"`          // Sticker width
	Height       int           `json:"height"`         // Sticker height
	IsAnimated   bool          `json:"is_animated"`    // True, if the sticker is animated
	Thumb        *PhotoSize    `json:"thumb"`          // Optional. Sticker thumbnail in the .WEBP or .JPG format
	Emoji        string        `json:"emoji"`          // Optional. Emoji associated with the sticker
	SetName      string        `json:"set_name"`       // Optional. Name of the sticker set to which the sticker belongs
	MaskPosition *MaskPosition `json:"mask_position"`  // Optional. For mask stickers, the position where the mask should be placed
	FileSize     int           `json:"file_size"`      // Optional. File size
}

type StickerSet struct {
	Name          string     `json:"name"`           // Sticker set name
	Title         string     `json:"title"`          // Sticker set title
	IsAnimated    bool       `json:"is_animated"`    // True, if the sticker set contains animated stickers
	ContainsMasks bool       `json:"contains_masks"` // True, if the sticker set contains masks
	Stickers      []*Sticker `json:"stickers"`       // List of all set stickers
	Thumb         *PhotoSize `json:"thumb"`          // Optional. Sticker set thumbnail in the .WEBP or .TGS format
}

type MaskPosition struct {
	Point  string  `json:"point"`   // The part of the face relative to which the mask should be placed. One of "forehead", "eyes", "mouth", or "chin".
	XShift float64 `json:"x_shift"` // Shift by X-axis measured in widths of the mask scaled to the face size, from left to right. For example, choosing -1.0 will place mask just to the left of the default mask position.
	YShift float64 `json:"y_shift"` // Shift by Y-axis measured in heights of the mask scaled to the face size, from top to bottom. For example, 1.0 will place the mask just below the default mask position.
	Scale  float64 `json:"scale"`   // Mask scaling coefficient. For example, 2.0 means double size.
}

type InlineQuery struct {
	Id       string    `json:"id"`        // Unique identifier for this query
	From     *User     `json:"from"`      // Sender
	Query    string    `json:"query"`     // Text of the query (up to 256 characters)
	Offset   string    `json:"offset"`    // Offset of the results to be returned, can be controlled by the bot
	ChatType string    `json:"chat_type"` // Optional. Type of the chat, from which the inline query was sent. Can be either "sender" for a private chat with the inline query sender, "private", "group", "supergroup", or "channel". The chat type should be always known for requests sent from official clients and most third-party clients, unless the request was sent from a secret chat
	Location *Location `json:"location"`  // Optional. Sender location, only for bots that request user location
}

type InlineQueryResultArticle struct {
	Type                string                `json:"type"`                  // Type of the result, must be article
	Id                  string                `json:"id"`                    // Unique identifier for this result, 1-64 Bytes
	Title               string                `json:"title"`                 // Title of the result
	InputMessageContent interface{}           `json:"input_message_content"` // Content of the message to be sent
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`          // Optional. Inline keyboard attached to the message
	Url                 string                `json:"url"`                   // Optional. URL of the result
	HideUrl             bool                  `json:"hide_url"`              // Optional. Pass True, if you don't want the URL to be shown in the message
	Description         string                `json:"description"`           // Optional. Short description of the result
	ThumbUrl            string                `json:"thumb_url"`             // Optional. Url of the thumbnail for the result
	ThumbWidth          int                   `json:"thumb_width"`           // Optional. Thumbnail width
	ThumbHeight         int                   `json:"thumb_height"`          // Optional. Thumbnail height
}

type InlineQueryResultPhoto struct {
	Type                string                `json:"type"`                  // Type of the result, must be photo
	Id                  string                `json:"id"`                    // Unique identifier for this result, 1-64 bytes
	PhotoUrl            string                `json:"photo_url"`             // A valid URL of the photo. Photo must be in jpeg format. Photo size must not exceed 5MB
	ThumbUrl            string                `json:"thumb_url"`             // URL of the thumbnail for the photo
	PhotoWidth          int                   `json:"photo_width"`           // Optional. Width of the photo
	PhotoHeight         int                   `json:"photo_height"`          // Optional. Height of the photo
	Title               string                `json:"title"`                 // Optional. Title for the result
	Description         string                `json:"description"`           // Optional. Short description of the result
	Caption             string                `json:"caption"`               // Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	ParseMode           string                `json:"parse_mode"`            // Optional. Mode for parsing entities in the photo caption. See formatting options for more details.
	CaptionEntities     []*MessageEntity      `json:"caption_entities"`      // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`          // Optional. Inline keyboard attached to the message
	InputMessageContent interface{}           `json:"input_message_content"` // Optional. Content of the message to be sent instead of the photo
}

type InlineQueryResultGif struct {
	Type                string                `json:"type"`                  // Type of the result, must be gif
	Id                  string                `json:"id"`                    // Unique identifier for this result, 1-64 bytes
	GifUrl              string                `json:"gif_url"`               // A valid URL for the GIF file. File size must not exceed 1MB
	GifWidth            int                   `json:"gif_width"`             // Optional. Width of the GIF
	GifHeight           int                   `json:"gif_height"`            // Optional. Height of the GIF
	GifDuration         int                   `json:"gif_duration"`          // Optional. Duration of the GIF
	ThumbUrl            string                `json:"thumb_url"`             // URL of the static (JPEG or GIF) or animated (MPEG4) thumbnail for the result
	ThumbMimeType       string                `json:"thumb_mime_type"`       // Optional. MIME type of the thumbnail, must be one of "image/jpeg", "image/gif", or "video/mp4". Defaults to "image/jpeg"
	Title               string                `json:"title"`                 // Optional. Title for the result
	Caption             string                `json:"caption"`               // Optional. Caption of the GIF file to be sent, 0-1024 characters after entities parsing
	ParseMode           string                `json:"parse_mode"`            // Optional. Mode for parsing entities in the caption. See formatting options for more details.
	CaptionEntities     []*MessageEntity      `json:"caption_entities"`      // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`          // Optional. Inline keyboard attached to the message
	InputMessageContent interface{}           `json:"input_message_content"` // Optional. Content of the message to be sent instead of the GIF animation
}

type InlineQueryResultMpeg4Gif struct {
	Type                string                `json:"type"`                  // Type of the result, must be mpeg4_gif
	Id                  string                `json:"id"`                    // Unique identifier for this result, 1-64 bytes
	Mpeg4Url            string                `json:"mpeg4_url"`             // A valid URL for the MP4 file. File size must not exceed 1MB
	Mpeg4Width          int                   `json:"mpeg4_width"`           // Optional. Video width
	Mpeg4Height         int                   `json:"mpeg4_height"`          // Optional. Video height
	Mpeg4Duration       int                   `json:"mpeg4_duration"`        // Optional. Video duration
	ThumbUrl            string                `json:"thumb_url"`             // URL of the static (JPEG or GIF) or animated (MPEG4) thumbnail for the result
	ThumbMimeType       string                `json:"thumb_mime_type"`       // Optional. MIME type of the thumbnail, must be one of "image/jpeg", "image/gif", or "video/mp4". Defaults to "image/jpeg"
	Title               string                `json:"title"`                 // Optional. Title for the result
	Caption             string                `json:"caption"`               // Optional. Caption of the MPEG-4 file to be sent, 0-1024 characters after entities parsing
	ParseMode           string                `json:"parse_mode"`            // Optional. Mode for parsing entities in the caption. See formatting options for more details.
	CaptionEntities     []*MessageEntity      `json:"caption_entities"`      // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`          // Optional. Inline keyboard attached to the message
	InputMessageContent interface{}           `json:"input_message_content"` // Optional. Content of the message to be sent instead of the video animation
}

type InlineQueryResultVideo struct {
	Type                string                `json:"type"`                  // Type of the result, must be video
	Id                  string                `json:"id"`                    // Unique identifier for this result, 1-64 bytes
	VideoUrl            string                `json:"video_url"`             // A valid URL for the embedded video player or video file
	MimeType            string                `json:"mime_type"`             // Mime type of the content of video url, "text/html" or "video/mp4"
	ThumbUrl            string                `json:"thumb_url"`             // URL of the thumbnail (jpeg only) for the video
	Title               string                `json:"title"`                 // Title for the result
	Caption             string                `json:"caption"`               // Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
	ParseMode           string                `json:"parse_mode"`            // Optional. Mode for parsing entities in the video caption. See formatting options for more details.
	CaptionEntities     []*MessageEntity      `json:"caption_entities"`      // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	VideoWidth          int                   `json:"video_width"`           // Optional. Video width
	VideoHeight         int                   `json:"video_height"`          // Optional. Video height
	VideoDuration       int                   `json:"video_duration"`        // Optional. Video duration in seconds
	Description         string                `json:"description"`           // Optional. Short description of the result
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`          // Optional. Inline keyboard attached to the message
	InputMessageContent interface{}           `json:"input_message_content"` // Optional. Content of the message to be sent instead of the video. This field is required if InlineQueryResultVideo is used to send an HTML-page as a result (e.g., a YouTube video).
}

type InlineQueryResultAudio struct {
	Type                string                `json:"type"`                  // Type of the result, must be audio
	Id                  string                `json:"id"`                    // Unique identifier for this result, 1-64 bytes
	AudioUrl            string                `json:"audio_url"`             // A valid URL for the audio file
	Title               string                `json:"title"`                 // Title
	Caption             string                `json:"caption"`               // Optional. Caption, 0-1024 characters after entities parsing
	ParseMode           string                `json:"parse_mode"`            // Optional. Mode for parsing entities in the audio caption. See formatting options for more details.
	CaptionEntities     []*MessageEntity      `json:"caption_entities"`      // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	Performer           string                `json:"performer"`             // Optional. Performer
	AudioDuration       int                   `json:"audio_duration"`        // Optional. Audio duration in seconds
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`          // Optional. Inline keyboard attached to the message
	InputMessageContent interface{}           `json:"input_message_content"` // Optional. Content of the message to be sent instead of the audio
}

type InlineQueryResultVoice struct {
	Type                string                `json:"type"`                  // Type of the result, must be voice
	Id                  string                `json:"id"`                    // Unique identifier for this result, 1-64 bytes
	VoiceUrl            string                `json:"voice_url"`             // A valid URL for the voice recording
	Title               string                `json:"title"`                 // Recording title
	Caption             string                `json:"caption"`               // Optional. Caption, 0-1024 characters after entities parsing
	ParseMode           string                `json:"parse_mode"`            // Optional. Mode for parsing entities in the voice message caption. See formatting options for more details.
	CaptionEntities     []*MessageEntity      `json:"caption_entities"`      // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	VoiceDuration       int                   `json:"voice_duration"`        // Optional. Recording duration in seconds
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`          // Optional. Inline keyboard attached to the message
	InputMessageContent interface{}           `json:"input_message_content"` // Optional. Content of the message to be sent instead of the voice recording
}

type InlineQueryResultDocument struct {
	Type                string                `json:"type"`                  // Type of the result, must be document
	Id                  string                `json:"id"`                    // Unique identifier for this result, 1-64 bytes
	Title               string                `json:"title"`                 // Title for the result
	Caption             string                `json:"caption"`               // Optional. Caption of the document to be sent, 0-1024 characters after entities parsing
	ParseMode           string                `json:"parse_mode"`            // Optional. Mode for parsing entities in the document caption. See formatting options for more details.
	CaptionEntities     []*MessageEntity      `json:"caption_entities"`      // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	DocumentUrl         string                `json:"document_url"`          // A valid URL for the file
	MimeType            string                `json:"mime_type"`             // Mime type of the content of the file, either "application/pdf" or "application/zip"
	Description         string                `json:"description"`           // Optional. Short description of the result
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`          // Optional. Inline keyboard attached to the message
	InputMessageContent interface{}           `json:"input_message_content"` // Optional. Content of the message to be sent instead of the file
	ThumbUrl            string                `json:"thumb_url"`             // Optional. URL of the thumbnail (jpeg only) for the file
	ThumbWidth          int                   `json:"thumb_width"`           // Optional. Thumbnail width
	ThumbHeight         int                   `json:"thumb_height"`          // Optional. Thumbnail height
}

type InlineQueryResultLocation struct {
	Type                 string                `json:"type"`                   // Type of the result, must be location
	Id                   string                `json:"id"`                     // Unique identifier for this result, 1-64 Bytes
	Latitude             float64               `json:"latitude"`               // Location latitude in degrees
	Longitude            float64               `json:"longitude"`              // Location longitude in degrees
	Title                string                `json:"title"`                  // Location title
	HorizontalAccuracy   float64               `json:"horizontal_accuracy"`    // Optional. The radius of uncertainty for the location, measured in meters; 0-1500
	LivePeriod           int                   `json:"live_period"`            // Optional. Period in seconds for which the location can be updated, should be between 60 and 86400.
	Heading              int                   `json:"heading"`                // Optional. For live locations, a direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.
	ProximityAlertRadius int                   `json:"proximity_alert_radius"` // Optional. For live locations, a maximum distance for proximity alerts about approaching another chat member, in meters. Must be between 1 and 100000 if specified.
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup"`           // Optional. Inline keyboard attached to the message
	InputMessageContent  interface{}           `json:"input_message_content"`  // Optional. Content of the message to be sent instead of the location
	ThumbUrl             string                `json:"thumb_url"`              // Optional. Url of the thumbnail for the result
	ThumbWidth           int                   `json:"thumb_width"`            // Optional. Thumbnail width
	ThumbHeight          int                   `json:"thumb_height"`           // Optional. Thumbnail height
}

type InlineQueryResultVenue struct {
	Type                string                `json:"type"`                  // Type of the result, must be venue
	Id                  string                `json:"id"`                    // Unique identifier for this result, 1-64 Bytes
	Latitude            float64               `json:"latitude"`              // Latitude of the venue location in degrees
	Longitude           float64               `json:"longitude"`             // Longitude of the venue location in degrees
	Title               string                `json:"title"`                 // Title of the venue
	Address             string                `json:"address"`               // Address of the venue
	FoursquareId        string                `json:"foursquare_id"`         // Optional. Foursquare identifier of the venue if known
	FoursquareType      string                `json:"foursquare_type"`       // Optional. Foursquare type of the venue, if known. (For example, "arts_entertainment/default", "arts_entertainment/aquarium" or "food/icecream".)
	GooglePlaceId       string                `json:"google_place_id"`       // Optional. Google Places identifier of the venue
	GooglePlaceType     string                `json:"google_place_type"`     // Optional. Google Places type of the venue. (See supported types.)
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`          // Optional. Inline keyboard attached to the message
	InputMessageContent interface{}           `json:"input_message_content"` // Optional. Content of the message to be sent instead of the venue
	ThumbUrl            string                `json:"thumb_url"`             // Optional. Url of the thumbnail for the result
	ThumbWidth          int                   `json:"thumb_width"`           // Optional. Thumbnail width
	ThumbHeight         int                   `json:"thumb_height"`          // Optional. Thumbnail height
}

type InlineQueryResultContact struct {
	Type                string                `json:"type"`                  // Type of the result, must be contact
	Id                  string                `json:"id"`                    // Unique identifier for this result, 1-64 Bytes
	PhoneNumber         string                `json:"phone_number"`          // Contact's phone number
	FirstName           string                `json:"first_name"`            // Contact's first name
	LastName            string                `json:"last_name"`             // Optional. Contact's last name
	Vcard               string                `json:"vcard"`                 // Optional. Additional data about the contact in the form of a vCard, 0-2048 bytes
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`          // Optional. Inline keyboard attached to the message
	InputMessageContent interface{}           `json:"input_message_content"` // Optional. Content of the message to be sent instead of the contact
	ThumbUrl            string                `json:"thumb_url"`             // Optional. Url of the thumbnail for the result
	ThumbWidth          int                   `json:"thumb_width"`           // Optional. Thumbnail width
	ThumbHeight         int                   `json:"thumb_height"`          // Optional. Thumbnail height
}

type InlineQueryResultGame struct {
	Type          string                `json:"type"`            // Type of the result, must be game
	Id            string                `json:"id"`              // Unique identifier for this result, 1-64 bytes
	GameShortName string                `json:"game_short_name"` // Short name of the game
	ReplyMarkup   *InlineKeyboardMarkup `json:"reply_markup"`    // Optional. Inline keyboard attached to the message
}

type InlineQueryResultCachedPhoto struct {
	Type                string                `json:"type"`                  // Type of the result, must be photo
	Id                  string                `json:"id"`                    // Unique identifier for this result, 1-64 bytes
	PhotoFileId         string                `json:"photo_file_id"`         // A valid file identifier of the photo
	Title               string                `json:"title"`                 // Optional. Title for the result
	Description         string                `json:"description"`           // Optional. Short description of the result
	Caption             string                `json:"caption"`               // Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	ParseMode           string                `json:"parse_mode"`            // Optional. Mode for parsing entities in the photo caption. See formatting options for more details.
	CaptionEntities     []*MessageEntity      `json:"caption_entities"`      // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`          // Optional. Inline keyboard attached to the message
	InputMessageContent interface{}           `json:"input_message_content"` // Optional. Content of the message to be sent instead of the photo
}

type InlineQueryResultCachedGif struct {
	Type                string                `json:"type"`                  // Type of the result, must be gif
	Id                  string                `json:"id"`                    // Unique identifier for this result, 1-64 bytes
	GifFileId           string                `json:"gif_file_id"`           // A valid file identifier for the GIF file
	Title               string                `json:"title"`                 // Optional. Title for the result
	Caption             string                `json:"caption"`               // Optional. Caption of the GIF file to be sent, 0-1024 characters after entities parsing
	ParseMode           string                `json:"parse_mode"`            // Optional. Mode for parsing entities in the caption. See formatting options for more details.
	CaptionEntities     []*MessageEntity      `json:"caption_entities"`      // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`          // Optional. Inline keyboard attached to the message
	InputMessageContent interface{}           `json:"input_message_content"` // Optional. Content of the message to be sent instead of the GIF animation
}

type InlineQueryResultCachedMpeg4Gif struct {
	Type                string                `json:"type"`                  // Type of the result, must be mpeg4_gif
	Id                  string                `json:"id"`                    // Unique identifier for this result, 1-64 bytes
	Mpeg4FileId         string                `json:"mpeg4_file_id"`         // A valid file identifier for the MP4 file
	Title               string                `json:"title"`                 // Optional. Title for the result
	Caption             string                `json:"caption"`               // Optional. Caption of the MPEG-4 file to be sent, 0-1024 characters after entities parsing
	ParseMode           string                `json:"parse_mode"`            // Optional. Mode for parsing entities in the caption. See formatting options for more details.
	CaptionEntities     []*MessageEntity      `json:"caption_entities"`      // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`          // Optional. Inline keyboard attached to the message
	InputMessageContent interface{}           `json:"input_message_content"` // Optional. Content of the message to be sent instead of the video animation
}

type InlineQueryResultCachedSticker struct {
	Type                string                `json:"type"`                  // Type of the result, must be sticker
	Id                  string                `json:"id"`                    // Unique identifier for this result, 1-64 bytes
	StickerFileId       string                `json:"sticker_file_id"`       // A valid file identifier of the sticker
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`          // Optional. Inline keyboard attached to the message
	InputMessageContent interface{}           `json:"input_message_content"` // Optional. Content of the message to be sent instead of the sticker
}

type InlineQueryResultCachedDocument struct {
	Type                string                `json:"type"`                  // Type of the result, must be document
	Id                  string                `json:"id"`                    // Unique identifier for this result, 1-64 bytes
	Title               string                `json:"title"`                 // Title for the result
	DocumentFileId      string                `json:"document_file_id"`      // A valid file identifier for the file
	Description         string                `json:"description"`           // Optional. Short description of the result
	Caption             string                `json:"caption"`               // Optional. Caption of the document to be sent, 0-1024 characters after entities parsing
	ParseMode           string                `json:"parse_mode"`            // Optional. Mode for parsing entities in the document caption. See formatting options for more details.
	CaptionEntities     []*MessageEntity      `json:"caption_entities"`      // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`          // Optional. Inline keyboard attached to the message
	InputMessageContent interface{}           `json:"input_message_content"` // Optional. Content of the message to be sent instead of the file
}

type InlineQueryResultCachedVideo struct {
	Type                string                `json:"type"`                  // Type of the result, must be video
	Id                  string                `json:"id"`                    // Unique identifier for this result, 1-64 bytes
	VideoFileId         string                `json:"video_file_id"`         // A valid file identifier for the video file
	Title               string                `json:"title"`                 // Title for the result
	Description         string                `json:"description"`           // Optional. Short description of the result
	Caption             string                `json:"caption"`               // Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
	ParseMode           string                `json:"parse_mode"`            // Optional. Mode for parsing entities in the video caption. See formatting options for more details.
	CaptionEntities     []*MessageEntity      `json:"caption_entities"`      // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`          // Optional. Inline keyboard attached to the message
	InputMessageContent interface{}           `json:"input_message_content"` // Optional. Content of the message to be sent instead of the video
}

type InlineQueryResultCachedVoice struct {
	Type                string                `json:"type"`                  // Type of the result, must be voice
	Id                  string                `json:"id"`                    // Unique identifier for this result, 1-64 bytes
	VoiceFileId         string                `json:"voice_file_id"`         // A valid file identifier for the voice message
	Title               string                `json:"title"`                 // Voice message title
	Caption             string                `json:"caption"`               // Optional. Caption, 0-1024 characters after entities parsing
	ParseMode           string                `json:"parse_mode"`            // Optional. Mode for parsing entities in the voice message caption. See formatting options for more details.
	CaptionEntities     []*MessageEntity      `json:"caption_entities"`      // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`          // Optional. Inline keyboard attached to the message
	InputMessageContent interface{}           `json:"input_message_content"` // Optional. Content of the message to be sent instead of the voice message
}

type InlineQueryResultCachedAudio struct {
	Type                string                `json:"type"`                  // Type of the result, must be audio
	Id                  string                `json:"id"`                    // Unique identifier for this result, 1-64 bytes
	AudioFileId         string                `json:"audio_file_id"`         // A valid file identifier for the audio file
	Caption             string                `json:"caption"`               // Optional. Caption, 0-1024 characters after entities parsing
	ParseMode           string                `json:"parse_mode"`            // Optional. Mode for parsing entities in the audio caption. See formatting options for more details.
	CaptionEntities     []*MessageEntity      `json:"caption_entities"`      // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup"`          // Optional. Inline keyboard attached to the message
	InputMessageContent interface{}           `json:"input_message_content"` // Optional. Content of the message to be sent instead of the audio
}

type InputTextMessageContent struct {
	MessageText           string           `json:"message_text"`             // Text of the message to be sent, 1-4096 characters
	ParseMode             string           `json:"parse_mode"`               // Optional. Mode for parsing entities in the message text. See formatting options for more details.
	Entities              []*MessageEntity `json:"entities"`                 // Optional. List of special entities that appear in message text, which can be specified instead of parse_mode
	DisableWebPagePreview bool             `json:"disable_web_page_preview"` // Optional. Disables link previews for links in the sent message
}

type InputLocationMessageContent struct {
	Latitude             float64 `json:"latitude"`               // Latitude of the location in degrees
	Longitude            float64 `json:"longitude"`              // Longitude of the location in degrees
	HorizontalAccuracy   float64 `json:"horizontal_accuracy"`    // Optional. The radius of uncertainty for the location, measured in meters; 0-1500
	LivePeriod           int     `json:"live_period"`            // Optional. Period in seconds for which the location can be updated, should be between 60 and 86400.
	Heading              int     `json:"heading"`                // Optional. For live locations, a direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.
	ProximityAlertRadius int     `json:"proximity_alert_radius"` // Optional. For live locations, a maximum distance for proximity alerts about approaching another chat member, in meters. Must be between 1 and 100000 if specified.
}

type InputVenueMessageContent struct {
	Latitude        float64 `json:"latitude"`          // Latitude of the venue in degrees
	Longitude       float64 `json:"longitude"`         // Longitude of the venue in degrees
	Title           string  `json:"title"`             // Name of the venue
	Address         string  `json:"address"`           // Address of the venue
	FoursquareId    string  `json:"foursquare_id"`     // Optional. Foursquare identifier of the venue, if known
	FoursquareType  string  `json:"foursquare_type"`   // Optional. Foursquare type of the venue, if known. (For example, "arts_entertainment/default", "arts_entertainment/aquarium" or "food/icecream".)
	GooglePlaceId   string  `json:"google_place_id"`   // Optional. Google Places identifier of the venue
	GooglePlaceType string  `json:"google_place_type"` // Optional. Google Places type of the venue. (See supported types.)
}

type InputContactMessageContent struct {
	PhoneNumber string `json:"phone_number"` // Contact's phone number
	FirstName   string `json:"first_name"`   // Contact's first name
	LastName    string `json:"last_name"`    // Optional. Contact's last name
	Vcard       string `json:"vcard"`        // Optional. Additional data about the contact in the form of a vCard, 0-2048 bytes
}

type InputInvoiceMessageContent struct {
	Title                     string          `json:"title"`                         // Product name, 1-32 characters
	Description               string          `json:"description"`                   // Product description, 1-255 characters
	Payload                   string          `json:"payload"`                       // Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use for your internal processes.
	ProviderToken             string          `json:"provider_token"`                // Payment provider token, obtained via Botfather
	Currency                  string          `json:"currency"`                      // Three-letter ISO 4217 currency code, see more on currencies
	Prices                    []*LabeledPrice `json:"prices"`                        // Price breakdown, a JSON-serialized list of components (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.)
	MaxTipAmount              int             `json:"max_tip_amount"`                // Optional. The maximum accepted amount for tips in the smallest units of the currency (integer, not float/double). For example, for a maximum tip of US$ 1.45 pass max_tip_amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies). Defaults to 0
	SuggestedTipAmounts       []int           `json:"suggested_tip_amounts"`         // Optional. A JSON-serialized array of suggested amounts of tip in the smallest units of the currency (integer, not float/double). At most 4 suggested tip amounts can be specified. The suggested tip amounts must be positive, passed in a strictly increased order and must not exceed max_tip_amount.
	ProviderData              string          `json:"provider_data"`                 // Optional. A JSON-serialized object for data about the invoice, which will be shared with the payment provider. A detailed description of the required fields should be provided by the payment provider.
	PhotoUrl                  string          `json:"photo_url"`                     // Optional. URL of the product photo for the invoice. Can be a photo of the goods or a marketing image for a service. People like it better when they see what they are paying for.
	PhotoSize                 int             `json:"photo_size"`                    // Optional. Photo size
	PhotoWidth                int             `json:"photo_width"`                   // Optional. Photo width
	PhotoHeight               int             `json:"photo_height"`                  // Optional. Photo height
	NeedName                  bool            `json:"need_name"`                     // Optional. Pass True, if you require the user's full name to complete the order
	NeedPhoneNumber           bool            `json:"need_phone_number"`             // Optional. Pass True, if you require the user's phone number to complete the order
	NeedEmail                 bool            `json:"need_email"`                    // Optional. Pass True, if you require the user's email address to complete the order
	NeedShippingAddress       bool            `json:"need_shipping_address"`         // Optional. Pass True, if you require the user's shipping address to complete the order
	SendPhoneNumberToProvider bool            `json:"send_phone_number_to_provider"` // Optional. Pass True, if user's phone number should be sent to provider
	SendEmailToProvider       bool            `json:"send_email_to_provider"`        // Optional. Pass True, if user's email address should be sent to provider
	IsFlexible                bool            `json:"is_flexible"`                   // Optional. Pass True, if the final price depends on the shipping method
}

type ChosenInlineResult struct {
	ResultId        string    `json:"result_id"`         // The unique identifier for the result that was chosen
	From            *User     `json:"from"`              // The user that chose the result
	Location        *Location `json:"location"`          // Optional. Sender location, only for bots that require user location
	InlineMessageId string    `json:"inline_message_id"` // Optional. Identifier of the sent inline message. Available only if there is an inline keyboard attached to the message. Will be also received in callback queries and can be used to edit the message.
	Query           string    `json:"query"`             // The query that was used to obtain the result
}

type LabeledPrice struct {
	Label  string `json:"label"`  // Portion label
	Amount int    `json:"amount"` // Price of the product in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
}

type Invoice struct {
	Title          string `json:"title"`           // Product name
	Description    string `json:"description"`     // Product description
	StartParameter string `json:"start_parameter"` // Unique bot deep-linking parameter that can be used to generate this invoice
	Currency       string `json:"currency"`        // Three-letter ISO 4217 currency code
	TotalAmount    int    `json:"total_amount"`    // Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
}

type ShippingAddress struct {
	CountryCode string `json:"country_code"` // ISO 3166-1 alpha-2 country code
	State       string `json:"state"`        // State, if applicable
	City        string `json:"city"`         // City
	StreetLine1 string `json:"street_line1"` // First line for the address
	StreetLine2 string `json:"street_line2"` // Second line for the address
	PostCode    string `json:"post_code"`    // Address post code
}

type OrderInfo struct {
	Name            string           `json:"name"`             // Optional. User name
	PhoneNumber     string           `json:"phone_number"`     // Optional. User's phone number
	Email           string           `json:"email"`            // Optional. User email
	ShippingAddress *ShippingAddress `json:"shipping_address"` // Optional. User shipping address
}

type ShippingOption struct {
	Id     string          `json:"id"`     // Shipping option identifier
	Title  string          `json:"title"`  // Option title
	Prices []*LabeledPrice `json:"prices"` // List of price portions
}

type SuccessfulPayment struct {
	Currency                string     `json:"currency"`                   // Three-letter ISO 4217 currency code
	TotalAmount             int        `json:"total_amount"`               // Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	InvoicePayload          string     `json:"invoice_payload"`            // Bot specified invoice payload
	ShippingOptionId        string     `json:"shipping_option_id"`         // Optional. Identifier of the shipping option chosen by the user
	OrderInfo               *OrderInfo `json:"order_info"`                 // Optional. Order info provided by the user
	TelegramPaymentChargeId string     `json:"telegram_payment_charge_id"` // Telegram payment identifier
	ProviderPaymentChargeId string     `json:"provider_payment_charge_id"` // Provider payment identifier
}

type ShippingQuery struct {
	Id              string           `json:"id"`               // Unique query identifier
	From            *User            `json:"from"`             // User who sent the query
	InvoicePayload  string           `json:"invoice_payload"`  // Bot specified invoice payload
	ShippingAddress *ShippingAddress `json:"shipping_address"` // User specified shipping address
}

type PreCheckoutQuery struct {
	Id               string     `json:"id"`                 // Unique query identifier
	From             *User      `json:"from"`               // User who sent the query
	Currency         string     `json:"currency"`           // Three-letter ISO 4217 currency code
	TotalAmount      int        `json:"total_amount"`       // Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	InvoicePayload   string     `json:"invoice_payload"`    // Bot specified invoice payload
	ShippingOptionId string     `json:"shipping_option_id"` // Optional. Identifier of the shipping option chosen by the user
	OrderInfo        *OrderInfo `json:"order_info"`         // Optional. Order info provided by the user
}

type PassportData struct {
	Data        []*EncryptedPassportElement `json:"data"`        // Array with information about documents and other Telegram Passport elements that was shared with the bot
	Credentials *EncryptedCredentials       `json:"credentials"` // Encrypted credentials required to decrypt the data
}

type PassportFile struct {
	FileId       string `json:"file_id"`        // Identifier for this file, which can be used to download or reuse the file
	FileUniqueId string `json:"file_unique_id"` // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	FileSize     int    `json:"file_size"`      // File size
	FileDate     int    `json:"file_date"`      // Unix time when the file was uploaded
}

type EncryptedPassportElement struct {
	Type        string          `json:"type"`         // Element type. One of "personal_details", "passport", "driver_license", "identity_card", "internal_passport", "address", "utility_bill", "bank_statement", "rental_agreement", "passport_registration", "temporary_registration", "phone_number", "email".
	Data        string          `json:"data"`         // Optional. Base64-encoded encrypted Telegram Passport element data provided by the user, available for "personal_details", "passport", "driver_license", "identity_card", "internal_passport" and "address" types. Can be decrypted and verified using the accompanying EncryptedCredentials.
	PhoneNumber string          `json:"phone_number"` // Optional. User's verified phone number, available only for "phone_number" type
	Email       string          `json:"email"`        // Optional. User's verified email address, available only for "email" type
	Files       []*PassportFile `json:"files"`        // Optional. Array of encrypted files with documents provided by the user, available for "utility_bill", "bank_statement", "rental_agreement", "passport_registration" and "temporary_registration" types. Files can be decrypted and verified using the accompanying EncryptedCredentials.
	FrontSide   *PassportFile   `json:"front_side"`   // Optional. Encrypted file with the front side of the document, provided by the user. Available for "passport", "driver_license", "identity_card" and "internal_passport". The file can be decrypted and verified using the accompanying EncryptedCredentials.
	ReverseSide *PassportFile   `json:"reverse_side"` // Optional. Encrypted file with the reverse side of the document, provided by the user. Available for "driver_license" and "identity_card". The file can be decrypted and verified using the accompanying EncryptedCredentials.
	Selfie      *PassportFile   `json:"selfie"`       // Optional. Encrypted file with the selfie of the user holding a document, provided by the user; available for "passport", "driver_license", "identity_card" and "internal_passport". The file can be decrypted and verified using the accompanying EncryptedCredentials.
	Translation []*PassportFile `json:"translation"`  // Optional. Array of encrypted files with translated versions of documents provided by the user. Available if requested for "passport", "driver_license", "identity_card", "internal_passport", "utility_bill", "bank_statement", "rental_agreement", "passport_registration" and "temporary_registration" types. Files can be decrypted and verified using the accompanying EncryptedCredentials.
	Hash        string          `json:"hash"`         // Base64-encoded element hash for using in PassportElementErrorUnspecified
}

type EncryptedCredentials struct {
	Data   string `json:"data"`   // Base64-encoded encrypted JSON-serialized data with unique user's payload, data hashes and secrets required for EncryptedPassportElement decryption and authentication
	Hash   string `json:"hash"`   // Base64-encoded data hash for data authentication
	Secret string `json:"secret"` // Base64-encoded secret, encrypted with the bot's public RSA key, required for data decryption
}

type PassportElementErrorDataField struct {
	Source    string `json:"source"`     // Error source, must be data
	Type      string `json:"type"`       // The section of the user's Telegram Passport which has the error, one of "personal_details", "passport", "driver_license", "identity_card", "internal_passport", "address"
	FieldName string `json:"field_name"` // Name of the data field which has the error
	DataHash  string `json:"data_hash"`  // Base64-encoded data hash
	Message   string `json:"message"`    // Error message
}

type PassportElementErrorFrontSide struct {
	Source   string `json:"source"`    // Error source, must be front_side
	Type     string `json:"type"`      // The section of the user's Telegram Passport which has the issue, one of "passport", "driver_license", "identity_card", "internal_passport"
	FileHash string `json:"file_hash"` // Base64-encoded hash of the file with the front side of the document
	Message  string `json:"message"`   // Error message
}

type PassportElementErrorReverseSide struct {
	Source   string `json:"source"`    // Error source, must be reverse_side
	Type     string `json:"type"`      // The section of the user's Telegram Passport which has the issue, one of "driver_license", "identity_card"
	FileHash string `json:"file_hash"` // Base64-encoded hash of the file with the reverse side of the document
	Message  string `json:"message"`   // Error message
}

type PassportElementErrorSelfie struct {
	Source   string `json:"source"`    // Error source, must be selfie
	Type     string `json:"type"`      // The section of the user's Telegram Passport which has the issue, one of "passport", "driver_license", "identity_card", "internal_passport"
	FileHash string `json:"file_hash"` // Base64-encoded hash of the file with the selfie
	Message  string `json:"message"`   // Error message
}

type PassportElementErrorFile struct {
	Source   string `json:"source"`    // Error source, must be file
	Type     string `json:"type"`      // The section of the user's Telegram Passport which has the issue, one of "utility_bill", "bank_statement", "rental_agreement", "passport_registration", "temporary_registration"
	FileHash string `json:"file_hash"` // Base64-encoded file hash
	Message  string `json:"message"`   // Error message
}

type PassportElementErrorFiles struct {
	Source     string   `json:"source"`      // Error source, must be files
	Type       string   `json:"type"`        // The section of the user's Telegram Passport which has the issue, one of "utility_bill", "bank_statement", "rental_agreement", "passport_registration", "temporary_registration"
	FileHashes []string `json:"file_hashes"` // List of base64-encoded file hashes
	Message    string   `json:"message"`     // Error message
}

type PassportElementErrorTranslationFile struct {
	Source   string `json:"source"`    // Error source, must be translation_file
	Type     string `json:"type"`      // Type of element of the user's Telegram Passport which has the issue, one of "passport", "driver_license", "identity_card", "internal_passport", "utility_bill", "bank_statement", "rental_agreement", "passport_registration", "temporary_registration"
	FileHash string `json:"file_hash"` // Base64-encoded file hash
	Message  string `json:"message"`   // Error message
}

type PassportElementErrorTranslationFiles struct {
	Source     string   `json:"source"`      // Error source, must be translation_files
	Type       string   `json:"type"`        // Type of element of the user's Telegram Passport which has the issue, one of "passport", "driver_license", "identity_card", "internal_passport", "utility_bill", "bank_statement", "rental_agreement", "passport_registration", "temporary_registration"
	FileHashes []string `json:"file_hashes"` // List of base64-encoded file hashes
	Message    string   `json:"message"`     // Error message
}

type PassportElementErrorUnspecified struct {
	Source      string `json:"source"`       // Error source, must be unspecified
	Type        string `json:"type"`         // Type of element of the user's Telegram Passport which has the issue
	ElementHash string `json:"element_hash"` // Base64-encoded element hash
	Message     string `json:"message"`      // Error message
}

type Game struct {
	Title        string           `json:"title"`         // Title of the game
	Description  string           `json:"description"`   // Description of the game
	Photo        []*PhotoSize     `json:"photo"`         // Photo that will be displayed in the game message in chats.
	Text         string           `json:"text"`          // Optional. Brief description of the game or high scores included in the game message. Can be automatically edited to include current high scores for the game when the bot calls setGameScore, or manually edited using editMessageText. 0-4096 characters.
	TextEntities []*MessageEntity `json:"text_entities"` // Optional. Special entities that appear in text, such as usernames, URLs, bot commands, etc.
	Animation    *Animation       `json:"animation"`     // Optional. Animation that will be displayed in the game message in chats. Upload via BotFather
}

type CallbackGame struct {
}

type GameHighScore struct {
	Position int   `json:"position"` // Position in high score table for the game
	User     *User `json:"user"`     // User
	Score    int   `json:"score"`    // Score
}
