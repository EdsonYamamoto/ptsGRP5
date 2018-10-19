package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GETIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl.html", nil)
}

func GETDocumentation(c *gin.Context) {
	c.HTML(http.StatusOK, "documentation.tmpl.html", nil)
}

func GETTeste(c *gin.Context) {
	c.HTML(http.StatusOK, "testImplementation.tmpl.html", nil)
}
