package repository

import (
	"database/sql"
	"github.com/JohanAanesen/CSAMS/webservice/model"
)

// UserPendingRepository struct
type UserPendingRepository struct {
	db *sql.DB
}

// NewUserPendingRepository func
func NewUserPendingRepository(db *sql.DB) *UserPendingRepository {
	return &UserPendingRepository{
		db: db,
	}
}

// Insert inserts a new userPending in the db
func (repo *UserPendingRepository) Insert(pending model.UserRegistrationPending) (int, error) {
	var id int64

	query := "INSERT INTO `users_pending` (`name`, `email`, `password`, `validation_id`) VALUES (?, ?, ?, ?)"

	tx, err := repo.db.Begin()
	if err != nil {
		return int(id), err
	}

	rows, err := tx.Exec(query, pending.Name, pending.Email, pending.Password, pending.ValidationID)
	if err != nil {
		tx.Rollback()
		return int(id), err
	}

	id, err = rows.LastInsertId()
	if err != nil {
		tx.Rollback()
		return int(id), err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return int(id), err
	}

	return int(id), nil
}

// InsertNewEmail inserts a new userPending for emails in the db
func (repo *UserPendingRepository) InsertNewEmail(pending model.UserRegistrationPending) (int, error) {
	var id int64

	query := "INSERT INTO `users_pending` (`email`, `validation_id`) VALUES (?, ?)"

	tx, err := repo.db.Begin()
	if err != nil {
		return int(id), err
	}

	rows, err := tx.Exec(query, pending.Email, pending.ValidationID)
	if err != nil {
		tx.Rollback()
		return int(id), err
	}

	id, err = rows.LastInsertId()
	if err != nil {
		tx.Rollback()
		return int(id), err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return int(id), err
	}

	return int(id), nil
}

// FetchAll fetches all userPending in the db, but not the password
func (repo *UserPendingRepository) FetchAll() ([]*model.UserRegistrationPending, error) {
	result := make([]*model.UserRegistrationPending, 0)

	query := "SELECT `id`, `name`, `email`, `validation_id` FROM `users_pending` "

	rows, err := repo.db.Query(query)
	if err != nil {
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		temp := model.UserRegistrationPending{}

		err = rows.Scan(&temp.ID, &temp.Name, &temp.Email, &temp.ValidationID)
		if err != nil {
			return result, err
		}

		result = append(result, &temp)
	}

	return result, err
}

// FetchPassword fetches the password to one user through the id
func (repo *UserPendingRepository) FetchPassword(id int) (string, error) {
	var password string

	query := "SELECT `password` FROM `users_pending` WHERE `id` = ?"

	rows, err := repo.db.Query(query, id)
	if err != nil {
		return password, err
	}

	defer rows.Close()

	if rows.Next() {

		err = rows.Scan(&password)
		if err != nil {
			return password, err
		}

	}

	return password, nil

}
