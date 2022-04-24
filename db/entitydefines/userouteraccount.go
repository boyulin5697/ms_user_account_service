//UserOuterAccount definition, Author @by. since 2022/4/22
//Describe the outside system accounts and information.

package entitydefines

import (
	"gorm.io/gorm"
	"time"
)

type UserOuterAccount struct {
	id                       int16     //primary key
	userOuterAccountId       string    //user outer accountId
	userId                   string    //the owner's userId
	userOuterAccountIdentify string    //user's outer account's identification, for example, his Google account.
	platformId               string    //The account's belonging platform, there could be an enum here.
	createdTime              time.Time // the created time of this outer account.
	updatedTime              time.Time // the updated time
	isUsing                  bool      // Is this account still usable, If not, please set it as false.
}

//getters and setters

func (u *UserOuterAccount) UserOuterAccountId() string {
	return u.userOuterAccountId
}
func (u *UserOuterAccount) SetUserOuterAccountId(accountId string) {
	u.userOuterAccountId = accountId
}
func (u *UserOuterAccount) UserId() string {
	return u.userId
}
func (u *UserOuterAccount) SetUserId(userId string) {
	u.userId = userId
}
func (u *UserOuterAccount) UserOuterAccountIdentify() string {
	return u.userOuterAccountIdentify
}
func (u *UserOuterAccount) SetUserOuterAccountIdentify(identify string) {
	u.userOuterAccountIdentify = identify
}
func (u *UserOuterAccount) PlatformId() string {
	return u.platformId
}
func (u *UserOuterAccount) SetPlatformId(platformId string) {
	u.platformId = platformId
}
func (u *UserOuterAccount) CreatedTime() time.Time {
	return u.createdTime
}

func (u *UserOuterAccount) SetCreatedTime(createdTime time.Time) {
	u.createdTime = createdTime
}

func (u *UserOuterAccount) UpdatedTime() time.Time {
	return u.updatedTime
}

func (u *UserOuterAccount) SetUpdatedTime(updatedTime time.Time) {
	u.updatedTime = updatedTime
}

func (u *UserOuterAccount) IsUsing() bool {
	return u.isUsing
}

func (u *UserOuterAccount) SetIsUsing(isUsing bool) {
	u.isUsing = isUsing
}

//database actions

func (u *UserOuterAccount) Save(db *gorm.DB) error {
	if u.userOuterAccountId == "" {
		return nil
	}
	u.updatedTime = time.Now()
	if err := db.Save(u).Error; err != nil {
		return err
	}
	return nil
}
