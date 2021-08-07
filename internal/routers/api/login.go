package api

import "github.com/gin-gonic/gin"

type AccountHandler struct{}

/*登录方法注册*/

func (h *AccountHandler) Login(c *gin.Context) {}

func (h *AccountHandler) Regist(c *gin.Context) {}

func (h *AccountHandler) Logout(c *gin.Context) {}

func (h *AccountHandler) DoRegist(c *gin.Context) {}
