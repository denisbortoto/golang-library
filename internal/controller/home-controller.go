package controller

import (
	"library/internal/utils"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	utils.ToJson(w, struct {
		Message string `json:"message"`
	}{
		Message: "Go RESTful Api",
	})
}
