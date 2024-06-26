package middlewares

import (
	"context"
	"net/http"
	"encoding/json"
	"github.com/ProgramandoComAndre/holistic-therapy-clientes2/src/infra/auth"
)


type errorBody struct {
	Message string `json:"message"`
}

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")

        if token == "" {
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusUnauthorized)
            json.NewEncoder(w).Encode(&errorBody{Message: "Unauthorized"})
            return
        }

        tokenWithoutBearer := token[7:]
        authorizedUser, err := auth.VerifyToken(tokenWithoutBearer)

        if err != nil {
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusUnauthorized)
            json.NewEncoder(w).Encode(&errorBody{Message: err.Error()})
            return
        }

		if(authorizedUser.Roleid <= 1) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(&errorBody{Message: "Operation not allowed"})
			return
		}

		ctx := context.WithValue(r.Context(), "AuthorizedUser", authorizedUser)
		r = r.WithContext(ctx)
        // If authentication is successful, call the next handler
        next.ServeHTTP(w, r)
    })
}