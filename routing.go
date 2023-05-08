package main

import (
	"database/sql"
	"net/http"
    "sort"

	"github.com/Euvaz/go-log"
	"github.com/gin-gonic/gin"
)

func registerRoutes (router *gin.Engine, db *sql.DB) {

    router.LoadHTMLGlob("website/tmpl/*.tmpl")
    router.GET("/", func(ctx *gin.Context) {
        logger.Info("Handling GET /")
        //GITHUB_USER := "Euvaz"
        //LINKEDIN_USER := "VirgilLopez"
        type GitHubRepositories struct {
            ID                  int
            Owner               string
            Name                string
            Preview             string
            Stargazers_count    int
        }

        rows, err := db.Query("SELECT id, owner, name, preview, stargazers_count FROM github_repositories") 
        if err != nil {
            logger.Fatal(err.Error())
        }
        defer rows.Close()

        github_repos := []GitHubRepositories{}
        for rows.Next() {
            var repo GitHubRepositories
            err = rows.Scan(&repo.ID, &repo.Owner, &repo.Name, &repo.Preview, &repo.Stargazers_count)
            if err != nil {
                logger.Fatal(err.Error())
            }
            github_repos = append(github_repos, repo)
        }

        if err = rows.Err(); err != nil {
            logger.Fatal(err.Error())
        }

        sort.Slice(github_repos, func(i, j int) bool {
            return github_repos[i].Stargazers_count > github_repos[j].Stargazers_count
        })
        
        data := gin.H {
            "title": "Virgil Lopez",
            "firstname": "Virgil",
            "lastname": "Lopez",
            "tagline": "Aspiring DevOps Engineer.",
            "github_repos": github_repos,
        }
        ctx.HTML(http.StatusOK, "index.tmpl", data)
    })
}

