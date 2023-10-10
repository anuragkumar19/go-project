package validations

type CreateSubredditParameters struct {
	Name string `json:"name" validate:"required,min=3,max=30,alphanum" mod:"trim"`
}

type UpdateSubredditTitleParameters struct {
	Title string `json:"title" validate:"max=30" mod:"trim"`
}

type UpdateSubredditAboutParameters struct {
	About string `json:"about" validate:"max=500" mod:"trim"`
}
