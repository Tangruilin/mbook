package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"mbook/global"
	"mbook/internal/model"
	"net/http"
)

type HomeHandler struct{}

//Index is the HomePage
func (h *HomeHandler) Index(c *gin.Context) {
	//	暂时只是做一个测试
	cates, err := new(model.MdCategory).Get(global.DBEngine, -1, true)
	if err != nil {
		log.Fatal(err)
		return
	}
	c.Set("Cates", cates)
	c.HTML(http.StatusOK, "list.html", c.Keys)
}
