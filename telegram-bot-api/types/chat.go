/*
Package types Chat相关结构体定义
Created by chenguolin 2018-12-12
*/
package types

import "net/url"

// Chattable is any config type that can be sent.
type Chattable interface {
	values() (url.Values, error)
	method() string
}

// Fileable is any config type that can be sent that includes a file.
type Fileable interface {
	Chattable
	params() (map[string]string, error)
	name() string
	getFile() interface{}
	useExistingFile() bool
}

// GroupChat is a group chat.
type GroupChat struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// ChatPhoto represents a chat photo.
type ChatPhoto struct {
	SmallFileID string `json:"small_file_id"`
	BigFileID   string `json:"big_file_id"`
}

// ChatConfig contains information about getting information on a chat.
type ChatConfig struct {
	ChatID             int64  `json:"chat_id"`
	SuperGroupUsername string `json:"super_group_username"`
}

// Chat contains information about the place a message was sent.
type Chat struct {
	ID                  int64      `json:"id"`
	Type                string     `json:"type"`
	Title               string     `json:"title"`                          // optional
	UserName            string     `json:"username"`                       // optional
	FirstName           string     `json:"first_name"`                     // optional
	LastName            string     `json:"last_name"`                      // optional
	AllMembersAreAdmins bool       `json:"all_members_are_administrators"` // optional
	Photo               *ChatPhoto `json:"photo"`
	Description         string     `json:"description,omitempty"` // optional
	InviteLink          string     `json:"invite_link,omitempty"` // optional
}

// IsPrivate returns if the Chat is a private conversation.
func (c Chat) IsPrivate() bool {
	return c.Type == "private"
}

// IsGroup returns if the Chat is a group.
func (c Chat) IsGroup() bool {
	return c.Type == "group"
}

// IsSuperGroup returns if the Chat is a supergroup.
func (c Chat) IsSuperGroup() bool {
	return c.Type == "supergroup"
}

// IsChannel returns if the Chat is a channel.
func (c Chat) IsChannel() bool {
	return c.Type == "channel"
}

// ChatConfig returns a ChatConfig struct for chat related methods.
func (c Chat) ChatConfig() ChatConfig {
	return ChatConfig{
		ChatID: c.ID,
	}
}
