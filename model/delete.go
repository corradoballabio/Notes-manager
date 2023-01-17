package model

func Delete(title string) error {
	deletion, err := con.Query("DELETE FROM Note WHERE title = ?", title)

	if err != nil {
		return err
	}

	defer deletion.Close()

	return nil
}