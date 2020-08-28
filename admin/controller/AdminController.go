package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Admin struct {
}

func (this Admin) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/index.html", gin.H{})
}
func (this Admin) User(c *gin.Context) {

}
