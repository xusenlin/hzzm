package routes

import (
	"github.com/gin-gonic/gin"
	"hzzm/config"
	"hzzm/controller"
	"hzzm/middlewares"
	"net/http"
)

func InitRouter() *gin.Engine {
	// 设置gin模式
	gin.SetMode(config.GinMode)

	r := gin.Default()
	r.Use(middlewares.Cors())

	r.GET("/ping", func(c *gin.Context) { //服务健康检查
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})


	v1 := r.Group("/v1")
	{
		v1.GET("/ts",controller.TangShiPaginate)
		v1.GET("/ts/favour",controller.TangShiFavour)

		v1.GET("/ss",controller.SongShiPaginate)
		v1.GET("/ss/favour",controller.SongShiFavour)

		v1.GET("/sc",controller.SongCiPaginate)
		v1.GET("/sc/favour",controller.SongCiFavour)

		v1.GET("/yq",controller.YuanQuPaginate)
		v1.GET("/yq/favour",controller.YuanQuFavour)

		v1.GET("/ly",controller.LunYuPaginate)
		v1.GET("/ly/favour",controller.LunYuFavour)

		v1.GET("/sswj",controller.SiShuWuJingPaginate)
		v1.GET("/sswj/favour",controller.SiShuWuJingFavour)

		v1.GET("/ymy",controller.YouMengYingPaginate)
		v1.GET("/ymy/favour",controller.YouMengYingFavour)
	}

	return r
}
