package api

import "github.com/gin-gonic/gin"

type ManageHandler struct{}

func (m *ManageHandler) Category(c *gin.Context) {}

func (m *ManageHandler) UpdateCate(c *gin.Context) {}

func (m *ManageHandler) DelCate(c *gin.Context) {}

func (m *ManageHandler) UpdateCateIcon(c *gin.Context) {}
