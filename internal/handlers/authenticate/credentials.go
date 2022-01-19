package authenticate

type credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c *credentials) checkCredentials() bool {
	// Check username and password against
	// a saved record

	return true
}
