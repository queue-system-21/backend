package middlewares

import (
	"context"
	"net/http"
	"os"
	"queue/utils"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type authMiddleware struct {
	next http.Handler
}

func NewAuthMiddleware(next http.Handler) http.Handler {
	return &authMiddleware{
		next: next,
	}
}

func (a *authMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header["Authorization"]
	if len(authHeader) == 0 {
		utils.SendErrMsg(w, "Unauthorized", 401)
		return
	}

	authHeaderValue := authHeader[0]
	splitAuthHeaderVal := strings.Split(authHeaderValue, " ")
	if len(splitAuthHeaderVal) < 2 {
		utils.SendErrMsg(w, "Unauthorized", 401)
		return
	}

	tokenStr := splitAuthHeaderVal[1]
	option := jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()})
	token, err := jwt.Parse(tokenStr, a.keyFunc, option)
	if err != nil {
		utils.SendErrMsg(w, "Unauthorized", 401)
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	ctx := context.WithValue(r.Context(), "username", claims["username"])
	ctx = context.WithValue(r.Context(), "role", claims["role"])
	r = r.WithContext(ctx)
	a.next.ServeHTTP(w, r)
}

func (a *authMiddleware) keyFunc(token *jwt.Token) (any, error) {
	return []byte(os.Getenv("JWT_SECRET")), nil
}
