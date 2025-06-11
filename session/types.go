package session

type CredentialsRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
