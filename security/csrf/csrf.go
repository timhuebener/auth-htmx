package csrf

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"time"
)

const CookieName = "csrf_token"

type ctxKey struct{}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie(CookieName)
		token := ""
		if err == nil && c.Value != "" {
			token = c.Value
		} else {
			b := make([]byte, 32)
			_, _ = rand.Read(b)
			token = base64.RawURLEncoding.EncodeToString(b)
			http.SetCookie(w, &http.Cookie{
				Name:     CookieName,
				Value:    token,
				Path:     "/",
				Expires:  time.Now().Add(24 * time.Hour),
				HttpOnly: true,
				Secure:   true,
				SameSite: http.SameSiteLaxMode,
			})
		}

		if r.Method != http.MethodGet && r.Method != http.MethodHead && r.Method != http.MethodOptions {
			hdr := r.Header.Get("X-CSRF-Token")
			if hdr == "" || hdr != token {
				http.Error(w, "csrf error", http.StatusUnauthorized)
				return
			}
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKey{}, token)))
	})
}

func Token(r *http.Request) string {
	if v, ok := r.Context().Value(ctxKey{}).(string); ok && v != "" {
		return v
	}
	c, err := r.Cookie(CookieName)
	if err != nil {
		return ""
	}
	return c.Value
}
