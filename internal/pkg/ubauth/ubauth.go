package ubauth

import ubauth "github.com/ahmdyaasiin/ub-auth-without-notification/v2"

type IUBAuth interface {
	AuthUB(username, password string) (*ubauth.StudentDetails, error)
}

type ubAuth struct{}

func NewUBAuth() IUBAuth {
	return &ubAuth{}
}

func (ubAuth) AuthUB(username, password string) (*ubauth.StudentDetails, error) {
	return ubauth.AuthUB(username, password)
}
