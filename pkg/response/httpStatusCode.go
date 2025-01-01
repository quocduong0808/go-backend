package response

const (
	SuccessCode         = 20001
	ErrCodeParamInvalid = 20002
	ErrCodeTokenInvalid = 20003
	ErrorCodeEmailExist = 20004
	ErrCodeGenerateOTP  = 20005
	ErrCodeSaveOTPRedis = 20006
	ErrCodeSendMail     = 20007
)

var MSG = map[int]string{
	SuccessCode:         "success",
	ErrCodeParamInvalid: "phone number is invalid",
	ErrCodeTokenInvalid: "token invalid",
	ErrorCodeEmailExist: "email exist",
	ErrCodeGenerateOTP:  "generate otp error",
	ErrCodeSaveOTPRedis: "save otp error",
	ErrCodeSendMail:     "send mail error",
}
