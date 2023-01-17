package controller

import (
	"fmt"
	"encoding/json"
	"net/http"
	"../model"
	"../view"
)
func crud() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			data := view.TodoTable{}
			json.NewDecoder(r.Body).Decode(&data)

			fmt.Println(data)

			if err := model.Create(data.Name, data.Todo); err != nil {
				fmt.Println(err)
				w.Write([]byte("Internal Error: "))
				w.Write([]byte(err.Error()))
				return
			}

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(&data)
		}
	}
}