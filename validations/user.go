package validations

type UpdateNameParameters struct {
	Name string `json:"name" validate:"required,min=3,max=30" mod:"trim"`
}

type UpdateUsernameParameters struct {
	Username string `json:"username" validate:"required,min=3,max=30,alphanum" mod:"trim,lcase"`
}

type UpdatePasswordParameters struct {
	OldPassword string `json:"old_password" validate:"required,min=8,max=60"`
	NewPassword string `json:"new_password" validate:"required,min=8,max=60"`
}
