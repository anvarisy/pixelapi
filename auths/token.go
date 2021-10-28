package auths

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(username string, is_admin bool) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["auth_username"] = username
	claims["is_admin"] = is_admin
	claims["exp"] = time.Now().Add(time.Hour * 8760).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}

func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if err_token := token.Claims.Valid(); err_token != nil && !token.Valid {
		return err_token
	}
	return nil
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

//get the token from the request body
func ExtractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func ExtractTokenAuth(r *http.Request) (map[string]interface{}, error) {
	res := make(map[string]interface{})
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
	if ok && token.Valid {
		username, ok := claims["auth_username"].(string) //convert the interface to string
		if !ok {
			return nil, err
		}
		isAdmin, ok := claims["is_admin"].(bool)
		res["username"] = username
		res["is_admin"] = isAdmin
		return res, nil
	}
	return nil, err
}
