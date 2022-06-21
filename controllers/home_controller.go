package controllers

import (
	"net/http"

	"github.com/luannv1203/golang/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome to this Awesome API")
}
