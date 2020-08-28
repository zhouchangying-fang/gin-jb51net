package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
}

func (this Login) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "login/index.html", gin.H{})
}
func (this Login) Load(c *gin.Context) {
	var user = struct {
		Name string `form:"name"`
		Pwd  string `form:"pwd"`
	}{}
	if c.ShouldBindJSON(&user) == nil {

	}
}
