package auth

import (
	"context"

	"github.com/00Duck/wishr-api/models"
	"golang.org/x/crypto/bcrypt"
)

// Key is used specifically here as a key for the request context
type Key int

/*
SessionKey is the arbitrarily defined key for the user claim. Specifically, a type "Key" with an integer value 0 is used to identify
the user's key, which is the unique property set to find the user claim in a context. In order to find a key in the context, its value
and type must match exactly which is why the key type is defined in the auth package.
*/
const SessionKey Key = 0

// SessionCookieName is called by various functions to identify which cookie is used for the session
const SessionCookieName = "sessionID"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// NewContext creates a new context with the SessionKey associated to its corresponding Key type
func NewContext(ctx context.Context, session *models.Session) context.Context {
	return context.WithValue(ctx, SessionKey, session)
}

// FromContext returns the value stored in the SessionKey, should be jwt.MapClaims
func FromContext(ctx context.Context) *models.Session {
	return ctx.Value(SessionKey).(*models.Session)
}
