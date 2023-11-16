package middleware

// https://auth0.com/docs/secure/tokens/json-web-tokens/create-custom-claims
type CustomClaims struct {
	Scope string `json:"scope"`
}
