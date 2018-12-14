/*
Package types Message结构体定义
Created by chenguolin 2018-12-12
*/
package types

import (
	"errors"
	"net/url"
	"strings"
	"time"

	"golang/telegram-bot-api/common"
	"golang/telegram-bot-api/config"
)

// Message is returned by almost every request, and contains data about
// almost anything.
type Message struct {
	MessageID             int                  `json:"message_id"`
	From                  *User                `json:"from"` // optional
	Date                  int                  `json:"date"`
	Chat                  *Chat                `json:"chat"`
	ForwardFrom           *User                `json:"forward_from"`            // optional
	ForwardFromChat       *Chat                `json:"forward_from_chat"`       // optional
	ForwardFromMessageID  int                  `json:"forward_from_message_id"` // optional
	ForwardDate           int                  `json:"forward_date"`            // optional
	ReplyToMessage        *Message             `json:"reply_to_message"`        // optional
	EditDate              int                  `json:"edit_date"`               // optional
	Text                  string               `json:"text"`                    // optional
	Entities              *[]MessageEntity     `json:"entities"`                // optional
	Audio                 *Audio               `json:"audio"`                   // optional
	Document              *Document            `json:"document"`                // optional
	Animation             *ChatAnimation       `json:"animation"`               // optional
	Game                  *Game                `json:"game"`                    // optional
	Photo                 *[]PhotoSize         `json:"photo"`                   // optional
	Sticker               *Sticker             `json:"sticker"`                 // optional
	Video                 *Video               `json:"video"`                   // optional
	VideoNote             *VideoNote           `json:"video_note"`              // optional
	Voice                 *Voice               `json:"voice"`                   // optional
	Caption               string               `json:"caption"`                 // optional
	Contact               *Contact             `json:"contact"`                 // optional
	Location              *Location            `json:"location"`                // optional
	Venue                 *Venue               `json:"venue"`                   // optional
	NewChatMembers        *[]User              `json:"new_chat_members"`        // optional
	LeftChatMember        *User                `json:"left_chat_member"`        // optional
	NewChatTitle          string               `json:"new_chat_title"`          // optional
	NewChatPhoto          *[]PhotoSize         `json:"new_chat_photo"`          // optional
	DeleteChatPhoto       bool                 `json:"delete_chat_photo"`       // optional
	GroupChatCreated      bool                 `json:"group_chat_created"`      // optional
	SuperGroupChatCreated bool                 `json:"supergroup_chat_created"` // optional
	ChannelChatCreated    bool                 `json:"channel_chat_created"`    // optional
	MigrateToChatID       int64                `json:"migrate_to_chat_id"`      // optional
	MigrateFromChatID     int64                `json:"migrate_from_chat_id"`    // optional
	PinnedMessage         *Message             `json:"pinned_message"`          // optional
	Invoice               *Invoice             `json:"invoice"`                 // optional
	SuccessfulPayment     *SuccessfulPayment   `json:"successful_payment"`      // optional
	PassportData          *config.PassportData `json:"passport_data,omitempty"` // optional
}

// Time converts the message timestamp into a Time.
func (m *Message) Time() time.Time {
	return time.Unix(int64(m.Date), 0)
}

// IsCommand returns true if message starts with a "bot_command" entity.
func (m *Message) IsCommand() bool {
	if m.Entities == nil || len(*m.Entities) == 0 {
		return false
	}

	entity := (*m.Entities)[0]
	return entity.Offset == 0 && entity.Type == "bot_command"
}

// Command checks if the message was a command and if it was, returns the
// command. If the Message was not a command, it returns an empty string.
// If the command contains the at name syntax, it is removed. Use
// CommandWithAt() if you do not want that.
func (m *Message) Command() string {
	command := m.CommandWithAt()

	if i := strings.Index(command, "@"); i != -1 {
		command = command[:i]
	}

	return command
}

// CommandWithAt checks if the message was a command and if it was, returns the
// command. If the Message was not a command, it returns an empty string.
// If the command contains the at name syntax, it is not removed. Use Command()
// if you want that.
func (m *Message) CommandWithAt() string {
	if !m.IsCommand() {
		return ""
	}

	// IsCommand() checks that the message begins with a bot_command entity
	entity := (*m.Entities)[0]
	return m.Text[1:entity.Length]
}

// CommandArguments checks if the message was a command and if it was,
// returns all text after the command name. If the Message was not a
// command, it returns an empty string.
// Note: The first character after the command name is omitted:
// - "/foo bar baz" yields "bar baz", not " bar baz"
// - "/foo-bar baz" yields "bar baz", too
// Even though the latter is not a command conforming to the spec, the API
// marks "/foo" as command entity.
func (m *Message) CommandArguments() string {
	if !m.IsCommand() {
		return ""
	}

	// IsCommand() checks that the message begins with a bot_command entity
	entity := (*m.Entities)[0]
	if len(m.Text) == entity.Length {
		return "" // The command makes up the whole message
	}

	return m.Text[entity.Length+1:]
}

// MessageEntity contains information about data in a Message.
type MessageEntity struct {
	Type   string `json:"type"`
	Offset int    `json:"offset"`
	Length int    `json:"length"`
	URL    string `json:"url"`  // optional
	User   *User  `json:"user"` // optional
}

// ParseURL attempts to parse a URL contained within a MessageEntity.
func (entity MessageEntity) ParseURL() (*url.URL, error) {
	if entity.URL == "" {
		return nil, errors.New(common.ErrBadURL)
	}

	return url.Parse(entity.URL)
}
