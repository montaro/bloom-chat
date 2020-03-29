package auth

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/schema"
	"github.com/hackebrot/go-repr/repr"
)

type Claims struct {
	AtHash        string  `json:"at_hash,omitempty"`
	EmailVerified bool    `json:"email_verified,omitempty"`
	EventId       string  `json:"event_id,omitempty"`
	TokenUse      string  `json:"token_use,omitempty"`
	AuthTime      int64   `json:"auth_time,omitempty"`
	Username      string  `json:"cognito:username,omitempty"`
	Email         string  `json:"email,omitempty"`
	jwt.StandardClaims
}

type UserTokens struct {
	IdToken     string `schema:"id_token"`
	AccessToken string `schema:"access_token"`
	ExpiresIn   string `schema:"expires_in"`
	TokenType   string `schema:"token_type"`
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userTokens := new(UserTokens)
	if err := schema.NewDecoder().Decode(userTokens, r.Form); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	claims := &Claims{}

	_, err = jwt.Parse(userTokens.IdToken, func(token *jwt.Token) (interface{}, error) {
		if claimsMap, ok := token.Claims.(jwt.MapClaims); ok {
			fmt.Println(repr.Repr(claimsMap))
			claims.Username = claimsMap["cognito:username"].(string)
			claims.Email = claimsMap["email"].(string)
		}
		return nil, nil
	})

	//tkn, err := jwt.ParseWithClaims(userTokens.IdToken, claims, func(token *jwt.Token) (interface{}, error) {
	//	return jwtKey, nil
	//})
	//if err != nil {
	//	if err == jwt.ErrSignatureInvalid {
	//		w.WriteHeader(http.StatusUnauthorized)
	//		return
	//	}
	//	w.WriteHeader(http.StatusBadRequest)
	//	return
	//}
	//if !tkn.Valid {
	//	w.WriteHeader(http.StatusUnauthorized)
	//	return
	//}

	// Finally, return the welcome message to the user, along with their
	// username given in the token
	_, _ = w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))
}
