package db

import "github.com/gofrs/uuid"

func AddProduct(name, description, category, price, image, total string,) error {
	productId, err := uuid.NewV4()
	if err != nil {
		return err
	}
	id := productId.String()[:8]

	stm := `INSERT INTO products(id, title, price, description, category, image, total) VALUES(?,?,?,?,?,?,?)`

	if _, err = DB.Exec(stm, id, name, price, description, category, image, total); err != nil {
		return err
	}

	return nil
}