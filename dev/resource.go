package dev

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Resource struct {
}

func (r *Resource) Doc(c *gin.Context) {
	c.HTML(http.StatusOK, "apidoc.html.tmpl", gin.H{})
}
