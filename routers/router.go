package routers

import (
	"example/middlerware/jwt"
	"example/pkg/settings"
	"example/routers/api"
	v1 "example/routers/api/v1"
	_ "github.com/EDDYCJY/go-gin-example/docs"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(settings.RunMode)

	r.GET("/auth", api.GetAuth)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiV1 := r.Group("/api/v1")

	apiV1.Use(jwt.JWT())

	{
		apiV1.GET("/tags", v1.GetTags)

		apiV1.GET("/tags/:id", v1.GetATag)

		apiV1.POST("/tags", v1.AddTag)

		apiV1.PUT("/tags/:id", v1.EditTag)

		apiV1.DELETE("/tags/:id", v1.DeleteTag)

		apiV1.GET("/articles", v1.GetArticles)

		apiV1.GET("/articles/:id", v1.GetAnArticle)

		apiV1.POST("/articles", v1.AddAnArticle)

		apiV1.PUT("/articles/:id", v1.EditAnArticle)

		apiV1.DELETE("/articles/:id", v1.DeleteAnArticle)

	}
	return r
}
