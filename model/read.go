package model

import (
	"../view"
)

func ReadAll() ([]view.TodoTable, error) {
	rows, err := con.Query("SELECT name, todo FROM TODO")

	if err != nil {
		return nil, err
	}

	todos := []view.TodoTable{}

	defer rows.Close()

	for rows.Next() {
		data := view.TodoTable{}
		rows.Scan(&data.Name, &data.Todo)

		todos = append(todos, data)
	}

	return todos, nil
}

func ReadByName(name string) ([]view.TodoTable, error) {
	rows, err := con.Query("SELECT name, todo FROM TODO WHERE Name = ?", name)

	if err != nil {
		return nil, err
	}

	todos := []view.TodoTable{}

	defer rows.Close()

	for rows.Next() {
		data := view.TodoTable{}
		rows.Scan(&data.Name, &data.Todo)

		todos = append(todos, data)
	}

	return todos, nil
}