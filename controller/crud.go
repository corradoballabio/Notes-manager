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
		} else if r.Method == http.MethodGet {
			var data []view.TodoTable
			var err error

			name := r.URL.Query().Get("name")

			if len(name) < 1 {
				data, err = model.ReadAll()
			} else {
				data, err = model.ReadByName(name[:])
			}

			if err != nil {
				fmt.Println(err)
				w.Write([]byte("Internal Error: "))
				w.Write([]byte(err.Error()))
				return
			}

			json.NewEncoder(w).Encode(data)
		} else if r.Method == http.MethodDelete {
			name := r.URL.Path[1:]

			if len(name) < 1 {
				w.Write([]byte("Error: no id specified in a delete call"))
				return
			}

			if err := model.Delete(name); err != nil {
				fmt.Println(err)
				w.Write([]byte("Internal Error: "))
				w.Write([]byte(err.Error()))
				return
			}

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(struct{
				Status string `json:"status"`
			}{"Item Deleted"})
		}
	}
}