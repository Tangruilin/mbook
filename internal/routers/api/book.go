package api

import "github.com/gin-gonic/gin"

type BookHandler struct{}

//图书管理API

func (b *BookHandler) Index(c *gin.Context) {}

func (b *BookHandler) Create(c *gin.Context) {}

func (b *BookHandler) Setting(c *gin.Context) {}

func (b *BookHandler) UploadCover(c *gin.Context) {}

func (b *BookHandler) Collection(c *gin.Context) {}

func (b *BookHandler) SaveBook(c *gin.Context) {}

func (b *BookHandler) Release(c *gin.Context) {}

func (b *BookHandler) CreatToken(c *gin.Context) {}

/*个人中心内容*/

func (b *BookHandler) Score(c *gin.Context) {}

func (b *BookHandler) Comment(c *gin.Context) {}
