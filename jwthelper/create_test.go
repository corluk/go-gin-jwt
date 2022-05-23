package jwthelper

import "testing"

func TestCreateToken(t *testing.T) {

	secret := "12345"
	jwtHelper := New(secret)
	token, err := jwtHelper.SetValue("user", "testuser").Create()
	if err != nil {
		t.Log(err.Error())
		t.Fail()
		t.FailNow()
	}
	t.Log(token)
}
