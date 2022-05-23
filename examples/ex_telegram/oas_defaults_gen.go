// Code generated by ogen, DO NOT EDIT.

package api

// setDefaults set default value of fields.
func (s *Error) setDefaults() {
	{
		val := bool(false)
		s.Ok = val
	}
}

// setDefaults set default value of fields.
func (s *GetUpdates) setDefaults() {
	{
		val := int(0)
		s.Offset.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *GetUserProfilePhotos) setDefaults() {
	{
		val := int(0)
		s.Offset.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *InlineQueryResultCachedAudio) setDefaults() {
	{
		val := string("audio")
		s.Type = val
	}
}

// setDefaults set default value of fields.
func (s *InlineQueryResultCachedDocument) setDefaults() {
	{
		val := string("document")
		s.Type = val
	}
}

// setDefaults set default value of fields.
func (s *InlineQueryResultCachedGif) setDefaults() {
	{
		val := string("gif")
		s.Type = val
	}
}

// setDefaults set default value of fields.
func (s *InlineQueryResultCachedMpeg4Gif) setDefaults() {
	{
		val := string("mpeg4_gif")
		s.Type = val
	}
}

// setDefaults set default value of fields.
func (s *InlineQueryResultCachedPhoto) setDefaults() {
	{
		val := string("photo")
		s.Type = val
	}
}

// setDefaults set default value of fields.
func (s *InlineQueryResultCachedVideo) setDefaults() {
	{
		val := string("video")
		s.Type = val
	}
}

// setDefaults set default value of fields.
func (s *InlineQueryResultCachedVoice) setDefaults() {
	{
		val := string("voice")
		s.Type = val
	}
}

// setDefaults set default value of fields.
func (s *MessageEntity) setDefaults() {
	{
		val := int(0)
		s.Offset = val
	}
}

// setDefaults set default value of fields.
func (s *PassportElementErrorDataField) setDefaults() {
	{
		val := string("data")
		s.Source = val
	}
}

// setDefaults set default value of fields.
func (s *PassportElementErrorFile) setDefaults() {
	{
		val := string("file")
		s.Source = val
	}
}

// setDefaults set default value of fields.
func (s *PassportElementErrorFiles) setDefaults() {
	{
		val := string("files")
		s.Source = val
	}
}

// setDefaults set default value of fields.
func (s *PassportElementErrorFrontSide) setDefaults() {
	{
		val := string("front_side")
		s.Source = val
	}
}

// setDefaults set default value of fields.
func (s *PassportElementErrorReverseSide) setDefaults() {
	{
		val := string("reverse_side")
		s.Source = val
	}
}

// setDefaults set default value of fields.
func (s *PassportElementErrorSelfie) setDefaults() {
	{
		val := string("selfie")
		s.Source = val
	}
}

// setDefaults set default value of fields.
func (s *PassportElementErrorTranslationFile) setDefaults() {
	{
		val := string("translation_file")
		s.Source = val
	}
}

// setDefaults set default value of fields.
func (s *PassportElementErrorTranslationFiles) setDefaults() {
	{
		val := string("translation_files")
		s.Source = val
	}
}

// setDefaults set default value of fields.
func (s *PassportElementErrorUnspecified) setDefaults() {
	{
		val := string("unspecified")
		s.Source = val
	}
}

// setDefaults set default value of fields.
func (s *Result) setDefaults() {
	{
		val := bool(true)
		s.Ok = val
	}
}

// setDefaults set default value of fields.
func (s *ResultArrayOfBotCommand) setDefaults() {
	{
		val := bool(true)
		s.Ok = val
	}
}

// setDefaults set default value of fields.
func (s *ResultArrayOfChatMember) setDefaults() {
	{
		val := bool(true)
		s.Ok = val
	}
}

// setDefaults set default value of fields.
func (s *ResultArrayOfGameHighScore) setDefaults() {
	{
		val := bool(true)
		s.Ok = val
	}
}

// setDefaults set default value of fields.
func (s *ResultArrayOfMessage) setDefaults() {
	{
		val := bool(true)
		s.Ok = val
	}
}

// setDefaults set default value of fields.
func (s *ResultArrayOfUpdate) setDefaults() {
	{
		val := bool(true)
		s.Ok = val
	}
}

// setDefaults set default value of fields.
func (s *ResultChat) setDefaults() {
	{
		val := bool(true)
		s.Ok = val
	}
}

// setDefaults set default value of fields.
func (s *ResultChatInviteLink) setDefaults() {
	{
		val := bool(true)
		s.Ok = val
	}
}

// setDefaults set default value of fields.
func (s *ResultChatMember) setDefaults() {
	{
		val := bool(true)
		s.Ok = val
	}
}

// setDefaults set default value of fields.
func (s *ResultFile) setDefaults() {
	{
		val := bool(true)
		s.Ok = val
	}
}

// setDefaults set default value of fields.
func (s *ResultInt) setDefaults() {
	{
		val := bool(true)
		s.Ok = val
	}
}

// setDefaults set default value of fields.
func (s *ResultMessage) setDefaults() {
	{
		val := bool(true)
		s.Ok = val
	}
}

// setDefaults set default value of fields.
func (s *ResultMessageId) setDefaults() {
	{
		val := bool(true)
		s.Ok = val
	}
}

// setDefaults set default value of fields.
func (s *ResultPoll) setDefaults() {
	{
		val := bool(true)
		s.Ok = val
	}
}

// setDefaults set default value of fields.
func (s *ResultString) setDefaults() {
	{
		val := bool(true)
		s.Ok = val
	}
}

// setDefaults set default value of fields.
func (s *ResultUser) setDefaults() {
	{
		val := bool(true)
		s.Ok = val
	}
}

// setDefaults set default value of fields.
func (s *ResultUserProfilePhotos) setDefaults() {
	{
		val := bool(true)
		s.Ok = val
	}
}

// setDefaults set default value of fields.
func (s *ResultWebhookInfo) setDefaults() {
	{
		val := bool(true)
		s.Ok = val
	}
}

// setDefaults set default value of fields.
func (s *UploadStickerFile) setDefaults() {
	{
		val := string("up to 512 kilobytes in size, dimensions must not exceed 512px, and either width or height must be exactly 512px. More info on Sending Files »")
		s.PNGSticker = val
	}
}
