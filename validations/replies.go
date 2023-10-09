package validations

type ReplyParameters struct {
	Content string `json:"content" validation:"required,max=500" mod:"trim"`
}
