package responses

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JSON(c gin.ResponseWriter, statusCode int, data interface{}) {
	c.WriteHeader(statusCode)
	err := json.NewEncoder(c).Encode(data)
	if err != nil {
		fmt.Fprint(c, "%s", err.Error())
	}
}
func JSONPAGINATION(w http.ResponseWriter, statusCode int, data interface{}, pagination interface{}) {
	fmt.Println(w)
	w.WriteHeader(statusCode)
	w.Write([]byte("data"))
	w.Write([]byte("pagination"))
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprint(w, "%s", err.Error())
	}
	err = json.NewEncoder(w).Encode(pagination)
	if err != nil {
		fmt.Fprint(w, "%s", err.Error())
	}
}

func ERROR(c gin.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(c, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	JSON(c, http.StatusBadRequest, nil)
}
