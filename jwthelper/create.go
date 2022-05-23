package jwthelper

import (
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
type JWTHelper struct {
	Claims *Claims
	Secret string
	Token  string
}

func New(secret string) JWTHelper {

	return JWTHelper{
		Claims: &jwt.MapClaims{},
		Secret: secret,
	}

}

func (jwtHelper *JWTHelper) SetExpire(expires int64) *JWTHelper {

	(*jwtHelper.Claims)["exp"] = expires
	return jwtHelper
}

func (jwtHelper *JWTHelper) SetValue(key string, value string) *JWTHelper {
	(*jwtHelper.Claims)[key] = value
	return jwtHelper
}

func (jwtHelper *JWTHelper) Build() (string, error) {

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtHelper.Claims)
	token, err := at.SignedString([]byte(jwtHelper.Secret))
	if err != nil {
		return "", err

	}
	jwtHelper.Token = token
	return token, nil
}

func (jwtHelper *JWTHelper) Create() (string, error) {

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtHelper.Claims)
	token, err := at.SignedString([]byte(jwtHelper.Secret))
	if err != nil {
		return "", err

	}
	jwtHelper.Token = token
	return token, nil
}
