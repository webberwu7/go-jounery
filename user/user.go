package user

import (
	"github.com/jmoiron/sqlx"
)

type User struct {
	Id   int
	Name string
}

type Repo struct {
	Db *sqlx.DB
}

func RepoInterface(db *sqlx.DB) *Repo {
	return &Repo{Db: db}
}

func (repo *Repo) GetUser(id int) (*User, error) {

	var user User
	err := repo.Db.Get(&user, "SELECT id, name FROM user WHERE id = ? ", id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
