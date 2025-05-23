package model

import (
	"encoding/json"
	"net/http"
)

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
}

// func SendResponse(w http.ResponseWriter, httpCode int, code int, message string, data interface{}) {
// 	w.Header().Set("Content-Type", "application/json")
// 	// w.WriteHeader(httpCode) // go 默认返回 200
// 	json.NewEncoder(w).Encode(&Result{
// 		Code:    code,
// 		Message: message,
// 		Data:    data,
// 	})
// }

func Success(w http.ResponseWriter, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&Result{
		Code:    1,
		Message: message,
		Data:    data,
	})
}

func Error(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&Result{
		Code:    0,
		Message: message,
	})
}
