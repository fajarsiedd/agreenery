package entities

type User struct {
	Base
	DisplayName  string
	Photo        string
	Phone        string
	CredentialID string
	Credential   Credential
	AccessToken  string
	RefreshToken string
}
