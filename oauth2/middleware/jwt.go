package middleware

import (
	"context"
	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
	adapter "github.com/gwatts/gin-adapter"
	"log"
	"net/http"
	"net/url"
	"time"
)

type CustomClaims struct {
	Scope string `json:"scope"`
}

// Validate does nothing for this example, but we need
// it to satisfy validator.CustomClaims interface.
func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}

func EnsureValidToken() gin.HandlerFunc {
	issuerUrl, _ := url.Parse("https://dev-qi1ij6arqhrlv7ds.us.auth0.com/")
	jwks.NewCachingProvider(issuerUrl, 5*time.Minute)
	jwtValidator, err := validator.New(func(ctx context.Context) (interface{}, error) {
		return []byte("YAEnfNNGgUvJOKK0QyEYSF7lWrJI0z3O"), nil
	}, validator.HS256, issuerUrl.String(), []string{"http://localhost:3002"}, validator.WithCustomClaims(
		func() validator.CustomClaims {
			return &CustomClaims{}
		},
	))
	if err != nil {
		log.Fatalf("Failed to set up the jwt validator")
	}
	errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("Encountered error while validating JWT: %v", err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message":"Failed to validate JWT."}`))
	}

	m := jwtmiddleware.New(jwtValidator.ValidateToken, jwtmiddleware.WithErrorHandler(errorHandler))
	return adapter.Wrap(m.CheckJWT)
}
