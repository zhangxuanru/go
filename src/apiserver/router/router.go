package router

import (
	"github.com/gin-gonic/gin"
	"apiserver/router/middleware"
	"net/http"
	"apiserver/handler/sd"
)

func Load(g *gin.Engine,mw ...gin.HandlerFunc) *gin.Engine  {
     g.Use(gin.Recovery())
     g.Use(middleware.NoCache)
     g.Use(middleware.Options)
     g.Use(middleware.Secure)
     g.Use(mw...)

     g.NoRoute(func(context *gin.Context) {
		    context.String(http.StatusNotFound,"The incorrect API route.")
	 })

     svcd := g.Group("/sd")
     {
           svcd.GET("/health",sd.HealthCheck)
	 }

	 return  g
}