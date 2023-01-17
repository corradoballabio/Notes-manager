package model

import (
	"../view"
)

func ReadAll() ([]view.NoteTable, error) {
	rows, err := con.Query("SELECT title, body FROM Note")

	if err != nil {
		return nil, err
	}

	notes := []view.NoteTable{}

	defer rows.Close()

	for rows.Next() {
		data := view.NoteTable{}
		rows.Scan(&data.Title, &data.Body)

		notes = append(notes, data)
	}

	return notes, nil
}

func ReadByTitle(title string) ([]view.NoteTable, error) {
	rows, err := con.Query("SELECT title, body FROM Note WHERE title = ?", title)

	if err != nil {
		return nil, err
	}

	notes := []view.NoteTable{}

	defer rows.Close()

	for rows.Next() {
		data := view.NoteTable{}
		rows.Scan(&data.Title, &data.Body)

		notes = append(notes, data)
	}

	return notes, nil
}