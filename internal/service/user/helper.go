package user

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/Beep-Technologies/beepbeep3-ocpp/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func HashAPIKey(apiKey string) (string, error) {
	hasher := sha256.New()
	hasher.Write([]byte(apiKey))
	hashedAPiKey := hex.EncodeToString(hasher.Sum(nil))
	return hashedAPiKey, nil
}

func EncryptPw(password string) (encPass string, err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return
	}

	return string(hashedPassword), nil
}

func VerifyPw(username, password string, userObj models.User) (isLoginSuccess bool, user *models.User) {
	err := bcrypt.CompareHashAndPassword([]byte(userObj.Password), []byte(password))
	if err != nil {
		fmt.Println("Incorrect password")
		return
	} else {
		isLoginSuccess = true
		user = &userObj
		fmt.Println("Login successful!")
	}

	return isLoginSuccess, user
}
