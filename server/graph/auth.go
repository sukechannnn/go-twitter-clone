package graph

import (
	"context"
	"net/http"

	"github.com/sukechannnn/go-twitter-clone/graph/model"
	"gorm.io/gorm"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

// TODO: セッションカラムを追加して、User.id を直接見ないようにする
func validateAndGetUserID(db *gorm.DB, userId string) (*model.User, error) {
	user, err := model.FindUserById(db, userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Middleware decodes the share session cookie and packs the session into context
func Middleware(db *gorm.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c, err := r.Cookie("auth")

			// Allow unauthenticated users in
			if err != nil || c == nil {
				next.ServeHTTP(w, r)
				return
			}

			user, err := validateAndGetUserID(db, c.Value)
			if err != nil || user == nil {
				http.Error(w, "Invalid", http.StatusForbidden)
				return
			}

			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, user)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *model.User {
	raw, _ := ctx.Value(userCtxKey).(*model.User)
	return raw
}
