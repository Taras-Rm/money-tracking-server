package models

type CreateUserInput struct {
	Name     string
	Email    string
	Password string
}

type LoginUserInput struct {
	Email    string
	Password string
}

type LoginUserWithGoogleInput struct {
	AccessToken string
}

type GetMeInput struct {
	AccessToken string
}
