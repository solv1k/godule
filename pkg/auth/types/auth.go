package types

// Entity for authentication and authorization
type AuthEntity interface {
	GetAuthID() string
	GetAuthCodeType() string
	GetAuthCodeIdentifier() string
	GetAuthPayload() map[string]interface{}
}
