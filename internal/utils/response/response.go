package response

import "net/http"

func WriteJson(w http.ResponseWriter,status int,data interface {}) error{
	w.Header().set("Content-Type","application/json")
	w.w
}