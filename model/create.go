package model

func Create(title, body string) error {
	insertQ, err := con.Query("INSERT INTO Note VALUES(?, ?)", title, body)

	if err != nil {
		return err
	}

	defer insertQ.Close()

	return nil
}