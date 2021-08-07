package api

import "github.com/gin-gonic/gin"

type SearchHandler struct{}

/*搜索模块的接口*/

func (s *SearchHandler) Search(c *gin.Context) {}

func (s *SearchHandler) Result(c *gin.Context) {}
