package main

import (
    "database/sql"
    "net/http"
    "sort"

    "github.com/Euvaz/go-log"
    "github.com/gin-contrib/multitemplate"
    "github.com/gin-gonic/gin"
)

func createMyRender() multitemplate.Renderer {
  renderer := multitemplate.NewRenderer()
  renderer.AddFromFiles("index", "website/tmpl/layout.tmpl", "website/tmpl/index.tmpl")
  renderer.AddFromFiles("projects", "website/tmpl/layout.tmpl", "website/tmpl/projects.tmpl")
  renderer.AddFromFiles("contact", "website/tmpl/layout.tmpl", "website/tmpl/contact.tmpl")
  return renderer
}

func registerRoutes (router *gin.Engine, db *sql.DB) {

    router.HTMLRender = createMyRender()

    EMAIL := "6f4baa76-cf90-4803-bc66-19838566253e@simplelogin.com"
    GITHUB_USER := "Euvaz"
    LINKEDIN_USER := "VirgilLopez"

    router.GET("/", func(ctx *gin.Context) {
        logger.Info("Handling GET /")
        
        data := gin.H {
            "title": "Virgil Lopez",
            "tagline": "Aspiring DevOps Engineer.",
            "firstname": "Virgil",
            "lastname": "Lopez",
            "linkedin_user": LINKEDIN_USER,
        }
        ctx.HTML(http.StatusOK, "index", data)
    })

    router.GET("/projects", func(ctx *gin.Context) {
        logger.Info("Handling GET /projects")

        type GitHubRepository struct {
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

        github_repos := []GitHubRepository{}
        for rows.Next() {
            var repo GitHubRepository
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
            "tagline": "Projects",
            "firstname": "Virgil",
            "lastname": "Lopez",
            "github_repos": github_repos,
        }
        ctx.HTML(http.StatusOK, "projects", data)
    })

    router.GET("/contact", func(ctx *gin.Context) {
        logger.Info("Handling GET /contact")
        data := gin.H {
            "title": "Virgil Lopez",
            "tagline": "Contact Me",
            "firstname": "Virgil",
            "lastname": "Lopez",
            "email": EMAIL,
            "github_user": GITHUB_USER,
            "linkedin_user": LINKEDIN_USER,
        }
        ctx.HTML(http.StatusOK, "contact", data)
    })
}

