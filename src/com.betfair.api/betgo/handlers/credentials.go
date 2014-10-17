package handlers


import (
	"net/http"
)

func getApiCredentials(r *http.Request) (appKey string, sessionToken string) {
	appKeyCookie,_ := r.Cookie("appKey")
	if appKeyCookie != nil {
		appKey = appKeyCookie.Value
	}

	sessionTokenCookie,_ := r.Cookie("ssoid")
	if sessionTokenCookie != nil {
		sessionToken = sessionTokenCookie.Value
	}

	return appKey, sessionToken
}

func addCredentialToHeader(r *http.Request, appKey string, sessionToken string) {
	r.Header.Add("X-Application", appKey)
	r.Header.Add("X-Authentication", sessionToken)
}
