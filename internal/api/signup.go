package api

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelgbenle/ZeinaMFI/internal/models"
	"github.com/michaelgbenle/ZeinaMFI/internal/util"
)

func (u *HTTPHandler) SignUpHandler(c *gin.Context) {
	var user *models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		util.Response(c, "error", 400, nil, []string{"invalid request"})
		return
	}

	// check if email is valid
	if !util.IsValidEmail(user.Email) {
		util.Response(c, "error", 400, nil, []string{"invalid email"})
		return
	}

	//check if password is valid
	if !util.IsValidPassword(user.Password) {
		util.Response(c, "error", 400, nil, []string{"minimum of 8 characters containing upper case, lower case, number and special character required"})
		return
	}

	//check if user already exists
	_, err = u.Repository.FindUserByEmail(user.Email)
	if err == nil {
		util.Response(c, "error", 400, nil, []string{"user already exists"})
		return
	}

	//save user to database
	err = u.Repository.CreateUser(user)
	if err != nil {
		util.Response(c, "unable to sign up user", 500, nil, []string{"internal server error"})
		return
	}

	// successful sign up
	util.Response(c, "sign up successful", 201, nil, nil)
}
