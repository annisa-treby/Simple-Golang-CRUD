package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleSuccess(resp http.ResponseWriter, status int, data interface{})  {
	response := Respond{
		Status:  true,
		Message: "Success",
		Data:    data,
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(status)

	err := json.NewEncoder(resp).Encode(response)
	if err!= nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("wow amazing error"))
		fmt.Printf("[HandleSuccess] Error when encode data with error :%w\n", err.Error())
	}
}

func HandleError(resp http.ResponseWriter, status int, msg string)  {

	response := Respond{
		Status:  false,
		Message: msg,
		Data:    nil,
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(status)

	err := json.NewEncoder(resp).Encode(response)
	if err!=nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("amazing error"))
		fmt.Printf("[HandleError<] Error when encode data with error :%w\n", err.Error())
	}
}
