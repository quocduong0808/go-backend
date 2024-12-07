package response

const (
	SuccessCode         = 20001
	ErrCodeParamInvalid = 20002
	ErrCodeTokenInvalid = 20003
)

var MSG = map[int]string{
	SuccessCode:         "success",
	ErrCodeParamInvalid: "phone number is invalid",
	ErrCodeTokenInvalid: "token invalid",
}
