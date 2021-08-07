package api

import "github.com/gin-gonic/gin"

type UserHandler struct{}

// user 文件表示的是个人中心的api

func (u *UserHandler) Index(c *gin.Context) {}

func (u *UserHandler) Collection(c *gin.Context) {}

func (u *UserHandler) Follow(c *gin.Context) {}

func (u *UserHandler) Fans(c *gin.Context) {}
