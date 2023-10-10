package validations

type CreatePostWithTextParameters struct {
	Title string `json:"title" form:"title" validate:"max=30,required" mod:"trim"`
	Text  string `json:"text" form:"text" validate:"max=2000,required" mod:"trim"`
}

type CreatePostWithMediaParameters struct {
	Title string `json:"title" form:"title" validate:"max=30,required" mod:"trim"`
	Text  string `json:"text" form:"text" validate:"max=2000" mod:"trim"`
}

type CreatePostWithLinkParameters struct {
	Title string `json:"title" validate:"max=30,required" mod:"trim"`
	Text  string `json:"text" validate:"max=2000,required" mod:"trim"`
	Link  string `json:"link" validate:"http_url" mod:"trim"`
}

type VotePostParameters struct {
	Down bool `json:"down" validate:"boolean"`
}
