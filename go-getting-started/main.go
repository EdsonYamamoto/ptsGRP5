package main

import (
	"os"

	"github.com/robfig/cron"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ctr "github.com/heroku/go-getting-started/src/controller"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	//var err error
	port := os.Getenv("PORT")

	if port == "" {
		//log.Fatal("$PORT must be set")
		port = "8080"
	}

	c := cron.New()

	c.Start()
	c.AddFunc("0 0 17 * * *", ctr.CronFazAlgo) //executa uma ação as 17h UTC

	router := gin.Default()
	router.Use(cors.Default())
	router.LoadHTMLGlob("templates/*.tmpl.html")

	router.Static("/static", "static")

	router.GET("/", ctr.GETIndex)
	router.GET("/documentation", ctr.GETDocumentation)
	router.GET("/teste", ctr.GETTeste)
	router.GET("/collection/:name", ctr.GETAllCollection)
	router.POST("/collection/POSTSaveCollection", ctr.POSTSaveCollection)
	router.POST("/collection/POSTgetAllNotDeletedFromCollection", ctr.POSTGetAllDocumentsNotDeletedFromCollection)
	router.POST("/collection/POSTgetAllSearchFromCollection", ctr.POSTGetAllSearchFromCollection)
	router.POST("/collection/POSTgetOneDocumentFromCollection", ctr.POSTGetOneSearchFromCollection)
	router.POST("/collection/POSTdeleteDocumentCollection", ctr.POSTdeleteDocumentCollection)
	router.POST("/collection/POSTupdateDocumentCollection", ctr.POSTupdateDocumentCollection)

	//WEBHOOK
	router.GET("/facebook/webhook", ctr.GETFacebookWebhook)
	router.POST("/facebook/webhook", ctr.POSTFacebookWebhook)

	router.Run(":" + port)
}
