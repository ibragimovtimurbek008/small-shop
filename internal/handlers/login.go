package handlers

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/ibragimovtimurbek008/small-shop/internal/pkg/jwt"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	token, err := jwt.CreateToken("timur")
	if err != nil {
		slog.Error("create token", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, token)
}
