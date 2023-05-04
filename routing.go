package main

import (
    "database/sql"
    _ "net/http"

    "github.com/Euvaz/go-log"
    "github.com/gin-gonic/gin"
)

func registerRoutes (router *gin.Engine, db *sql.DB) {
    router.GET("/", func(ctx *gin.Context) {
        logger.Info("Handling GET /")
    })
}

