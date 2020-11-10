package model

import (
	"ginblog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type: varchar(20);not null" json:"username"`
	Password string `gorm:"type: varchar(20);not null" json:"password"`
	Role     int    `gorm:"type: int" json:"role"`
}

//
func CheckUser(name string) (code int) {
	var users User
	//First 查询出第一个参数
	db.Select("id").Where("username = ?", name).First(&users)
	//fmt.Println(users.ID)
	if users.ID > 0 { //有记录
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS //证明可用
}

// 新增用户
func CreateUser(data *User) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS
}

// 查询用户列表
func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

// 编辑用户

// 删除用户
