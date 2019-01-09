package admin

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"web_test/model"
	"web_test/pkg/e"
)

/**
登录
 */
func (AdminController *admin) Login(c *gin.Context) {
	//开启session
	//	c.Header("Cache-Control", "no-cache, no-store, max-age=0")//设置header头
	session := sessions.Default(c) //产生一个session
	//接受参数
	mobile := c.PostForm("mobile")
	password := c.PostForm("password")
	//调用model判断参数
	err, user := model.UserModel.UserLogin(mobile, password)
	if err != "" {
		ReturnJSON(c, &Result{0, err, nil}) //返回前端错误信息
		return
	}
	// 保存登录信息
	session.Set("user_session", strconv.Itoa(int(user.ID))+"|"+user.Password)
	_ = session.Save()
	fmt.Println("user_session = ", session.Get("user_session"))
	// 给前端返回状态
	ReturnJSON(c, &Result{200, err, nil})
}

/**
创建用户
 */
func (AdminController *admin) CreateUser(c *gin.Context) {
	//接受参数
	mobile := c.PostForm("mobile")
	password := c.PostForm("password")
	err, _ := model.UserModel.UserCreate(mobile, password)
	if err != "" {
		ReturnJSON(c, &Result{0, err, nil}) //返回前端错误信息
		return
	}
	// 给前端返回状态
	ReturnJSON(c, &Result{200, err, nil})
}

/*
修改密码
 */
func (AdminController *admin) EditPassword(c *gin.Context)  {
	//接受前端的参数
	password := c.PostForm("password")
	if len(password) < 6 || len(password) > 32 {
		ReturnJSON(c, &Result{0, "密码格式不正确", nil})
		return
	}
	//查询session
	session :=sessions.Default(c)
	user_session := session.Get("user_session")
	s := strings.Split(user_session.(string), "|")//分成两个切片
	id := s[0]
	db := model.DbInstance()//调用db
	db.First(&AdminController.User,"id = ?",id)
	AdminController.User.Password = e.MD5(password)
	err:=db.Save(&AdminController.User).Error
	if err !=nil {
		ReturnJSON(c, &Result{0, "密码数据保存失败", nil})
		return
	}
	// 保存更改后的信息
	session.Set("user_session", strconv.Itoa(int(AdminController.User.ID))+"|"+AdminController.User.Password)
	_ = session.Save()
	fmt.Println("user_session = ", session.Get("user_session"))

	ReturnJSON(c, &Result{200, "", nil})
}

/*
退出登录
 */
func (AdminController *admin)Logout(c *gin.Context)  {
	session := sessions.Default(c)
	session.Delete("user_session")
	_=session.Save()
	AdminError.T = 2
	AdminError.Url = "/"
	AdminError.Msg = "退出成功"
	c.Redirect(http.StatusFound, "/err")//重定向
	return

}