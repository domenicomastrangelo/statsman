package authenticate

import (
	"encoding/json"
	"io"
	"net/http"
)

func Auth(w http.ResponseWriter, r *http.Request) {
	creds := credentials{}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		w.Write([]byte("An error occurred reading the body of the request"))
		return
	}

	err = json.Unmarshal(body, &creds)

	if err != nil {
		w.Write([]byte("JSON request not correctly formatted"))
		return
	}

	if !creds.checkCredentials() {
		w.Write([]byte("Unauthorized"))
		return
	}

	authResponseHeader, err := authenticateJWT(creds)

	if err != nil {
		w.Write([]byte("Unauthorized"))
		return
	}

	w.Header().Add("authentication", authResponseHeader)
	w.Write([]byte("Authorized"))
}
