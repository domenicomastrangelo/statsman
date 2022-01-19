package authenticate

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/domenicomastrangelo/statsman/internal/database"
)

func Auth(w http.ResponseWriter, r *http.Request) {
	creds := database.Credentials{}

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

	if !creds.CheckCredentials() {
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

func Validate(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] == nil {
			w.Write([]byte("Unauthorized"))
		} else {
			token := r.Header["Token"][0]

			jwtToken, err := validateJWT(token)

			if err != nil {
				w.Write([]byte("Unauthorized"))
			} else {
				w.Header().Add("authentication", jwtToken)

				h.ServeHTTP(w, r)
			}
		}
	})
}
