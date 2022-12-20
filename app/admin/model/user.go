package model

import (
	"go-mage-admin/app/core"
	"go-mage-admin/app/core/model"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type User struct {
	//匿名继承
	model.Collection
	UserId     int    `gorm:"primary_key" json:"user_id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	DateCreate int    `json:"date_create"`
	IsActive   int    `json:"is_active"`
}

/*
	gorm.DefaultTableNameHandler = func(db *gorm.DB,defaultTableName string){
		// 全局表前缀
	    return "sys_" + defaultTableName
	}
*/
func (*User) TableName() string {
	return "admin_user"
}

// Authenticate 用户名和密码验证
func (user *User) Authenticate(username string, password string) (bool, string, *User) {
	u := new(User)
	core.AppDb["read"].First(&u, "username = ?", username)
	msg := ""
	flag := false
	if u.UserId != 0 {
		if u.IsActive == 1 {
			err := bcrypt.CompareHashAndPassword([]byte(password), []byte(u.Password))
			if err != nil {
				flag = true
			} else {
				msg = "登录密码不正确！"
			}
		} else {
			msg = "当前帐号已禁用！"
		}
	} else {
		msg = "用户名不存在！"
	}
	return flag, msg, u
}
func (user *User) GetCollection(filters map[string]interface{}, orders [2]string, page int, size int) ([]User, int64) {
	var users []User
	var total int64
	db := core.AppDb["read"].Model(User{})

	//多条件过滤
	collection := model.Collection{}
	db = collection.PrepareCollection(db, filters)
	log.Println("filters:", filters)
	db2 := db
	db.Debug().Count(&total)
	res := collection.LoadCollection(db2, orders, page, size).Find(&users)
	if res.Error != nil {
		log.Println(res.Error)
		panic("获取User失败")
	}
	return users, total
}

// GetGridColumnsConfig 获取用户自定义的Grid
func (user *User) GetGridColumnsConfig(uid int, gridCode string) []string {
	configGrid := new(ConfigGrid).getConfig(uid, gridCode)
	return configGrid
}
