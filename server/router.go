package server

import (
	"singo/api"
	"singo/serializer/handler"
	
	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()
	
	v1 := handler.NewHandler("/api/v1")
	defer v1.Mount(r)
	
	// 路由
	{
		v1.POST("ping", api.Ping)
	}
	return r
}


