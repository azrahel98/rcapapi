package mysql

import (
	"database/sql"
	"errors"

	"villas.com/src/modelos"
	"villas.com/src/service"
)

type UserImpl struct{}

func (u UserImpl) Login(username, pass string) (*modelos.User, error) {
	db, _ := service.GetInstance()
	var user modelos.User
	err := db.Query("select userid ,nickname ,password,Area ,admin  from `User` where nickname = ?", func(r *sql.Rows) error {
		for r.Next() {
			err := r.Scan(&user.Id, &user.Username, &user.Password, &user.Areas, &user.IsAmin)
			if err != nil {
				return err
			}
		}
		return nil
	}, username)
	if user.Id == 0 {
		return nil, errors.New("no existe")
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}
