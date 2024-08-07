package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // Load HTML templates from the "templates" folder
    r.LoadHTMLGlob("templates/*")

    // Home Page
    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "layout.html", gin.H{
            "title": "Welcome to the Home Page!",
            "content": `<h1>Welcome to the Home Page!</h1><p>This is the home page of our web app.</p>`,
        })
    })

    // Projects Page
    r.GET("/projects", func(c *gin.Context) {
        c.HTML(http.StatusOK, "layout.html", gin.H{
            "title": "Projects",
            "content": `<h1>Projects</h1><p>This is a list of our exciting projects.</p>`,
        })
    })

    // Contact Page
    r.GET("/contact", func(c *gin.Context) {
        c.HTML(http.StatusOK, "layout.html", gin.H{
            "title": "Contact Us",
            "content": `<h1>Contact Us</h1><p>Contact us at <a href="mailto:info@example.com">info@example.com</a> or call 123-456-7890.</p>`,
        })
    })

    r.Run() // Listen and serve
}