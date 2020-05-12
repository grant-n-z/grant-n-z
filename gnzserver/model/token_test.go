package model

import (
	"strings"
	"testing"
)

func TestIsRefresh_True(t *testing.T) {
	tokenRequest := TokenRequest{
		GrantType: "refresh_token",
	}

	if !tokenRequest.IsRefresh() {
		t.Errorf("Incorrect TestIsRefresh_True test")
		t.FailNow()
	}
}

func TestIsRefresh_False(t *testing.T) {
	tokenRequest := TokenRequest{
		GrantType: "password",
	}

	if tokenRequest.IsRefresh() {
		t.Errorf("Incorrect TestIsRefresh_False test")
		t.FailNow()
	}
}

func TestTokenString(t *testing.T) {
	password := GrantPassword.String()
	if !strings.EqualFold(password, "password") {
		t.Errorf("Incorrect TestTokenString test")
		t.FailNow()
	}

	refreshToken := GrantRefreshToken.String()
	if !strings.EqualFold(refreshToken, "refresh_token") {
		t.Errorf("Incorrect TestString test")
		t.FailNow()
	}
}
