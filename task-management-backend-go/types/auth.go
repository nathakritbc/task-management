package types

type AuthSignIn struct {
	Username string `json:"username"  form:"username" validate:"required"`
	Password string `json:"password"  form:"password" validate:"required"`
	OtpCode  string `json:"otp_code" form:"otp_code"`
}
