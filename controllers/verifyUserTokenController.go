package controllers

import (
	"context"
	"fmt"
	"github.com/aminkt/google-api-bridge/lib"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	idToken "google.golang.org/api/idtoken"
	"net/http"
	"os"
)

func VerifyOathTokenAction(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token == "" {
		http.Error(w, "Authorization token is missing", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	userInfo, err := verifyGoogleToken(ctx, token, os.Getenv("GOOGLE_SERVER_CLIENT_ID"))
	if err != nil {
		lib.PrintStackTrace(err)
		http.Error(w, fmt.Sprintf("Failed to verify token: %v", err), http.StatusUnauthorized)
		return
	}

	// You can access user information from `userInfo` variable
	// For example, to get user's email:
	userEmail := userInfo["email"]
	fmt.Fprintf(w, "User email: %s", userEmail)
}

func verifyGoogleToken(ctx context.Context, token string, googleClientID string) (map[string]string, error) {
	// Create OAuth2 configuration
	config := oauth2.Config{
		ClientID: googleClientID,
		Endpoint: google.Endpoint,
	}

	// Verify the token
	idToken, err := idToken.Validate(ctx, token, config.ClientID)
	if err != nil {
		return nil, err
	}

	// Extract user data from the token's payload
	payload := idToken.Claims
	result := make(map[string]string)
	result["email"] = payload["email"].(string)
	result["name"] = payload["given_name"].(string)
	result["familyName"] = payload["family_name"].(string)

	return result, nil
}
