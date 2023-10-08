package validations

type CreateSubredditParameters struct {
	Name string `json:"subreddit" validate:"required,min=3,max=30,alphanum" mod:"trim"`
}

type UpdateSubredditTitleParameters struct {
	Title string `json:"title" validate:"max=30" mod:"trim"`
}
