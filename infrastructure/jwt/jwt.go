package jwt

import (
	"crypto/rsa"
	"estrellitas-crud/infrastructure/env"
	"estrellitas-crud/infrastructure/logger"
	"estrellitas-crud/infrastructure/models"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"time"
)

type UserToken models.User

var (
	signKey *rsa.PrivateKey
)

// JWT personzalizado
type jwtCustomClaims struct {
	User      models.User `json:"user"`
	IPAddress string      `json:"ip_address"`
	jwt.StandardClaims
}

// init lee los archivos de firma y validación RSA
func init() {
	c := env.NewConfiguration()
	signBytes, err := ioutil.ReadFile(c.App.RSAPrivateKey)
	if err != nil {
		logger.Error.Printf("leyendo el archivo privado de firma: %s", err)
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		logger.Error.Printf("realizando el parse en authentication RSA private: %s", err)
	}
}

// Genera el token
func GenerateJWT(u models.User) (string, int, error) {
	c := &jwtCustomClaims{
		User:      u,
		IPAddress: "127.0.0.1",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 1200).Unix(),
			Issuer:    "Estrellitas",
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodRS256, c)
	token, err := t.SignedString(signKey)
	if err != nil {
		logger.Error.Printf("firmando el token: %v", err)
		return "", 70, err
	}
	// TODO encript Token
	return token, 29, nil
}
