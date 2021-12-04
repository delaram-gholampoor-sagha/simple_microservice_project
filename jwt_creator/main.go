package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// "github.com/dgrijalva/jwt-go"

var MySigningKey = []byte(os.Getenv("SECRET_KEY"))

func GetJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	// now you can have multiple claims , claim is basically the information you want to add
	claims["authorized"] = true
	claims["client"] = "delarammajestic"
	claims["aud"] = "billing.jwt.io"
	claims["iss"] = "jwtgo.io"
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()
	tokenString, err := token.SignedString(MySigningKey)
	if err != nil {
		fmt.Errorf("something went wrong : %s", err.Error())
		return "", err

	}

	return tokenString, nil
}

func Index(w http.ResponseWriter, r *http.Request) {
	// when you hit the "/" route you hit the indx function and index function call the fetjwt function
	validtoken, err := GetJWT()
	fmt.Println(validtoken)
	if err != nil {
		fmt.Println("failed to generate the token ")

	}
	fmt.Fprintf(w, validtoken)
}

func handleRequests() {
	http.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
