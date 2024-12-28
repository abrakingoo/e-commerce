package db

func CreateTables() error {
	for _, stm := range Statements {
		if _, err := DB.Exec(stm); err != nil {
			return err
		}
	}
	return nil
}