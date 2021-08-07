package middleware

import (
	"github.com/gin-gonic/gin"
	"mbook/internal/model"
)

// Prepare方法中进行方法的预处理，这里不包含session 的处理，在session有专门的session中间件来处理
func Prepare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//	不提供匿名访问功能，这部分暂时不做处理

		//上方式session和cookie的实现部分，目前不做处理
		//	只做本地端调试，目前BaseUrl也不需要在意
		member := model.NewMember()
		ctx.Set("Member", member)
		ctx.Set("SITE_NAME", "MBOOK")
		ctx.Next()
	}
}
