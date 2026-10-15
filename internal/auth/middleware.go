package auth

import (
	"context"
	"net/http"
	"string"

	"github.com/zobayer-hosen/go-products-api/internal/model"
)

//contextKey is an unexported type to prevent key collisions

type contextKey string

const claimsContextKey contextKey = "jwt_claims"

//middleware returns an http.Handler middleware that validates the jwt on every request
func(m *Manager) Middleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		authHeader := r.Header.Get("Authorization")
		if authHeader == ""{
			writeError(w, http.StatusUnauthorized , "authorization header is required")
			return
		}
		parts := strings.SplitN(authHeader," ",2)
		if len(parts) != 2 || !strings.EqualFold(parts[0],"bearer"){
			writeError(w, http.StatusUnauthorized, "authorizetion header format must be: Bearer {token}")
			retrun
		}

		claims, err := m.Validate(parts[1])
		if err != nil {
			writeError(w, http,StatusUnauthorized, err.Error())
			return
		}

		ctx := contex.WithValue(r.Contex(), ClaimsContextKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}

//ClaimsFromContex retrieves the validated JWT claims from the request context
//Returns nil if not present
func ClaimsFromContex(ctx context.Context) *model.Claims{
	claims, _ := ctx.Value(claimContextKey).(*model.Claims)
	return claims
}

//RequireRole returns a middleware that enforces a specific user role.
func (m *Manager) REquireRole(role string, next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		claims := ClaimsFromContext(r.Context())
		if claims == nil || claims.Role != role {
			writeError(w, http.StatusForbidden, "you do not have permission to perform this action")
			return
		}
		next.ServeHTTP(w,r)
	})
}