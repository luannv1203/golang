package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprint(w, "%s", err.Error())
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

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}
