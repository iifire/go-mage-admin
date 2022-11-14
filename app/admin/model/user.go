package model

type User struct {
	UserId     int
	Username   string
	Password   string
	Email      string
	DateCreate string
	IsActive   int
}

/*
	gorm.DefaultTableNameHandler = func(db *gorm.DB,defaultTableName string){
		// 全局表前缀
	    return "sys_" + defaultTableName
	}
*/
func (User) TableName() string {
	return "admin_user"
}
