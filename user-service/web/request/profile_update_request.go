package request

type ProfileUpdateRequest struct {
	Name        string `json:"name" validate:"required"`
	DateOfBirth string `json:"dateOfBirth" validate:"required"`
}
