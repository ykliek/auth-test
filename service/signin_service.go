package service

import "api-auth-test/auth"

type signInInterface interface {
	SignIn(auth.Details) (string, error)
}

type signInStruct struct {}

var (
	Authorize signInInterface = &signInStruct{}
)

func (si *signInStruct) SignIn(authD auth.Details) (string, error) {
	token, err := auth.CreateToken(authD)
	if err != nil {
		return "", err
	}
	return token , err
}
