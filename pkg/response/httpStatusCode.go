package response

const (
	SuccessCode         = 20001
	ErrCodeParamInvalid = 20002
)

var MSG = map[int]string{
	SuccessCode:         "success",
	ErrCodeParamInvalid: "phone number is invalid",
}
