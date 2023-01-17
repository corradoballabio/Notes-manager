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
			data := view.NoteTable{}
			json.NewDecoder(r.Body).Decode(&data)

			fmt.Println(data)

			if err := model.Create(data.Title, data.Body); err != nil {
				fmt.Println(err)
				w.Write([]byte("Internal Error: "))
				w.Write([]byte(err.Error()))
				return
			}

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(&data)
		} else if r.Method == http.MethodGet {
			var data []view.NoteTable
			var err error

			title := r.URL.Query().Get("title")

			if len(title) < 1 {
				data, err = model.ReadAll()
			} else {
				data, err = model.ReadByTitle(title[:])
			}

			if err != nil {
				fmt.Println(err)
				w.Write([]byte("Internal Error: "))
				w.Write([]byte(err.Error()))
				return
			}

			json.NewEncoder(w).Encode(data)
		} else if r.Method == http.MethodDelete {
			title := r.URL.Path[1:]

			if len(title) < 1 {
				w.Write([]byte("Error: no id specified in a delete call"))
				return
			}

			if err := model.Delete(title); err != nil {
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