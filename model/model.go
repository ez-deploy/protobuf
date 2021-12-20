package model

// NewCommonRespWithErrorMessage create error common response with error message.
func NewCommonRespWithErrorMessage(message string) *CommonResp {
	return &CommonResp{Error: &Error{Error: true, Message: message}}
}

// NewCommonRespWithError create error common response with error.
func NewCommonRespWithError(err error) *CommonResp {
	return &CommonResp{Error: &Error{Error: true, Message: err.Error()}}
}

// NewError from error.
func NewError(err error) *Error {
	return &Error{Error: true, Message: err.Error()}
}

// NewErrorWithMessage from error message.
func NewErrorWithMessage(message string) *Error {
	return &Error{Error: true, Message: message}
}
