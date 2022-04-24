//User core entity, Author @by. Since 2022/4/22

package entitydefines

import (
	"gorm.io/gorm"
	"time"
)

// User information define
type User struct {
	id          int16     // primary key in database
	userId      string    //user id
	name        string    //user's name
	password    string    //user's password
	telephone   string    //user's telephone number, requires state number before it, example +86-12334556
	email       string    //user's email address
	createTime  time.Time //The create time of this account
	updateTime  time.Time //The latest update time of this account
	nationality string    // the nationality of this account's owner
	right       int8      // the right level of the user
}

//getters and setters

func (u *User) UserId() string {
	return u.userId
}
func (u *User) SetUserId(userId string) {
	u.userId = userId
}

func (u *User) Name() string {
	return u.name
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) Password() string {
	return u.password
}
func (u *User) SetPassword(password string) {
	u.password = password
}

func (u *User) Telephone() string {
	return u.telephone
}

func (u *User) SetTelephone(telephone string) {
	u.telephone = telephone
}

func (u *User) Email() string {
	return u.email
}

func (u *User) SetEmail(email string) {
	u.email = email
}

func (u *User) CreateTime() time.Time {
	return u.createTime
}

func (u *User) SetCreateTime(createTime time.Time) {
	u.createTime = createTime
}

func (u *User) UpdateTime() time.Time {
	return u.updateTime
}

func (u *User) SetUpdateTime(updateTime time.Time) {
	u.updateTime = updateTime
}

func (u *User) Nationality() string {
	return u.nationality
}
func (u *User) SetNationality(nation string) {
	u.nationality = nation
}

func (u *User) Right() int8 {
	return u.right
}

func (u *User) SetRight(right int8) {
	u.right = right
}

//database operations

// Save , save user info to db
func (u *User) Save(db *gorm.DB) error {
	if u.userId == "" {
		return nil
	} else {
		u.updateTime = time.Now()
		if err := db.Save(u).Error; err != nil {
			return err
		}
	}
	return nil
}
