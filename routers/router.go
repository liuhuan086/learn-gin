package routers

import (
	"example/pkg/settings"
	v1 "example/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(settings.RunMode)

	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/tags", v1.GetTags)

		apiV1.POST("/tags", v1.AddTag)

		apiV1.PUT("/tags/:id", v1.EditTag)

		apiV1.DELETE("/tags/:id", v1.DeleteTag)
	}
	return r
}
