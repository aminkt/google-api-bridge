package bridges

import (
	"context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/idtoken"
)

func VerifyIdToken(ctx context.Context, idTokenString string, googleClientID string) (map[string]string, error) {
	// Create OAuth2 configuration
	config := oauth2.Config{
		ClientID: googleClientID,
		Endpoint: google.Endpoint,
	}

	// Verify the token
	TokenData, err := idtoken.Validate(ctx, idTokenString, config.ClientID)
	if err != nil {
		return nil, err
	}

	// Extract user data from the token's payload
	payload := TokenData.Claims
	result := make(map[string]string)
	result["email"] = payload["email"].(string)
	result["given_name"] = payload["given_name"].(string)
	result["family_name"] = payload["family_name"].(string)

	return result, nil
}
