package oauth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type OAuthManager struct {
	GoogleApiUrl string
}

func NewOAuthManager(googleApiUrl string) *OAuthManager {
	return &OAuthManager{
		GoogleApiUrl: googleApiUrl,
	}
}

type GoogleUser struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
}

func (o *OAuthManager) GetGoogleUser(ctx context.Context, accessToken string) (*GoogleUser, error) {
	serverRequest, err := http.NewRequestWithContext(ctx, http.MethodGet, o.GoogleApiUrl, nil)
	if err != nil {
		return nil, err
	}

	serverRequest.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := http.DefaultClient.Do(serverRequest)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("google api error: status %d, body: %s", resp.StatusCode, string(respBody))
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data GoogleUser

	err = json.Unmarshal(respBody, &data)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal google user: %w", err)
	}

	if data.Email == "" {
		return nil, fmt.Errorf("google user email is empty")
	}

	return &data, nil
}
