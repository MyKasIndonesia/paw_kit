package jwt

import (
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

const (
	minimumExpiration = 3 * time.Second
)

type (
	// Tokenizer ...
	Tokenizer struct {
		Secret     []byte
		Expiration time.Duration
	}

	// jwtCustomClaims are custom claims extending default ones.
	jwtCustomClaims struct {
		jwtgo.StandardClaims

		Data interface{} `json:"data"`
	}
)

// New ...
func New(secret string, expiration time.Duration) (*Tokenizer, error) {
	if secret == "" {
		return nil, errors.New("secret can not be empty")
	}

	if expiration < minimumExpiration {
		return nil, errors.New("expiration time too short")
	}

	return &Tokenizer{
		Secret:     []byte(secret),
		Expiration: expiration,
	}, nil
}

// Generate ...
func (t *Tokenizer) Generate(data interface{}) (string, error) {
	// Set custom claims
	claims := &jwtCustomClaims{
		jwtgo.StandardClaims{
			ExpiresAt: time.Now().Add(t.Expiration).Unix(),
		},
		data,
	}

	// Create token with claims
	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	signed, err := token.SignedString(t.Secret)
	if err != nil {
		return "", errors.Wrap(err, "failed to sign token")
	}
	return signed, nil
}
