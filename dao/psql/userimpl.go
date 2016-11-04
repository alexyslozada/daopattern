package psql

import "github.com/alexyslozada/daopattern/models"

type UserImplPsql struct {}

func (dao UserImplPsql) Create(u *models.User) error {
	query := "INSERT INTO users (firstname, lastname, email) VALUES ($1, $2, $3) RETURNING id"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	row := stmt.QueryRow(u.FirstName, u.LastName, u.Email)
	row.Scan(&u.ID)
	return nil
}

func (dao UserImplPsql) GetAll() ([]models.User, error) {
	query := "SELECT id, firstname, lastname, email FROM users"
	users := make([]models.User, 0)
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return users, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return users, err
	}

	for rows.Next() {
		var row models.User
		err := rows.Scan(&row.ID, &row.FirstName, &row.LastName, &row.Email)
		if err != nil {
			return users, err
		}
		users = append(users, row)
	}

	return users, nil
}

