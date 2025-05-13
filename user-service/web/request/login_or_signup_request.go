package request

type LoginOrSignupRequest struct {
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	OtpId       string `json:"otpId" validate:"required"`
}
