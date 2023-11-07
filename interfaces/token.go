package interfaces

type Token struct {
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	Iss           string `json:"iss"`
	Aud           string `json:"aud"`
	AuthTime      int    `json:"auth_time"`
	UserID        string `json:"user_id"`
	Sub           string `json:"sub"`
	Iat           int    `json:"iat"`
	Exp           int    `json:"exp"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Firebase      struct {
		Identities struct {
			GoogleCom []string `json:"google.com"`
			Email     []string `json:"email"`
		} `json:"identities"`
		SignInProvider string `json:"sign_in_provider"`
	} `json:"firebase"`
}
