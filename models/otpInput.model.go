package models

type OTPInput struct {
	UserId string `json:"user_id"`
	Token  string `json:"token"`
}
