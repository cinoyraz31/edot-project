package request

type OTPSendRequest struct {
	PhoneNumber string `json:"phoneNumber" validate:"required"`
}

type OTPValidateRequest struct {
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	Code        string `json:"code" validate:"required,min=6,max=6"`
}
