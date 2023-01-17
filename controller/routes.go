package controller

import (
	"net/http"
)

func Register() *http.ServeMux {
	mux := http.NewServeMux()

	return mux
}