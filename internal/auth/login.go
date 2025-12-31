package auth

import (
	"context"
	"time"

	"github.com/go-rod/rod"
)

type AuthService struct {
	Page *rod.Page
}

func NewAuthService(page *rod.Page) *AuthService {
	return &AuthService{Page: page}
}

func (a *AuthService) Login(email, password string) error {
	page := a.Page

	// Navigate to LinkedIn login
	page.MustNavigate("https://www.linkedin.com/login")

	// Wait for page load
	page.MustWaitLoad()

	// Fill email
	page.MustElement(`#username`).MustInput(email)

	// Small human delay
	time.Sleep(2 * time.Second)

	// Fill password
	page.MustElement(`#password`).MustInput(password)

	// Click login
	page.MustElement(`button[type="submit"]`).MustClick()

	// Wait for navigation
	page.MustWaitNavigation()

	return nil
}
