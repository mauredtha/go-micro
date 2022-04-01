package models

import (
	"database/sql"
	"microservices/libraries/api"
)

// User : struct of User
type User struct {
	ID       uint
	Username string
	Password string
	Email    string
	IsActive bool
}

const qUser = `SELECT id, username, password, email, is_active FROM users`

// List of users
func (u *User) List(db *sql.DB) ([]User, error) {
	var list []User

	rows, err := db.Query(qUser)
	if err != nil {
		return list, err
	}

	defer rows.Close()

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.IsActive); err != nil {
			return list, err
		}
		list = append(list, user)
	}

	return list, rows.Err()
}

// Create new user
func (u *User) Create(db *sql.DB) error {
	const query = `
        INSERT INTO users (username, password, email, is_active, created, updated)
        VALUES (?, ?, ?, 0, NOW(), NOW())
    `
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	res, err := stmt.Exec(u.Username, u.Password, u.Email)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	u.ID = uint(id)

	return nil
}

// Get user by id
func (u *User) Get(db *sql.DB) error {
	const q string = `SELECT id, username, password, email, is_active FROM users`
	err := db.QueryRow(q+" WHERE id=?", u.ID).Scan(&u.ID, &u.Username, &u.Password, &u.Email, &u.IsActive)

	if err == sql.ErrNoRows {
		err = api.ErrNotFound(err, "")
	}

	return err
}

// Update user by id
func (u *User) Update(db *sql.DB) error {
	const q string = `UPDATE users SET is_active = ? WHERE id = ?`
	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(u.IsActive, u.ID)
	return err
}

// Delete user by id
func (u *User) Delete(db *sql.DB) error {
	const q string = `DELETE FROM users WHERE id = ?`
	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(u.ID)
	return err
}
