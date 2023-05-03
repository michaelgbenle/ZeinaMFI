package api

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelgbenle/ZeinaMFI/internal/util"
)

//LiveHandler is to check if server is up
func (u *HTTPHandler) LiveHandler(c *gin.Context) {
	data := "server is up and running"

	// healthcheck
	util.Response(c, "Live", 200, data, nil)
}
