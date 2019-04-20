package main

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// The signing key is usually set in an environment variable:
// var mySigningKey = os.Get("MY_JWT")
var mySigningKey = []byte("mysecretphrase")

/* JWTs are used in SPAs (single page apps) for authorization, allowing the user to transfer data betwen
the client and the server using HMAC or public/private keypair encryption */
func generateJWT() (string, error) {

	// generates a signed token using HMAC SHA 256 hash algorithm
	token := jwt.New(jwt.SigningMethodHS256)

	// Generates an authorization for the user that expires after 30 minutes
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "Joseph Linder"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	//
	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func main() {
	fmt.Println("My Simple Client")

	// Generates a new JWT and prints it to the console.
	tokenString, err := generateJWT()
	if err != nil {
		fmt.Println("Error generating generating token string")
	}
	fmt.Println(tokenString)
}
