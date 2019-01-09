package admin

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"web_test/model"
	"web_test/pkg/e"
)

type Result struct {
	Code int    `json:"code"` //返回状态码
	Msg  string `json:"msg"`  //返回提示信息
	Data *gin.H `json:"data"` //返回内容
}

type admin struct {
	//admin类
	User    *model.User
	Context *gin.Context
	Session sessions.Session
}

type Error struct {
	T   int		//跳转时间
	Url string	//跳转路径
	Msg string	//提示信息
}

var AdminError = new(Error)//产生一个错误提示信息对象

var AdminController = admin{} //产生一个对象

func ReturnJSON(c *gin.Context, result *Result) {
	if result.Code == 200 {
		result.Msg = "success"
	}
	c.JSON(200, result)

	c.Abort() //如果授权失败（例如密码不匹配），则调用abort以确保其余的处理程序不调用此请求。
}

/*
错误提示
 */
func (AdminError *Error) Errors(c *gin.Context) {
	//fmt.Printf("adminerror = %+v",AdminError)
	c.HTML(http.StatusOK, "layouts/err.html", gin.H{
		"msg":  AdminError.Msg,
		"url":  AdminError.Url,
		"time": AdminError.T,
	})
}

/**
验证session
 */
func (AdminController *admin) CheckSession(c *gin.Context) {
	session := sessions.Default(c)
	//fmt.Println("session_ =",session.Get("user_session"))
	user_session := session.Get("user_session")
	if user_session == nil {
		//ReturnJSON(c, &Result{-1, "没有权限", nil})
		AdminError.T = 2
		AdminError.Url = "/"
		AdminError.Msg = "没有权限"
		c.Redirect(http.StatusFound, "/err")//重定向
		return
		//c.Redirect(301, "/err")永久重定向，坑爹
	}

	s := strings.Split(user_session.(string), "|")
	id := s[0]
	password := s[1]

	if id == "" || !e.IsNumeric(id) { //判断id是否为空，或者是否为数字
		/*ReturnJSON(c, &Result{-1, "没有权限", nil})
		return*/
		AdminError.T = 2
		AdminError.Url = "/"
		AdminError.Msg = "没有权限"
		c.Redirect(http.StatusFound, "/err")//重定向
		return
	}

	err, user := model.UserModel.UserFirst(id, password)

	if err != "" {
		/*ReturnJSON(c, &Result{-1, err, nil})
		return*/
		AdminError.T = 2
		AdminError.Url = "/"
		AdminError.Msg = err
		c.Redirect(http.StatusFound, "/err")//重定向
		return
	}
	AdminController.User = user
	c.Next() //执行其他的服务
}
