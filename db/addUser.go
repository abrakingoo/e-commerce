package db


func AddUser(uid, fname, lname, number, email, password string) error {
	stm := `INSERT INTO users(id, firstname, lastname, phonenumber, email, password)
	VALUES(?,?,?,?,?,?);`
	_, err := DB.Exec(stm, uid, fname, lname, number, email, password)
	if err != nil {
		return err
	}
	return nil
}