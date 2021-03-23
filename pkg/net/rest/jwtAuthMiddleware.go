package rest

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/nicholaslim94/messenger_backend/pkg/security"
)

//JwtAuthenticate is a middleware which validates http header's JWT token.
//Expects Authorization: Bearer <token>
//If tokenn is valid, add id from token to context
func JwtAuthenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headerToken := r.Header.Get("Authorization")
		if !strings.Contains(headerToken, "Bearer ") {
			log.Println(http.StatusText(http.StatusBadRequest))
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(http.StatusText(http.StatusBadRequest)))
			return
		}
		tokenString := strings.Split(headerToken, "Bearer ")[1]

		claims := &security.JwtClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			token.SigningString()
			return security.GetSecret(), nil
		})
		if err != nil || !token.Valid {
			log.Println(http.StatusText(http.StatusUnauthorized))
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
			return
		}
		ctx := context.WithValue(r.Context(), ContextIDKey, claims.ID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
