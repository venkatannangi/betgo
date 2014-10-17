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


