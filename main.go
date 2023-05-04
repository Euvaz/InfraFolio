package main

import (
    "database/sql"

    "github.com/Euvaz/go-log"
    "github.com/gin-gonic/gin"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    db := getDB()

    logger.Info("Starting server")
    initDB(db)
    defer closeDB(db)

    router := gin.Default()
    registerRoutes(router, db)

    router.Run()
}

func initDB(db *sql.DB) {
	var err error

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS github (
                        id SERIAL PRIMARY KEY,
                        name TEXT,
                        contributions INTEGER
                    );`)
	if err != nil {
		logger.Fatal(err.Error())
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS github_repositories (
                        id SERIAL PRIMARY KEY,
                        owner TEXT,
                        name TEXT
                    );`)
	if err != nil {
		logger.Fatal(err.Error())
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS github_stars (
                        id SERIAL PRIMARY KEY,
                        name TEXT,
                        github_repositories_id SERIAL
                        REFERENCES github_repositories (id)
                    );`)
	if err != nil {
		logger.Fatal(err.Error())
	}

	logger.Info("Tables successfully initialized")
}

func getDB () *sql.DB {
    database, err := sql.Open("sqlite3", "file:locked.sqlite?cache=shared")
    if err != nil {
        logger.Fatal(err.Error())
    }
    database.SetMaxOpenConns(1)
    return database
}

func closeDB(db *sql.DB) {
	err := db.Close()
	if err != nil {
		logger.Fatal(err.Error())
	}
	logger.Info("Database connection closed")
}

