package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web_test/model"
)

/*
显示内容
 */

func (AcminController *admin) ContentIndex(c *gin.Context)  {
	//调用model查询数据
	err,content := model.ContentModel.ContentIndex()
	if err != "" {
		ReturnJSON(c, &Result{0, err, nil})
		return
	}
	//fmt.Printf("content = %+v",content)
	c.HTML(http.StatusOK,"contents/contents.html",content)

}





//统计所有数据,供后台页面展示
func (AdminController *admin) Counts(c *gin.Context) {
	type ResultData struct {
		Accounts, Contents, Reviews int//账户，内容，评论
	}
	var data ResultData
	db := model.DbInstance()//调用db
	db.Model(&model.Account{}).Count(&data.Accounts)
	db.Model(&model.Content{}).Count(&data.Contents)
	db.Model(&model.Content{}).Where("review > ?", 0).Count(&data.Reviews)
	c.JSON(200, data)
}