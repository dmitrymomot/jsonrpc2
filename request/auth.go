package request

import "github.com/dmitrymomot/go-jwt"

type (
	// Auth interface
	Auth interface {
		GetAccessToken() string
		GetUserID() string
		SetUserID(uid string)
		GetUserRole() string
		SetUserRole(role string)
		GetAppID() string
		SetAppID(aid string)
		GetClaims() jwt.Claims
		SetClaims(claims jwt.Claims)
	}

	// AuthArgs struct
	AuthArgs struct {
		AccessToken string `json:"__access_token"`
		userID      string
		userRole    string
		appID       string
		claims      jwt.Claims
	}
)

// GetAccessToken ...
func (a *AuthArgs) GetAccessToken() string {
	return a.AccessToken
}

// GetUserID ...
func (a *AuthArgs) GetUserID() string {
	return a.userID
}

// SetUserID ...
func (a *AuthArgs) SetUserID(uid string) {
	a.userID = uid
}

// GetUserRole ...
func (a *AuthArgs) GetUserRole() string {
	return a.userRole
}

// SetUserRole ...
func (a *AuthArgs) SetUserRole(role string) {
	a.userRole = role
}

// GetAppID ...
func (a *AuthArgs) GetAppID() string {
	return a.appID
}

// SetAppID ...
func (a *AuthArgs) SetAppID(aid string) {
	a.appID = aid
}

// GetClaims ...
func (a *AuthArgs) GetClaims() jwt.Claims {
	return a.claims
}

// SetClaims ...
func (a *AuthArgs) SetClaims(claims jwt.Claims) {
	a.claims = claims
}
