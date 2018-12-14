/*
Package types ChatMember结构体定义
Created by chenguolin 2018-12-12
*/
package types

// ChatMember is information about a member in a chat.
type ChatMember struct {
	User                  *User  `json:"user"`
	Status                string `json:"status"`
	UntilDate             int64  `json:"until_date,omitempty"`                // optional
	CanBeEdited           bool   `json:"can_be_edited,omitempty"`             // optional
	CanChangeInfo         bool   `json:"can_change_info,omitempty"`           // optional
	CanPostMessages       bool   `json:"can_post_messages,omitempty"`         // optional
	CanEditMessages       bool   `json:"can_edit_messages,omitempty"`         // optional
	CanDeleteMessages     bool   `json:"can_delete_messages,omitempty"`       // optional
	CanInviteUsers        bool   `json:"can_invite_users,omitempty"`          // optional
	CanRestrictMembers    bool   `json:"can_restrict_members,omitempty"`      // optional
	CanPinMessages        bool   `json:"can_pin_messages,omitempty"`          // optional
	CanPromoteMembers     bool   `json:"can_promote_members,omitempty"`       // optional
	CanSendMessages       bool   `json:"can_send_messages,omitempty"`         // optional
	CanSendMediaMessages  bool   `json:"can_send_media_messages,omitempty"`   // optional
	CanSendOtherMessages  bool   `json:"can_send_other_messages,omitempty"`   // optional
	CanAddWebPagePreviews bool   `json:"can_add_web_page_previews,omitempty"` // optional
}

// IsCreator returns if the ChatMember was the creator of the chat.
func (chat ChatMember) IsCreator() bool {
	return chat.Status == "creator"
}

// IsAdministrator returns if the ChatMember is a chat administrator.
func (chat ChatMember) IsAdministrator() bool {
	return chat.Status == "administrator"
}

// IsMember returns if the ChatMember is a current member of the chat.
func (chat ChatMember) IsMember() bool {
	return chat.Status == "member"
}

// HasLeft returns if the ChatMember left the chat.
func (chat ChatMember) HasLeft() bool {
	return chat.Status == "left"
}

// WasKicked returns if the ChatMember was kicked from the chat.
func (chat ChatMember) WasKicked() bool {
	return chat.Status == "kicked"
}
