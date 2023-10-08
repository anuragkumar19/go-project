package validations

type RegisterParameters struct {
	Name     string `json:"name" validate:"required,min=3,max=30" mod:"trim"`
	Username string `json:"username" validate:"required,min=3,max=30,alphanum" mod:"trim,lcase"`
	Email    string `json:"email" validate:"required,email" mod:"trim,lcase"`
	Password string `json:"password" validate:"required,min=8,max=60"`
}

type VerifyEmailParameters struct {
	Email string `json:"email" validate:"required,email" mod:"trim,lcase"`
	OTP   int    `json:"otp" validate:"required,min=100000,max=999999"`
}

type ForgotPasswordParameters struct {
	Email string `json:"email" validate:"required,email" mod:"trim,lcase"`
}

type ResetPasswordParameters struct {
	Email    string `json:"email" validate:"required,email" mod:"trim,lcase"`
	OTP      int    `json:"otp" validate:"required,min=100000,max=999999"`
	Password string `json:"password" validate:"required,min=8,max=60"`
}

type LoginParameters struct {
	Identifier string `json:"identifier" validate:"required" mod:"trim,lcase"`
	Password   string `json:"password" validate:"required,min=8,max=60"`
}
