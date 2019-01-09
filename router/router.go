package routers

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"web_test/controller/admin"
	"web_test/pkg/poptime"
	"web_test/pkg/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New() // 创建一个默认的没有任何中间件的路由

	r.Use(gin.Logger())

	r.Use(gin.Recovery())
	//r.Use(gzip.Gzip(gzip.DefaultCompression))
	// store, _ := sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	store := sessions.NewCookieStore([]byte("session.secret.password"))
	r.Use(sessions.Sessions("session", store))//开启加载session

	gin.SetMode(setting.RunMode)//根据输入字符串设置gin模式，api里面的dubug
	//r.Static("/satatic","./satatic")
	r.SetFuncMap(template.FuncMap{//自定义模板函数，传入模板
		"sum": Sum,//加
		"reduce":Reduce,//减
		"formDate" :poptime.FormatUnixTimeAsStr,//时间转换函数
	})
	r.StaticFS("/static/",http.Dir("static"))//处理静态文件
	r.GET("/favicon.ico", func(c *gin.Context) {
		c.Header("Cache-Contorl", "public, max-age=2592000")//session存储时间2592000一个月时间
		c.File("./static/favicon.ico")
	})

	r.LoadHTMLGlob("view/**/*")//设置模板路径

	r.GET("/", func(c *gin.Context) {//显示登录界面视图
		c.HTML(http.StatusOK,"login/login.html",nil)
	})
	r.POST("/login",admin.AdminController.Login)

	//r.GET("/test",admin.AdminController.CheckSession)//测试
	//重定向错误页面
	r.GET("/err",admin.AdminError.Errors)

	admin1:=r.Group("/admin")
	admin1.Use(admin.AdminController.CheckSession)//可以分配验证
	{
		admin1.GET("/",func(c *gin.Context) {//后台首页
			c.HTML(http.StatusOK,"index/index.html",nil)
		})
		admin1.GET("/createuserv", func(c *gin.Context) {//添加管理员视图
			c.HTML(http.StatusOK,"user/create.html",nil)
		})
		admin1.POST("/createuser",admin.AdminController.CreateUser)//添加会员
		//退出登录
		admin1.GET("/logout",admin.AdminController.Logout)
		//统计所有记录
		admin1.GET("/count",admin.AdminController.Counts)
		//更改密码
		admin1.POST("/editpassword",admin.AdminController.EditPassword)
		//查看账户信息
		admin1.GET("/account",admin.AdminController.AccountIndex)
		//查看内容
		admin1.GET("/content",admin.AdminController.ContentIndex)

	}

	return r
}

//自定义模板函数
//加
func Sum(a,b int)(c int)  {
	c = a + b
	return 
}
//减
func Reduce(a,b int)(c int)  {
	if a ==0{
		a = 1 //不然a为零
	}
	c = a - b
	return
}

