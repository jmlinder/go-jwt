package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// The signing key is usually set in an environment variable:
// var mySigningKey = os.Get("MY_JWT")
var mySigningKey = []byte("mysecretphrase")

// REST API - Homepage
// Generates a new JWT upon request
func homePage(w http.ResponseWriter, r *http.Request) {
	validToken, err := generateJWT()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w, validToken)
}

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

// Handles requetsts to "localhost:9001/"
func handleRequests() {
	http.HandleFunc("/", homePage)

	log.Fatal(http.ListenAndServe(":9001", nil))
}
func main() {
	fmt.Println("My Simple Client")

	handleRequests()
}
