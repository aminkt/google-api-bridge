package main

import (
	"github.com/aminkt/google-api-bridge/httpHandlers"
	"github.com/aminkt/google-api-bridge/lib"
	"net/http"
)

func main() {
	verifyAppConfig()

	http.HandleFunc("/verify-id-token", httpHandlers.VerifyIdTokenHandler)
	http.ListenAndServe(lib.ReadEnvironmentVariables().AppServerAddress, nil)
}

func verifyAppConfig() {
	envVars := lib.ReadEnvironmentVariables()

	if envVars.GoogleClientId == "" {
		panic("GoogleClientId is not configured!")
	}

	if envVars.AllowedIps == "" {
		panic("This application is not safe for open ip environment. You need to whitelist some ips.")
	}
}
