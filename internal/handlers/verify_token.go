package handlers

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/ibragimovtimurbek008/small-shop/internal/pkg/jwt"
)

func VerifyTokenHandler(w http.ResponseWriter, r *http.Request) {
	h := r.Header.Get("Authorization")

	tokStr := getTokenFromHeader(h)
	if tokStr == "" {
		slog.Error("bad token", "header", h)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := jwt.VerifyToken(tokStr)
	if err != nil {
		slog.Error("verify token", "err", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func getTokenFromHeader(authHeader string) string {
	if authHeader == "" {
		return ""
	}

	ss := strings.Split(authHeader, " ")
	if len(ss) >= 2 {
		if strings.ToLower(ss[0]) == "bearer" {
			return ss[1]
		}
		return ss[0]
	}

	return ""
}
