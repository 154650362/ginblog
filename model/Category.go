package model

import (
	"ginblog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type: varchar(20);not null" json:"name"`
}

//查询分类是否存在
func CheckCategory(name string) (code int) {
	var cate Category
	//First 查询出第一个参数
	db.Select("id").Where("name = ?", name).First(&cate)
	//fmt.Println(users.ID)
	if cate.ID > 0 { //有记录
		return errmsg.ERROR_CATENAME_USED
	}
	return errmsg.SUCCESS //证明可用
}

// 新增分类
func CreateCate(data *Category) int {
	//data.Password = ScryprPw(data.Password) //加密密码
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS
}

//todo 查询分类下的文章

// 查询分类列表
func GetCate(pageSize int, pageNum int) []Category {
	var cate []Category
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cate).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return cate
}

// 编辑分类信息
func EditCate(id int, data *Category) int {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name

	err = db.Model(&cate).Where("id = ? ", id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除
func DeleteCate(id int) int {
	var cate Category
	err = db.Where("id = ? ", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
