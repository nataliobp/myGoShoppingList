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

func (r *SqliteUserRepository) findByEmail(email string) *User {
	s := fmt.Sprintf("select id, email, password from users where email = '%s'", email)
	rows, err := r.db.Query(s)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int64
		var email string
		var password string
		err = rows.Scan(&id, &email, &password)
		if err != nil {
			log.Fatal(err)
		}

		return &User{Id: id, Email: email, Password: password}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (r *SqliteUserRepository) findById(id int64) *User {
	s := fmt.Sprintf("select id, email, password from users where id = '%d'", id)
	rows, err := r.db.Query(s)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int64
		var email string
		var password string
		err = rows.Scan(&id, &email, &password)
		if err != nil {
			log.Fatal(err)
		}

		return &User{Id: id, Email: email, Password: password}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
