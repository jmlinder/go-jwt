package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

	// Signs the token with the singning key
	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

// Homepage setup
func homePage(w http.ResponseWriter, r *http.Request) {

	// Generates a new JWT and prints it to the console.
	validToken, err := generateJWT()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	// Sets the token generated to our HTTP Header
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:9000", nil)
	req.Header.Set("Token", validToken)
	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprintf(w, string(body))

	// Returns the valid token into the body of 9001.
	// fmt.Fprintf(w, validToken)
}

func handleRequests() {
	http.HandleFunc("/", homePage)

	log.Fatal(http.ListenAndServe(":9001", nil))
}

func main() {
	fmt.Println("Client started on port 9001")

	handleRequests()
}
