package main

import (
	"database/sql"
	"net/http"

	"github.com/Euvaz/go-log"
	"github.com/gin-gonic/gin"
)

func registerRoutes (router *gin.Engine, db *sql.DB) {
    router.LoadHTMLGlob("website/tmpl/*.tmpl")
    router.GET("/", func(ctx *gin.Context) {
        logger.Info("Handling GET /")
        data := gin.H {
            "title": "Virgil Lopez",
            "firstname": "Virgil",
            "lastname": "Lopez",
            "tagline": "Aspiring DevOps Engineer.",
        }
        ctx.HTML(http.StatusOK, "index.tmpl", data)
    })
}

