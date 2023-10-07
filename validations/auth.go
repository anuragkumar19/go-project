package validations

type RegisterParameters struct {
	Name     string `json:"name" validate:"required,min=3,max=30" mod:"trim"`
	Username string `json:"username" validate:"required,min=3,max=30,alphanum" mod:"trim"`
	Email    string `json:"email" validate:"required,email" mod:"trim"`
	Password string `json:"password" validate:"required,min=8,max=60"`
}
