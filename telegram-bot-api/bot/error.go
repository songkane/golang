package bot

// Error is an error containing extra information returned by the Telegram API.
type Error struct {
	Message string
	ResponseParameters
}

func (e Error) Error() string {
	return e.Message
}
