package api

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelgbenle/ZeinaMFI/internal/models"
	"github.com/michaelgbenle/ZeinaMFI/internal/util"
	"log"
	"net/http"
)

func (u *HTTPHandler) DepositHandler(c *gin.Context) {
	_, err := u.GetUserFromContext(c)
	if err != nil {
		util.Response(c, "Unauthorized", http.StatusUnauthorized, nil, []string{"unauthorized"})
		return
	}

	deposit := &models.Money{}
	err = c.ShouldBindJSON(&deposit)
	if err != nil {
		util.Response(c, "error", 400, nil, []string{"invalid request"})
		return
	}
	//validate amount
	if deposit.Amount <= 0 {
		util.Response(c, "error", 400, nil, []string{"invalid amount"})
		return
	}

	//check if account number exists
	user, err := u.Repository.FindUserByAccountNos(deposit.AccountNo)
	if err != nil {
		util.Response(c, "error", 400, nil, []string{"account number does not exist"})
		return
	}

	//ensure only customers can deposit into their account
	if user.UserType == models.Admin {
		util.Response(c, "error", 400, nil, []string{"cannot deposit into admin account"})
		return
	}

	//deposit user's money
	transaction, depositErr := u.Repository.Deposit(deposit, user)
	if depositErr != nil {
		log.Println(depositErr)
		util.Response(c, "unable to deposit into user account", 500, nil, []string{"deposit error"})
		return
	}

	util.Response(c, "account credited successfully", 200, transaction, nil)

}
