package identityAccess

import (
	"database/sql"
	"fmt"
	"log"
)

type SqliteUserRepository struct{ db *sql.DB }

func NewSqliteUserRepository(db *sql.DB) UserRepository {
	return &SqliteUserRepository{db}
}

func (r *SqliteUserRepository) save(user User) int64 {
	tx, err := r.db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into users(email, password) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Email, user.Password)

	if err != nil {
		log.Fatal(err)
	}

	lastId, _ := result.LastInsertId()

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	return lastId
}

func (r *SqliteUserRepository) findAll() []User {
	var users []User
	rows, err := r.db.Query("select email, password from users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var email string
		var password string
		err = rows.Scan(&email, &password)
		if err != nil {
			log.Fatal(err)
		}

		users = append(users, User{Email: email, Password: password})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return users
}

func (r *SqliteUserRepository) findByEmail(email string) (*User, error) {
	user := &User{}
	row := r.db.QueryRow(`SELECT id, email, password FROM users WHERE email = $1`, email)
	err := row.Scan(&user.Id, &user.Email, &user.Password)

	switch {
	case err == sql.ErrNoRows:
		return nil, fmt.Errorf("user not found")
	case err != nil:
		return nil, err
	}

	return user, nil
}

func (r *SqliteUserRepository) findById(id int64) (*User, error) {
	user := &User{}
	row := r.db.QueryRow(`SELECT id, email, password FROM users WHERE id = $1`, id)
	err := row.Scan(&user.Id, &user.Email, &user.Password)

	switch {
	case err == sql.ErrNoRows:
		return nil, fmt.Errorf("user not found")
	case err != nil:
		return nil, err
	}

	return user, nil
}
