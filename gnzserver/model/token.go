package model

const (
	GrantPassword GrantTypeConfig = iota
	GrantRefreshToken
)

// Token request
type TokenRequest struct {
	Password     string `json:"password"`
	Email        string `json:"email"`
	GrantType    string `json:"grant_type"`
	RefreshToken string `json:"refresh_token"`
}

// Token response
type TokenResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func (t TokenRequest) IsRefresh() bool {
	if t.GrantType == GrantRefreshToken.String() {
		return true
	}
	return false
}

// Group table config struct
type GrantTypeConfig int

func (gc GrantTypeConfig) String() string {
	switch gc {
	case GrantPassword:
		return "password"
	case GrantRefreshToken:
		return "refresh_token"
	}
	return ""
}
