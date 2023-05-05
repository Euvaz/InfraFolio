package main

import (
    "net/http"

    "github.com/Euvaz/go-log"
    "github.com/gin-gonic/gin"
)

func registerRoutes (router *gin.Engine) {
    router.LoadHTMLGlob("website/static/*.html")
    router.GET("/", func(ctx *gin.Context) {
        logger.Info("Handling GET /")
        ctx.HTML(http.StatusOK, "index.html", nil)
    })
}

