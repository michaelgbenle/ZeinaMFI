package api

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelgbenle/ZeinaMFI/internal/util"
)

func (u *HTTPHandler) GetAllUsersHandler(c *gin.Context) {
	users, err := u.Repository.GetAllUsers()
	if err != nil {
		util.Response(c, "error", 400, nil, []string{"unable to get all users"})
		return
	}
	util.Response(c, "success", 200, users, nil)
}
