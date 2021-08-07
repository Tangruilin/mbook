package api

import "github.com/gin-gonic/gin"

type DocumentHandler struct{}

// Index /*书籍搜索方法*/
func (d *DocumentHandler) Index(c *gin.Context) {}

/*读书模块*/
func (d *DocumentHandler) Read(c *gin.Context) {}

func (d *DocumentHandler) Search(c *gin.Context) {}

/*图书编辑模块*/

func (d *DocumentHandler) Edit(c *gin.Context) {}

func (d *DocumentHandler) Content(c *gin.Context) {}

func (d *DocumentHandler) Upload(c *gin.Context) {}

func (d *DocumentHandler) Create(c *gin.Context) {}

func (d *DocumentHandler) Delete(c *gin.Context) {}
