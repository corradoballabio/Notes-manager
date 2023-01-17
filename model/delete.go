package model

func Delete(name string) error {
	deletion, err := con.Query("DELETE FROM TODO WHERE name = ?", name)

	if err != nil {
		return err
	}

	defer deletion.Close()

	return nil
}