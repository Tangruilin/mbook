//Package routers is the router, the API with Restful
package routers

import (
	"github.com/astaxie/beego"
	"github.com/gin-gonic/gin"
	"html/template"
	"mbook/internal/routers/api"
	"mbook/middleware"
	"mbook/utils"
	"strings"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Prepare())
	r.FuncMap = template.FuncMap{
		"showImg": utils.ShowImg,
		"cdnjs": func(p string) string {
			cdn := beego.AppConfig.DefaultString("cdnjs", "")
			if strings.HasPrefix(p, "/") && strings.HasSuffix(cdn, "/") {
				return cdn + string(p[1:])
			}
			if !strings.HasPrefix(p, "/") && !strings.HasSuffix(cdn, "/") {
				return cdn + "/" + p
			}
			return cdn + p
		},
		"cdncss": func(p string) string {
			cdn := beego.AppConfig.DefaultString("cdcss", "")
			if strings.HasPrefix(p, "/") && strings.HasSuffix(cdn, "/") {
				return cdn + string(p[1:])
			}
			if !strings.HasPrefix(p, "/") && !strings.HasSuffix(cdn, "/") {
				return cdn + p
			}
			return cdn + p
		},
		"getUsernameByUid": func(id interface{}) string {
			return ""
		},
		"getNicknameByUid": func(id interface{}) string {
			return ""
		},
		"inMap": utils.InMap,
		"doesCollection": func(uid, bid interface{}) bool {
			return false
		},
		"IsFollow": func(mid, fansId interface{}) bool {
			return false
		},
		"isubstr": utils.Substr,
	}
	r.Static("/static", "./static")
	r.LoadHTMLGlob("views/**/*")

	var (
		base     = new(api.BaseHandler)
		book     = new(api.BookHandler)
		document = new(api.DocumentHandler)
		explore  = new(api.ExploreHandler)
		home     = new(api.HomeHandler)
		account  = new(api.AccountHandler)
		manager  = new(api.ManageHandler)
		search   = new(api.SearchHandler)
		setting  = new(api.SettingHandler)
		user     = new(api.UserHandler)
	)
	//	中间是路由的绑定部分
	//绑定跟目录
	r.GET("/", home.Index)
	r.Any("/explore", explore.Index)
	r.Any("/books/:key", document.Index)
	//登录
	{
		r.Any("/login", account.Login)
		r.Any("/regist", account.Regist)
		r.Any("/logout", account.Logout)
		r.Any("doregist", account.DoRegist)
	}
	//编辑
	documentApi := r.Group("/api")
	{
		documentApi.Any("/:key/edit/?:id", document.Edit)
		documentApi.Any("/:key/content/?:id", document.Content)
		documentApi.POST("/upload", document.Upload)
		documentApi.POST("/:key/create", document.Create)
		documentApi.POST("/:key/delete", document.Delete)
	}
	//读书
	readApi := r.Group("/read")
	{
		readApi.Any("/:key/:id", document.Read)
		readApi.POST("/:key/search", document.Search)
	}
	// 搜索
	{
		r.GET("/search", search.Search)
		r.GET("/search/result", search.Result)
	}
	// 图书管理 部分个人中心
	bookApi := r.Group("/book")
	{
		bookApi.Any("", book.Index)
		bookApi.POST("/create", book.Create)
		bookApi.Any("/:key/setting", book.Setting)
		bookApi.POST("/setting/upload", book.UploadCover)
		bookApi.Any("/star/:id", book.Collection)
		bookApi.POST("/setting/save", book.SaveBook)
		bookApi.POST("/:key/release", book.Release)
		bookApi.POST("/setting/token", book.CreatToken)
		bookApi.Any("/score/:id", book.Score)
		bookApi.POST("/comment/:id", book.Comment)
	}
	//个人中心
	r.GET("/follow/:uid", base.SetFollow)
	userApi := r.Group("/user")
	{
		userApi.GET("/:username", user.Index)
		userApi.GET("/:username/collection", user.Collection)
		userApi.GET("/:username/follow", user.Follow)
		userApi.GET("/:username/fans", user.Fans)
	}
	settingApi := r.Group("/setting")
	{
		settingApi.Any("", setting.Index)
		settingApi.Any("/upload", setting.Upload)
	}
	managerApi := r.Group("/manager")
	{
		managerApi.POST("/category", manager.Category)
		managerApi.GET("/update-cate", manager.UpdateCate)
		managerApi.GET("/del-cate", manager.DelCate)
		managerApi.POST("/icon-cate", manager.UpdateCateIcon)
	}
	return r
}
