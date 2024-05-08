package httpHandlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aminkt/google-api-bridge/bridges"
	"github.com/aminkt/google-api-bridge/lib"
	"net/http"
)

func VerifyIdTokenHandler(w http.ResponseWriter, r *http.Request) {
	isOk, err := lib.VerifyIsAllowedIPs(r)
	if !isOk {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	token := r.URL.Query().Get("token")
	if token == "" {
		http.Error(w, "Authorization token is missing", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	userInfo, err := bridges.VerifyIdToken(ctx, token, lib.ReadEnvironmentVariables().GoogleClientId)
	if err != nil {
		lib.PrintStackTrace(err)
		http.Error(w, fmt.Sprintf("Failed to verify token: %v", err), http.StatusUnauthorized)
		return
	}

	// You can access user information from `userInfo` variable
	// For example, to get user's email:
	//userEmail := userInfo["email"]

	// Set content type header
	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(userInfo)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
	// Write the JSON data to the response
	_, err = w.Write(jsonData)
	if err != nil {
		fmt.Println("Failed to write JSON to response:", err)
	}
}
