package postgres

import (
	"GravitumTest/internal/model"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

func NewUser(user model.User) (model.User, error) {
	res := model.User{
		FullName: user.FullName,
	}
	query := fmt.Sprintf(`INSERT INTO users (FullName) VALUES ('%s') RETURNING ID`, *user.FullName)
	err := DB.QueryRow(query).Scan(&res.Id)
	if err != nil {
		log.Println(err)
		return model.User{}, err
	}
	return res, nil
}
func GetUser(id int) (model.User, error) {
	var user model.User
	query := fmt.Sprintf(`select * from users where id = %d`, id)
	row := DB.QueryRow(query)
	err := row.Scan(
		&user.Id,
		&user.FullName,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return model.User{}, errors.New("User not found")
	} else if err != nil {
		return model.User{}, err
	}
	return user, nil
}
func UpdateUser(id int, user model.User) error {
	query := fmt.Sprintf(`UPDATE users SET FullName = '%s' WHERE id = %d`, *user.FullName, id)
	result, err := DB.Exec(query)
	if err != nil {
		fmt.Println(err)
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		log.Printf("User with id - %d Not found", id)
		return errors.New("User not found")
	} else {
		return nil
	}

}
