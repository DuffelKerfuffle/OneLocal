package main

import (
	"fmt"
	"net/http"
	"net/smtp"
	"os"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type area struct {
	Name             string
	BusinessesInArea []business
}

type business struct {
	Name        string
	Description string
	ContactInfo string
	Address     string
	ImageLink   string
}

func main() {
	email := "onelocal4@gmail.com"
	password := "h@llo123"
	receiver := []string{"onelocal4@gmail.com"}

	smtpHost := "smtp.gmail.com"
	smtpport := "587"
	r := gin.Default()
	areas := LoadAll()
	for _, v := range areas {
		fmt.Println(v.Name)
		for _, i := range v.BusinessesInArea {
			fmt.Println(i.Name)
		}
	}
	r.Use(static.Serve("/images", static.LocalFile("./images", true)))
	r.Use(static.Serve("/.fonts", static.LocalFile("./.fonts", true)))
	r.LoadHTMLFiles("BusinessWeb.html", "HomePage.html", "Contact.html", "Terms.html", "businesses.html")

	r.GET("/home", func(c *gin.Context) {
		areas = LoadAll()
		c.HTML(http.StatusOK, "HomePage.html", areas)
	})

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/home")
	})

	r.GET("/terms", func(c *gin.Context) {
		c.HTML(http.StatusOK, "Terms.html", areas)
	})

	r.GET("/contact", func(c *gin.Context) {
		c.HTML(http.StatusOK, "Contact.html", areas)
	})

	r.POST("/contact", func(c *gin.Context) {
		name := c.PostForm("name")
		contacts := c.PostForm("contact")
		details := c.PostForm("message")

		message := []byte(name + "\n" + contacts + "\n" + details)
		auth := smtp.PlainAuth("", email, password, smtpHost)

		err := smtp.SendMail(smtpHost+":"+smtpport, auth, email, receiver, message)
		if err != nil {
			fmt.Println(err)
			return
		}
		c.HTML(http.StatusOK, "Contact.html", areas)
	})

	// for _, v := range areas {
	// 	for _, j := range v.BusinessesInArea {
	// 		r.GET("/business/"+j.Name, func(c *gin.Context) {
	// 			areas = LoadAll()
	// 			c.HTML(http.StatusOK, "BusinessWeb.html", j)
	// 		})
	// 	}
	// }

	r.GET("/business", func(c *gin.Context) {
		areas = LoadAll()

		place := c.Request.URL.Query().Get("business")
		for _, v := range areas {
			for _, u := range v.BusinessesInArea {
				if place == u.Name {
					c.HTML(http.StatusOK, "BusinessWeb.html", u)
					break
				}
			}
		}
	})

	r.GET("/businesses", func(c *gin.Context) {
		areas = LoadAll()
		c.HTML(http.StatusOK, "businesses.html", areas)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000" // Default port if not specified
	}
	err := r.Run(":" + port)
	if err != nil {
		panic(err)
	}
}
