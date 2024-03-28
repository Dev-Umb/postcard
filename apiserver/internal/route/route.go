package route

import (
	"github.com/gin-gonic/gin"
	"postcard/apiserver/internal/store"
	"postcard/apiserver/pkg/core"
)

func ApiServer(g *gin.Engine, sro store.DataStore) *gin.Engine {
	g.Use(gin.Logger())
	g.Use(gin.Recovery())
	g.GET("/init", func(context *gin.Context) {
		core.SendResponse(context, nil, struct {
			Data string `json:"data"`
		}{
			Data: "ok",
		})
	})
	return g
}
