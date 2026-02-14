package dto

type RegistrationRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginWithGoogleRequest struct {
	AccessToken string `json:"access_token"`
}

type MeRequest struct {
	AccessToken string `json:"access_token"`
}
