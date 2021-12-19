package model

// NewCommonRespWithErrorMsg create error common response with error msg.
func NewCommonRespWithErrorMsg(msg string) *CommonResp {
	return &CommonResp{Error: &Error{Error: true, Message: msg}}
}
