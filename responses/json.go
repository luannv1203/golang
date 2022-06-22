package responses

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JSON(c *gin.Context, statusCode int, data interface{}, message string) {
	c.JSON(http.StatusOK, gin.H{
		"status":  statusCode,
		"message": message,
		"data":    data,
	})
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

func ERROR(c *gin.Context, statusCode int, err error) {
	if err != nil {
		c.JSON(statusCode, gin.H{
			"status":  statusCode,
			"message": "",
			"data": struct {
				Error string `json:"error"`
			}{
				Error: err.Error(),
			},
		})
		return
	}
	JSON(c, http.StatusBadRequest, nil, "")
}
