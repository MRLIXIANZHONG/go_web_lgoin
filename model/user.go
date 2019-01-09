package model

import (
	"github.com/jinzhu/gorm"
	"time"
	"web_test/pkg/e"
)

type User struct {
	Model
	Username string
	Mobile string
	Password string
	LastLogin time.Time
	LoginTimes int
}

var UserModel  = User{}//产生一个对象
// 或者userModel := new(User)


/*
查询是否有这条记录
 */
func (UserModel *User) UserFirst (id,password string)(err string, ret *User)  {

	db.First(UserModel,"id = ? AND password = ?",id,password)

	if UserModel.ID == 0 {
		return "没有权限",nil
	}
	return "",UserModel
}


/**
登录
 */
func (UserModel *User) UserLogin(mobile,password string)(err string, ret *User)  {
	//数据判断
	if len(password)<6 || !e.IsMobile(mobile){//验证密码和电话参数是否完整
		return "提交参数不完整", nil
	}
	db.First(&UserModel,"mobile = ?",mobile)
	if UserModel.ID == 0 {
		return  "账号不存在",nil
	}
	if e.MD5(password) != UserModel.Password{
		return  "密码不正确",nil
	}
	UserModel.LoginTimes++ //登录次数加一
	UserModel.LastLogin = time.Now() //登录时间改变
	db.Save(&UserModel)

	return "", UserModel
}

/**
创建会员
 */
func (UserModel *User) UserCreate(mobile,password string)(err string, ret *User)  {
	//数据判断
	if len(password)<6 || !e.IsMobile(mobile){//验证密码和电话参数是否完整
		return "提交参数不完整", nil
	}
	db.Where("mobile = ?",mobile).First(&UserModel)//查是否有这条记录
	//fmt.Printf("UserModel = %+v",UserModel)
	if UserModel.ID != 0 {
		return  "该电话号码已经被注册",nil
	}
	UserModel.Mobile = mobile
	UserModel.Password = e.MD5(password)
	db.Save(&UserModel)//创建会员
	return "", UserModel
}

//gorm创建时自动调用
func (UserModel *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now())//插入时间
	scope.SetColumn("UpdatedAt", time.Now())//更新时间默认为插入时间
	return nil
}
//gorm更改时自动调用
func (UserModel *User) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())

	return nil
}