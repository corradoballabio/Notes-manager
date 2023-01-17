package model

func Create(name, todo string) error {
	insertQ, err := con.Query("INSERT INTO TODO VALUES(?, ?)", name, todo)

	if err != nil {
		return err
	}

	defer insertQ.Close()

	return nil
}