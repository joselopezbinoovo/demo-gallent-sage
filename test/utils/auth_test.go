package utils

import (
	"testing"
	"vg-sage/internal/utils"
	"github.com/dgrijalva/jwt-go"
)

func TestCreateToken(t *testing.T) {
	token := utils.CreateToken(jwt.MapClaims{
		"password": "nain",
	})

	if token == "" {
		t.Error("Null token generated")
	}
	t.Log(token)
}

func TestValidateToken(t *testing.T) {
	token := utils.CreateToken(jwt.MapClaims{
		"password": "password",
	})

	if token == "" {
		t.Error("Null token generated")
	} else {
		t.Log("valid token", token)
	}

	tokenValid := utils.ValidateToken(token)
	if !tokenValid {
		t.Error("Not a valid token")
	}

	notAToken := utils.ValidateToken("not a token")
	if notAToken {
		t.Error("Is returning a valid token")
	}

	validTokenNoPassword := utils.CreateToken(jwt.MapClaims{
		"password": "not valid pass",
	})
	if utils.ValidateToken(validTokenNoPassword) {
		t.Error("Validating with wrong password")
	}
}
