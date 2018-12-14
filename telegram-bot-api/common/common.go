/*
Package common 通用常量、变量定义
Created by chenguolin 2018-12-13
*/
package common

// Telegram constants
const (
	// APIEndpoint is the endpoint for all API methods,
	// with formatting for Sprintf.
	APIEndpoint = "https://api.telegram.org/bot%s/%s"
	// FileEndpoint is the endpoint for downloading a file from Telegram.
	FileEndpoint = "https://api.telegram.org/file/bot%s/%s"
)

// Constant values for ChatActions
const (
	ChatTyping         = "typing"
	ChatUploadPhoto    = "upload_photo"
	ChatRecordVideo    = "record_video"
	ChatUploadVideo    = "upload_video"
	ChatRecordAudio    = "record_audio"
	ChatUploadAudio    = "upload_audio"
	ChatUploadDocument = "upload_document"
	ChatFindLocation   = "find_location"
)

// Constant values for ParseMode in MessageConfig
const (
	ModeMarkdown = "Markdown"
	ModeHTML     = "HTML"
)

// API errors
const (
	// ErrAPIForbidden happens when a token is bad
	ErrAPIForbidden = "forbidden"
)

// Library errors
const (
	// ErrBadFileType happens when you pass an unknown type
	ErrBadFileType = "bad file type"
	// ErrBadURL happens when bad or empty url
	ErrBadURL = "bad or empty url"
)
