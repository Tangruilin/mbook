package api

import "github.com/gin-gonic/gin"

type BaseHandler struct{}

type BaseData struct {
}

func (b *BaseHandler) SetFollow(c *gin.Context) {}
