//General ms dao service operations, Author @by. since 2022/4/23

package dao

import (
	"db/entitydefines"
	"github.com/go-playground/validator"
	"time"
)

var db = InitDB()

//add user

func AddUser(u *entitydefines.User) error {
	if u.UserId() == "" {
		u.SetUserId(GenerateUUID())
	}
	u.SetCreateTime(time.Now())
	err := u.Save(db)
	if err != nil {
		return err
	}
	return nil
}

//modify user

func ModifyUser(u *entitydefines.User) error {
	var validate = validator.New()
	errs := validate.Struct(u)
	if errs != nil {
		return errs
	}
	u.Save(db)
	return nil
}
