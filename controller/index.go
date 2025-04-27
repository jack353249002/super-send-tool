package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminView(c *gin.Context) {
	HtmlReturn(c, "index", nil)
	return
}

// Html 返回值
func HtmlReturn(c *gin.Context, name string, data interface{}) {
	c.HTML(http.StatusOK, name+".html", data)

}
