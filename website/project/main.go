package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.LoadHTMLGlob("templates/*") // Load HTML templates from the "templates" folder

    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "home.html", gin.H{
            "title": "Welcome to the Home Page!",
        })
    })

    r.GET("/projects", func(c *gin.Context) {
        c.HTML(http.StatusOK, "projects.html", gin.H{
            "title": "Projects",
            "info": "This is a list of our exciting projects.",
        })
    })

    r.GET("/contact", func(c *gin.Context) {
        c.HTML(http.StatusOK, "contact.html", gin.H{
            "title": "Contact Us",
            "email": "info@example.com",
            "phone": "123-456-7890",
        })
    })

    r.Run() // Listen and serve
}