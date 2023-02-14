package auth

import (
	"context"
	"github.com/MicahParks/keyfunc"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"net/http"
	"time"
)

// middleware for jwt authentication and authorization with cognito user pool
// cognito user sub is stored in context
func Middleware(next http.Handler, jwtToken string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get cognito user sub from jwt token
		jwksURL := "https://cognito-idp.us-east-1.amazonaws.com/us-east-1_XXXXXXXXX/.well-known/jwks.json"
		options := keyfunc.Options{
			RefreshErrorHandler: func(err error) {
				log.Printf("There was an error with the jwt.Keyfunc\nError: %s", err.Error())
			},
			RefreshInterval:   time.Hour,
			RefreshRateLimit:  time.Minute * 5,
			RefreshTimeout:    time.Second * 10,
			RefreshUnknownKID: true,
		}
		jwks, err := keyfunc.Get(jwksURL, options)
		token, err := jwt.Parse(jwtToken, jwks.Keyfunc)
		if err != nil {
			log.Printf("Failed to parse the JWT.\nError: %s", err.Error())
			return
		}
		if !token.Valid {
			log.Printf("The JWT is not valid.")
			return
		}
		// get cognito user sub from token
		sub := token.Claims.(jwt.MapClaims)["sub"]

		// set cognito user sub in context
		ctx := r.Context()
		ctx = context.WithValue(ctx, "sub", sub)
		r = r.WithContext(ctx)

		jwks.EndBackground()
		// and call next handler
		next.ServeHTTP(w, r)
	})
}
