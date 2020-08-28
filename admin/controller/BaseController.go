package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Base struct {
}

func (this Base) BaseController(c *gin.Context) {
	fmt.Println("====", c.FullPath())
}
